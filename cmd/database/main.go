package main

import (
	"log"
	"xserver/app/db"
	"xserver/configs"
	"xserver/model"

	"github.com/xaces/xutils"
	"github.com/xaces/xutils/orm"
)

type jsonMenu struct {
	Data []model.SysMenu `json:"data"`
}

func main() {
	if err := configs.Load(".config.yaml"); err != nil {
		log.Fatalln(err)
	}
	if err := db.Run(); err != nil {
		log.Fatalln(err)
	}

	if orm.DbCount(&model.SysMenu{}, nil) < 1 {
		var v jsonMenu
		if err := xutils.JSONFile("conf/menu.json", &v); err != nil {
			log.Println(err)
		} else {
			orm.DB().Create(&v.Data)
		}
	}
}
