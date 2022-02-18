package service

import "github.com/wlgd/xutils/orm"

// DictDataPage 查询页
type DictDataPage struct {
	basePage
	DictType string `form:"dictType"` // 字典名称
}

// Where 初始化
func (s *DictDataPage) Where() *orm.DbWhere {
	var where orm.DbWhere
	where.String("dict_type like ?", s.DictType)
	return &where
}

// DictTypePage 查询页
type DictTypePage struct {
	basePage
	DictType  string `form:"dictType"`  // 字典名称
	DictLabel string `form:"dictLabel"` // 字典标签
}

// Where 初始化
func (s *DictTypePage) Where() *orm.DbWhere {
	var where orm.DbWhere
	where.String("dict_label like ?", s.DictLabel)
	where.String("dict_type like ?", s.DictType)
	return &where
}
