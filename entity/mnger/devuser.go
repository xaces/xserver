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

func NewDevUser(userId uint64, idstr string) *devUser{
	if v, ok := UserDevs[userId]; ok {
		return v
	}
	u := &devUser{
		Val:       xutils.NewBitMap(64),
		DeviceIds: idstr,
	}
	u.Set(idstr)
	UserDevs[userId] = u
	return u
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
