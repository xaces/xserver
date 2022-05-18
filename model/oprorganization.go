package model

type OprOrganizationOpt struct {
	ID       uint   `json:"id" gorm:"primary_key"` //
	Name     string `json:"name"`                  // 名称
	ParentID uint   `json:"parentId"`
	Enable   bool   `json:"enable" gorm:"default:1;comment:0禁用 1启动;"`
	UserName string `json:"username"` // 管理账号
	Details  string `json:"details"`
}

type OprOrganization struct {
	OprOrganizationOpt
	GUID      string `json:"guid"`
	CreatedAt jtime  `json:"createdAt"`
	CreatedBy string `json:"createdBy" gorm:"comment:创建者;"`
}

// 第一级组织默认为公司
// 公司绑定绑定主账号，公司绑定工作站
// 已分组织，未分组织，把未分组织的设备分配到指定组织中

func (o *OprOrganization) TableName() string {
	return "t_oprorganization"
}
