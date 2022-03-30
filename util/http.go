package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// HttpGet http get 请求
func HttpGet(url string, result interface{}) error {
	recv, err := http.Get(url)
	if err != nil {
		return err
	}
	defer recv.Body.Close()
	content, err := ioutil.ReadAll(recv.Body)
	if err != nil {
		return err
	}
	var res response
	if err := jsoniter.Unmarshal(content, &res); err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Msg)
	}
	if res.Data != nil && result != nil {
		jsoniter.Get(content, "data").ToVal(result)
	}
	return nil
}

func SingleHostProxy(api *url.URL, path string, c *gin.Context) {
	c.Request.URL.Path = path
	httputil.NewSingleHostReverseProxy(api).ServeHTTP(c.Writer, c.Request)
}
