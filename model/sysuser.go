package model

const (
	// SysUserRoleId 系统用户角色
	SysUserRoleId = 0
	// SysUserName 系统用户名
	SysUserName = "admin"
)

// 超级用户可以创建主账号
// 主账号只能可以创建普通账号
const (
	SysUserTypeRoot = iota
	SysUserTypeAdmin
	SysUserTypeComm
)

type SysUserToken struct {
	Id           uint64   `json:"userId" gorm:"primary_key"`
	UserName     string   `json:"username" gorm:"not null;comment:登录账号;"`
	RoleId       uint64   `json:"roleId"`
	DeptId       JUint64s `json:"deptIds"`
	OrganizeGuid string   `json:"organizeGuid"`
	OrganizeName string   `json:"organizeName"`
	Host         string   `json:"host" gorm:"-"`
	Scheme       string   `json:"scheme" gorm:"-"`
}

// SysUserOpt
type SysUserOpt struct {
	SysUserToken
	NickName  string `json:"nickname" gorm:"not null;comment:用户昵称;"`
	UserType  int    `json:"userType"`
	DeviceIds string `json:"deviceIds"` // *,代表所有， 1,2,3
	Email     string `json:"email" gorm:"comment:用户邮箱;"`
	Phone     string `json:"phone" gorm:"size:11;default:'';comment:手机号码;"`
	Sex       string `json:"sex" gorm:"size:1;default:'0';comment:用户性别(1男 2女 0未知);"`
	Avatar    string `json:"avatar" gorm:"comment:头像路径;"`
	Password  string `json:"-" gorm:"not null;comment:密码;"`
	Salt      string `json:"-" gorm:"not null;comment:盐加密;"`
	Enable    uint8  `json:"enable" gorm:"default:1;comment:0禁用 1启动;"`
	Details   string `json:"details"`
}

//1、主账号绑定工作站
//2、每个子账号记录主账号，根据主账号获取绑定的工作站信息
//3、添加设备时，记录设备对应的主账号，获取设备根据主账号信息，只获取当前主账号的设备
//4、把设备分配给对应的账号DeviceIds, 上级账号可以可重复分配设备到下级账号

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
