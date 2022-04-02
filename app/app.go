package app

import (
	"xserver/app/db"
	"xserver/app/router"
	"xserver/configs"
	"xserver/entity/nats"
)

func Run() error {
	if err := db.Run(&configs.Default.Sql); err != nil {
		return err
	}
	router.Run(&configs.Default.Http)
	nats.Default.Run("")
	return nil
}

func Shutdown() error {
	router.Shutdown()
	return nil
}
