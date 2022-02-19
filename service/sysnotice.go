package service

import "github.com/wlgd/xutils/orm"

// NoticePage 查询页
type NoticePage struct {
	basePage
}

// Where 初始化
func (s *NoticePage) Where() *orm.DbWhere {
	var where orm.DbWhere
	return &where
}
