package service

import (
	"xserver/model"

	"github.com/wlgd/xutils/orm"
)

func SysPostCheckAdd(req *model.SysPost) error {
	var post model.SysPost
	if err := orm.DbFirstBy(&post, "name like ?", req.Name); err != nil {
		return err
	}
	return nil
}
