package configs

import (
	"fmt"
	"xserver/app/db"
	"xserver/app/router"
	"xserver/util"
)

type localConfigure struct {
	Address string `json:"address"`
}
type configure struct {
	Host string
	Http router.Option
	Sql  db.Options
}

// Default 所有配置参数
var (
	Default configure
	Local   localConfigure
)

// Load 初始化配置参数
func Load(filename string) error {
	if err := util.YamlFile(filename, &Default); err != nil {
		return err
	}
	Local.Address = fmt.Sprintf("%s:%d/%s", Default.Host, Default.Http.Port, Default.Http.Root)
	return nil
}
