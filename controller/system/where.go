package system

import "github.com/wlgd/xutils/orm"

// UserPage 查询页
type Where struct {
	orm.DbPage
	KeyWord   string `form:"keyWord"`
	PostName  string `form:"postName"` //
	PostCode  string `form:"postCode"` //
	Status    string `form:"status"`
	MenuName  string `form:"menuName"`  // 菜单名称
	DictType  string `form:"dictType"`  // 字典名称
	DictLabel string `form:"dictLabel"` // 字典标签

	createdBy string `form:"-"`
}

func (o *Where) Where() *orm.DbWhere {
	return o.DbWhere()
}

// User
func (o *Where) User() *orm.DbWhere {
	where := o.DbWhere()
	where.Equal("created_by", o.createdBy)
	return where
}

// User
func (o *Where) Menu() *orm.DbWhere {
	where := o.DbWhere()
	where.Like("title", o.MenuName)
	return where
}

// Role
func (o *Where) Role() *orm.DbWhere {
	where := o.DbWhere()
	where.Equal("created_by", o.createdBy)
	return where
}
