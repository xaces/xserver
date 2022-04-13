package cache

import (
	"xserver/util"

	"github.com/wlgd/xutils"
)

type devUser struct {
	Val       *xutils.BitMap
	DeviceIds string
}

var gUserDevs map[uint]*devUser = make(map[uint]*devUser)

func NewDevUser(userId uint, idstr string) *devUser {
	u := &devUser{
		Val:       xutils.NewBitMap(64),
		DeviceIds: idstr,
	}
	u.Set(idstr)
	gUserDevs[userId] = u
	return u
}

func UserDevs(userId uint) *devUser {
	if v, ok := gUserDevs[userId]; ok {
		return v
	}
	return nil
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

// 只更新bitmap
func (d *devUser) Update(id int) {
	if d.Val.Include(id) {
		return
	}
	d.Val.Set(id)
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
	return d.Val.Include(id)
}
