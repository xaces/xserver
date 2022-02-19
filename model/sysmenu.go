package model

// SysMenu 权限
type SysMenu struct {
	Id          uint64 `json:"id" gorm:"primary_key"`
	ParentId    uint64 `json:"parentId"`
	Title       string `json:"title" gorm:"size:50;not null;comment:菜单名称;"`
	Type        string `json:"type" gorm:"size:1;default:'0';comment:;"`
	OpenType    string `json:"openType" gorm:"comment:菜单类型;"`
	Icon        string `json:"icon" gorm:"comment:菜单图标;"`
	Href        string `json:"href" gorm:"comment:链接;"`
	CreatedAt   jtime  `json:"createTime" gorm:"column:created_time;"`
	Description string `json:"description"`
	CheckArr    string `json:"checkArr" gorm:"-"`
}

func (o *SysMenu) TableName() string {
	return "t_sysmenu"
}
