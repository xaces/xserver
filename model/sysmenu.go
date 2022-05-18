package model

// SysMenuOpt
type SysMenuOpt struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	ParentID uint   `json:"parentId"`
	Title    string `json:"title" gorm:"size:50;not null;comment:菜单名称;"`
	Type     string `json:"type" gorm:"size:1;default:'0';comment:;"`
	OpenType string `json:"openType" gorm:"comment:菜单类型;"`
	Icon     string `json:"icon" gorm:"comment:菜单图标;"`
	Href     string `json:"href" gorm:"comment:链接;"`
	Details  string `json:"details"`
	CheckArr string `json:"checkArr" gorm:"-"`
}

// SysMenu 权限
type SysMenu struct {
	SysMenuOpt
	CreatedAt jtime     `json:"createdAt"`
	Children  []SysMenu `json:"children,omitempty" gorm:"-"`
	// Children []SysMenu `json:"children,omitempty" gorm:"foreignKey:ParentID;"`
}

func (o *SysMenu) TableName() string {
	return "t_sysmenu"
}

func (o *SysMenu) FindChildren(data []SysMenu) {
	for _, v := range data {
		if v.ParentID != o.ID {
			continue
		}
		v.FindChildren(data)
		o.Children = append(o.Children, v)
	}
}
