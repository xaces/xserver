package db

import (
	"errors"
	"xserver/model"

	"github.com/wlgd/xutils/orm"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initTables(db *gorm.DB) {
	db.AutoMigrate(&model.SysMenu{}, &model.SysRole{}, &model.SysUser{},
		&model.SysDictType{}, &model.SysDictData{}, &model.SysDept{}, &model.SysPost{}, &model.SysFile{}, &model.SysStation{})
	db.AutoMigrate(&model.OprOrganization{})
	orm.SetDB(db.Debug())
}

var (
	gconf = gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	db *gorm.DB
)

type Options struct {
	Name    string
	Address string
}

// Run 初始化服务
func Run(o *Options) error {
	switch o.Name {
	case "sqlite":
		db, _ = gorm.Open(sqlite.Open(o.Address), &gconf)
	case "mysql":
		db, _ = gorm.Open(mysql.New(mysql.Config{
			DSN: o.Address,
			// DefaultStringSize:         64,    // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据版本自动配置
		}), &gconf)
	}
	if db == nil {
		return errors.New("db invalid")
	}
	sqldb, err := db.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqldb.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqldb.SetMaxOpenConns(100)
	initTables(db)
	return nil
}
