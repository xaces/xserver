package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/wlgd/xutils/ctx"

	"xserver/entity/cache"
	"xserver/entity/nats"
	"xserver/middleware"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func msgHandler(userId uint, conn *websocket.Conn, ch chan []byte) {
	devs := cache.UserDevs(userId)
	for v := range ch { // 关闭chan自动退出
		if v == nil {
			continue
		}
		deviceId := jsoniter.Get(v, "deviceId").ToInt()
		if devs != nil && devs.Include(deviceId) {
			conn.WriteMessage(websocket.TextMessage, v)
		}
	}
}

// WsHandler webs
func WsHandler(c *gin.Context) {
	tokenstr := c.Query("token")
	claims, err := middleware.NewJWT().ParseToken(tokenstr)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()
	msgChan := nats.Default.Push(c.Request.RemoteAddr)
	if msgChan == nil {
		ctx.JSONWriteError(errors.New("subscribe server"), c)
		return
	}
	go msgHandler(claims.SysUserToken.Id, ws, msgChan)
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	nats.Default.Pop(c.Request.RemoteAddr)
}
