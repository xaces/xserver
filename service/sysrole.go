package service

import (
	"xserver/model"

	"github.com/wlgd/xutils/orm"
)

// RolePage 查询页
type RolePage struct {
	basePage
	MenuName string `form:"menuName"` // 菜单名称
	Visible  string `form:"visible"`  // 菜单状态
}

// Where 初始化
func (s *RolePage) Where() *orm.DbWhere {
	var where orm.DbWhere
	where.String("menu_name like ?", s.MenuName)
	where.String("visible = ?", s.Visible)
	return &where
}

// QueryRoleById 查询指定权限byID
func DbQueryRoleById(id uint64) (role model.SysRole, err error) {
	err = orm.DB().Where("id = ?", id).Preload("SysMenus").First(&role).Error
	return
}

// DbDelSysMenusByRoleId 删除角色关联数据
func DbDelSysMenusByRoleId(id uint64) error {
	// role, err := DbQueryRoleById(id)
	// if err != nil {
	// 	return err
	// }
	// if len(role.SysMenus) <= 0 {
	// 	return nil
	// }
	// return orm.DB().Model(&role).Association("SysMenus").Delete(&role.SysMenus) // 删除关联数据
	return nil
}
