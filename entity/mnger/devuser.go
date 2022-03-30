package mnger

import (
	"strconv"
	"xserver/util"
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
