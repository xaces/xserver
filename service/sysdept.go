package service

import "github.com/wlgd/xutils/orm"

// DeptPage 查询页
type DeptPage struct {
	orm.DbPage
	DeptName string `form:"deptName"` // 部门名称
}

// Where 初始化
func (s *DeptPage) Where() *orm.DbWhere {
	var where orm.DbWhere
	where.String("dept_name like ?", s.DeptName)
	return &where
}
