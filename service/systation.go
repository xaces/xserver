package service

import (
	"xserver/model"

	"github.com/wlgd/xutils/orm"
)

func sysStationDefault() *model.SysStation {
	val := &model.SysStation{}
	val.Name = "default"
	val.Max = 10
	val.Host = "127.0.0.1:12100"
	orm.DbCreate(val)
	return val
}
