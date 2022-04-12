package mnger

import (
	"xserver/model"

	"github.com/wlgd/xutils/orm"
)

var (
	sysTation  = make(map[string]*model.SysTation)
	defStation = model.SysTation{
		SysTationOpt: model.SysTationOpt{
			Scheme: "http",
			Host:   "127.0.0.1:12100",
		},
	}
)

func SysTationGet(guid string) *model.SysTation {
	if guid == "" {
		return &defStation
	}
	if v, ok := sysTation[guid]; ok {
		return v
	}
	v := &model.SysTation{}
	if err := orm.DbFirstBy(v, "guid = ?", guid); err != nil {
		return nil
	}
	sysTation[guid] = v
	return v
}
