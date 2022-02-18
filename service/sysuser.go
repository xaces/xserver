package service

import (
	"xserver/model"
	"xserver/pkg/gmd5"
	"xserver/util"

	"github.com/wlgd/xutils/orm"
)

// NewSysPassword 生成密码
func NewSysPassword(u *model.SysUser, password string) string {
	if u.Salt == "" {
		u.Salt = util.RandomString(6)
	}
	token := u.UserName + password + u.Salt
	return gmd5.MustEncryptString(token)
}

// UserPage 查询页
type UserPage struct {
	basePage
	Page    int    `form:"page"`  // 每页数
	Limit   int    `form:"limit"` // 当前页码
	KeyWord string `form:"keyWord"`
}

// Where 初始化
func (s *UserPage) Where() *orm.DbWhere {
	var where orm.DbWhere
	if s.KeyWord != "" {
		where.Append("user_name like ? or nick_name like ?", s.KeyWord, s.KeyWord)
	}
	return &where
}

func CheckAddUser(req *model.SysUser) error {
	var user model.SysUser
	if err := orm.DbFirstBy(&user, "user_name like ?", req.UserName); err != nil {
		return err
	}

	if req.Phone != "" {
		if err := orm.DbFirstBy(&user, "phone like ?", req.Phone); err != nil {
			return err
		}
	}
	if req.Email != "" {
		if err := orm.DbFirstBy(&user, "email like ?", req.Email); err != nil {
			return err
		}
	}
	return nil
}
