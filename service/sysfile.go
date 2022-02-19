package service

import "github.com/wlgd/xutils/orm"

// FilePage 查询页
type FilePage struct {
	basePage
}

// Where 初始化
func (s *FilePage) Where() *orm.DbWhere {
	var where orm.DbWhere
	return &where
}
