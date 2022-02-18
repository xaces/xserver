package main

import (
	"log"
	"os"
	"xserver/app"
	"xserver/configs"

	svc "github.com/kardianos/service"
)

type program struct {
}

func (p *program) Start(s svc.Service) error {
	if err := configs.Load(".config.toml"); err != nil {
		return err
	}
	return p.run()
}

func (p *program) run() error {
	return app.Run()
}

func (p *program) Stop(s svc.Service) error {
	return app.Shutdown()
}

func main() {
	svvconfig := &svc.Config{
		Name:        "xvms.server",
		DisplayName: "xserver",
		Description: "This is server application",
	}
	s, err := svc.New(&program{}, svvconfig)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err = s.Install()
		} else if os.Args[1] == "uninstall" {
			err = s.Uninstall()
		}
		log.Println(err)
		return
	}
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
