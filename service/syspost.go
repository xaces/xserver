package service

import (
	"xserver/model"

	"github.com/wlgd/xutils/orm"
)

// PostPage 查询页
type PostPage struct {
	BasePage
	PostName string `form:"postName"` //
	PostCode string `form:"postCode"` //
	Status   string `form:"status"`
}

// Where 初始化
func (s *PostPage) Where() *orm.DbWhere {
	var where orm.DbWhere
	where.String("post_name like ?", s.PostName)
	where.String("post_code like ?", s.PostCode)
	where.String("status = ?", s.Status)
	return &where
}

func PostCheckAdd(req *model.SysPost) error {
	var post model.SysPost
	if err := orm.DbFirstBy(&post, "post_name like ?", req.PostName); err != nil {
		return err
	}
	return nil
}
