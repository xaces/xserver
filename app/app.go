package app

import (
	"xserver/app/db"
	"xserver/app/router"
	"xserver/entity/subject"
)

func Run() error {
	if err := db.Run(); err != nil {
		return err
	}
	if err := router.Run(); err != nil {
		return err
	}
	subject.Default.Run("nats://127.0.0.1:4222")
	return nil
}

func Shutdown() error {
	router.Shutdown()
	return nil
}
