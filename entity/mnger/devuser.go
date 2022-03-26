package mnger

import (
	"fmt"
	"strconv"
	"xserver/middleware"
	"xserver/model"
	"xserver/util"

	"github.com/gin-gonic/gin"
	"github.com/wlgd/xutils/ctx"
)

type devUser struct {
	Val       map[int]string
	DeviceIds string
}

var UserDevs map[uint64]*devUser = make(map[uint64]*devUser)

func NewDevUser(userId uint64, idstr string) {
	u := &devUser{
		Val:       make(map[int]string),
		DeviceIds: idstr,
	}
	u.Set(idstr)
	UserDevs[userId] = u
}

func (d *devUser) Set(idstr string) {
	if idstr == "*" {
		d.DeviceIds = "*"
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	for _, v := range ids {
		d.Val[v] = ""
	}
	if d.DeviceIds != "" {
		d.DeviceIds += ","
	}
	d.DeviceIds += idstr
}

func (d *devUser) Dels(idstr string) {
	if idstr == "*" {
		for k := range d.Val {
			delete(d.Val, k)
		}
		d.DeviceIds = ""
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	for _, v := range ids {
		delete(d.Val, v)
	}
	var s string
	for k := range d.Val {
		s += strconv.Itoa(int(k))
		s += ","
	}
	if s != "" {
		s = s[:len(s)-1]
	}
	d.DeviceIds = s
}

func (d *devUser) Include(id int) bool {
	_, ok := d.Val[id]
	return ok
}

type Device struct {
	Id             int    `json:"deviceId"`
	DeviceNo       string `json:"deviceNo"`
	DeviceName     string `json:"deviceName"`
	ChlCount       int    `json:"chlCount"`
	ChlNames       string `json:"chlNames"`
	IoCount        int    `json:"ioCount"`
	IoNames        string `json:"ioNames"`
	Icon           string `json:"icon"`
	OrganizeGuid   string `json:"organizeGuid"`
	Details        string `json:"details"`
	Type           string `json:"type"`
	Guid           string `json:"guid"`
	Version        string `json:"version"`
	Online         bool   `json:"online"`
	LastOnlineTime string `json:"lastOnlineTime"`
	EffectiveTime  string `json:"effectiveTime"`
	CreatedAt      string `json:"createdTime"`
	UpdatedAt      string `json:"updatedTime"`
}

type page struct {
	Num    int    `form:"pageNum"`  // 当前页码
	Size   int    `form:"pageSize"` // 每页数
	UserId uint64 `form:"userId"`
}

func DevUserList(c *gin.Context) {
	var p page
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	tok := middleware.GetUserToken(c)
	var res []Device
	if err := GetUserDevice(tok, &res); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	u, ok := UserDevs[p.UserId] // 指定用户
	if !ok {
		ctx.JSONOk().WriteTo(c)
		return
	}
	if u.DeviceIds == "*" {
		ctx.JSONOk().Write(gin.H{"total": len(res), "data": res}, c)
		return
	}
	var data []Device
	for _, v := range res {
		if !u.Include(v.Id) {
			continue
		}
		data = append(data, v)
	}
	total := len(data)
	if p.Size < 1 {
		ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
		return
	}
	offset := p.Num * p.Size
	ctx.JSONOk().Write(gin.H{"total": total, "data": data[offset-p.Size : offset]}, c)
}

func GetUserDevice(t *model.SysUserToken, data interface{}) error {
	url := fmt.Sprintf("http://%s/station/api/device/list?organizeGuid=%s", t.Host, t.OrganizeGuid)
	return util.HttpGet(url, data)
}
