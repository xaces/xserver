package service

import (
	"errors"
	"xserver/model"
	"xserver/pkg/gmd5"
	"xserver/util"

	"github.com/wlgd/xutils/orm"
)

const (
	defaultpwd = "123456"
)

type Vehicle struct {
	Id          int    `json:"deviceId"`
	DeviceNo    string `json:"deviceNo"`
	DeviceName  string `json:"deviceName"`
	ChlCount    int    `json:"chlCount"`
	ChlNames    string `json:"chlNames"`
	OrganizeId  int    `json:"organizeId"`  // 分组Id
	StationGuid string `json:"stationGuid"` // 工作站guid
}

// UserPage 查询页
type UserPage struct {
	orm.DbPage
	KeyWord string `form:"keyWord"`
}

// Where 初始化
func (s *UserPage) Where() *orm.DbWhere {
	where := s.DbWhere()
	if s.KeyWord != "" {
		where.Append("user_name like ? or nick_name like ?", s.KeyWord, s.KeyWord)
	}
	return where
}

// SysUserPassword 生成密码
func SysUserPassword(u *model.SysUser, password string) string {
	if u.Salt == "" {
		u.Salt = util.StringRandom(6)
	}
	token := u.UserName + password + u.Salt
	return gmd5.MustEncryptString(token)
}

func SysUserIsExist(req *model.SysUser) error {
	var user model.SysUser
	db := orm.DB().Or("user_name like ?", req.UserName)
	if req.Phone != "" {
		db = db.Or("phone like ?", req.Phone)
	}
	if req.Email != "" {
		db = db.Or("email like ?", req.Email)
	}
	return db.First(&user).Error
}

func SysUserCreate(u *model.SysUser) error {
	if SysUserIsExist(u) == nil {
		return errors.New("user already exists")
	}
	u.Enable = 1
	if u.Password == "" {
		u.Password = defaultpwd
	}
	u.Password = SysUserPassword(u, u.Password)
	return orm.DbCreate(u)
}

func SysUserDevice(t *model.SysUserToken) []Vehicle {
	var res []Vehicle
	orm.DB().Model(&model.OprVehicle{}).Where("organize_guid = ?", t.OrganizeGuid).Scan(&res)
	return res
}
