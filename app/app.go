package app

import (
	"xserver/app/db"
	"xserver/app/router"
	"xserver/configs"
)

func Run() error {
	if err := db.Run(&configs.Default.Sql); err != nil {
		return err
	}
	router.Run(&configs.Default.Http)
	return nil
}

func Shutdown() error {
	router.Shutdown()
	return nil
}
