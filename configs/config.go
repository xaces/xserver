package configs

import (
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	// Conf 配置
	GViper     *viper.Viper
	gPublicDir string
)

// Load 初始化配置参数
func Load(filename string) error {
	gPublicDir = filepath.Dir(os.Args[0])
	GViper = viper.New()
	GViper.SetConfigFile(Public(filename))
	if err := GViper.ReadInConfig(); err != nil {
		return err
	}
	public := GViper.GetString("public")
	if path.IsAbs(public) {
		gPublicDir = public
	} else {
		gPublicDir = Public(public)
	}
	return nil
}

func Public(dirs ...string) string {
	public := gPublicDir
	for _, v := range dirs {
		public = public + string(os.PathSeparator) + v
	}
	return public
}
