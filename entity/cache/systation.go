package cache

import (
	"xserver/model"

	"github.com/wlgd/xutils/orm"
)

var (
	gSysTations = make(map[string]*model.SysTation)
	gDefStation = model.SysTation{
		SysTationOpt: model.SysTationOpt{
			Scheme: "http",
			Host:   "127.0.0.1:12100",
		},
	}
)

func SysTation(guid string) *model.SysTation {
	if guid == "" {
		return &gDefStation
	}
	if v, ok := gSysTations[guid]; ok {
		return v
	}
	v := &model.SysTation{}
	if err := orm.DbFirstBy(v, "guid = ?", guid); err != nil {
		return nil
	}
	gSysTations[guid] = v
	return v
}
