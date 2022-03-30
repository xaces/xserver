package mnger

import (
	"xserver/util"

	"github.com/wlgd/xutils"
)

type devUser struct {
	Val       *xutils.BitMap
	DeviceIds string
}

var UserDevs map[uint64]*devUser = make(map[uint64]*devUser)

func NewDevUser(userId uint64, idstr string) {
	u := &devUser{
		Val:       xutils.DefaultBitMap,
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
		d.Val.Set(v)
	}
	if d.DeviceIds != "" {
		d.DeviceIds += ","
	}
	d.DeviceIds += idstr
}

func (d *devUser) Dels(idstr string) {
	if idstr == "*" {
		d.Val.Clear()
		d.DeviceIds = ""
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	for _, v := range ids {
		d.Val.Del(v)
	}
	d.DeviceIds = d.Val.String()
}

func (d *devUser) Include(id int) bool {
	if d.DeviceIds == "*" {
		return true
	}
	return d.Val.Include(id)
}
