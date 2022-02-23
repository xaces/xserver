package service

import "github.com/wlgd/xutils/orm"

// menuPage 查询页
type MenuPage struct {
	BasePage
	MenuName string `form:"menuName"` // 菜单名称
	Visible  string `form:"visible"`  // 菜单状态
}

// Where 初始化
func (s *MenuPage) Where() *orm.DbWhere {
	var where orm.DbWhere
	where.String("menu_name like ?", s.MenuName)
	where.String("visible = ?", s.Visible)
	return &where
}
