package model

// SysMenu 权限
type SysMenu struct {
	Id        uint64 `json:"id" gorm:"primary_key"`
	ParentId  uint64 `json:"parentId"`
	Title     string `json:"title" gorm:"size:50;not null;comment:菜单名称;"`
	Icon      string `json:"icon" gorm:"comment:菜单图标;"`
	Href      string `json:"href" gorm:"comment:链接;"`
	OpenType  string `json:"openType" gorm:"comment:菜单类型;"`
	Sort      int    `json:"sort" gorm:"comment:显示顺序;"`
	Enable    uint8  `json:"enable" gorm:"default:1;comment:0禁用1启动;"`
	PowerCode string `json:"powerCode" gorm:"comment:菜单状态（0显示 1隐藏）;"`
	Type      string `json:"type" gorm:"size:1;default:'0';comment:;"`
	CreatedAt jtime  `json:"createTime" gorm:"column:created_time;"`
	Remark    string `json:"remark" gorm:"size:500;comment:备注;"`
	CheckArr  string `json:"checkArr" gorm:"-"`
}

func (o *SysMenu) TableName() string {
	return "t_sysmenu"
}
