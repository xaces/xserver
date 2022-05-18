package service

import "xserver/model"

func SysMenuTree(data []model.SysMenu, id uint) (tree []model.SysMenu) {
	for _, v := range data {
		if v.ParentID != id {
			continue
		}
		v.FindChildren(data)
		tree = append(tree, v)
	}
	return
}
