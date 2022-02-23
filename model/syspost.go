package model

// SysPostOpt 操作请求(获取/修改/更新)
type SysPostOpt struct {
	Id       uint64 `json:"postId" gorm:"primary_key"`
	PostName string `json:"postName" gorm:"not null;comment:岗位名称;"`
	PostCode string `json:"postCode" gorm:"size:64;comment:岗位编码;"`
	PostSort int    `json:"postSort" gorm:"not null;comment:岗位排序;"`
	Status   string `json:"status" gorm:"size:1;default:'0';not null;comment:角色状态（0正常 1停用;"`
	Flag     bool   `json:"flag"`
	Details  string `json:"details"`
}

// SysRole 角色
// 下级角色权限最多和上级角色一样
type SysPost struct {
	SysPostOpt
	CreatedAt jtime `json:"createTime" gorm:"column:created_time;"`
}

func (o *SysPost) TableName() string {
	return "t_syspost"
}
