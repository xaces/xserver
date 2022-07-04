package db

import (
	"xserver/configs"
	"xserver/model"

	"github.com/xaces/xutils/orm"
)

type options struct {
	Name    string
	Address string
}

// Run 初始化服务
func Run() error {
	var o options
	if err := configs.GViper.UnmarshalKey("sql", &o); err != nil {
		return err
	}
	db, err := orm.NewGormV2(o.Name, o.Address)
	if err != nil {
		return err
	}
	db.AutoMigrate(
		&model.SysMenu{},
		&model.SysRole{},
		&model.SysUser{},
		&model.SysDictType{},
		&model.SysDictData{},
		&model.SysDept{},
		&model.SysPost{},
		&model.SysFile{},
		&model.SysTation{},
		&model.SysNotice{},
	)
	db.AutoMigrate(
		&model.OprOrganization{},
		&model.OprVehicle{},
	)
	orm.SetDB(db.Debug())
	return nil
}
