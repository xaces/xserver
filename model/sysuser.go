package model

const (
	// SysUserRoleId 系统用户角色
	SysUserRoleId = 0
	// SysUserName 系统用户名
	SysUserName = "admin"
)

// SysUserOpt
type SysUserOpt struct {
	Id       uint64 `json:"userId" gorm:"primary_key"`
	UserName string `json:"username" gorm:"not null;comment:登录账号;"`
	NickName string `json:"nickname" gorm:"not null;comment:用户昵称;"`
	Email    string `json:"email" gorm:"comment:用户邮箱;"`
	Phone    string `json:"phone" gorm:"size:11;default:'';comment:手机号码;"`
	Sex      string `json:"sex" gorm:"size:1;default:'0';comment:用户性别(1男 2女 0未知);"`
	Avatar   string `json:"avatar" gorm:"comment:头像路径;"`
	Password string `json:"-" gorm:"not null;comment:密码;"`
	Salt     string `json:"-" gorm:"not null;comment:盐加密;"`
	Enable   uint8  `json:"enable" gorm:"default:1;comment:0禁用 1启动;"`
	RoleId   uint64 `json:"roleId"`
	DeptId   uint64 `json:"deptId"`
	Details  string `json:"details"`
}

// SysUser 用户
type SysUser struct {
	SysUserOpt
	Login     uint8  `json:"login" gorm:"default:1;comment:0在线 1离线;"`
	CreatedAt jtime  `json:"createTime" gorm:"column:created_time;"`
	CreatedBy string `json:"createBy" gorm:"comment:创建者;"`
}

func (o *SysUser) TableName() string {
	return "t_sysuser"
}
