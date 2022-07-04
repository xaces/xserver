package service

import (
	"xserver/model"

	"github.com/xaces/xutils/orm"
)

// SysRoleQueryByID 查询指定权限byID
func SysRoleQueryByID(id uint64) (role model.SysRole, err error) {
	err = orm.DB().Where("id = ?", id).Preload("SysMenus").First(&role).Error
	return
}

// SysRoleDelSysMenusById 删除角色关联数据
func SysRoleDelSysMenusById(id uint64) error {
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
