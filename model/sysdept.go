package model

// SysDept
type SysDeptOpt struct {
	ID       uint   `json:"deptId" gorm:"primary_key"`
	ParentID uint   `json:"parentId"`
	DeptName string `json:"deptName"`
	Address  string `json:"address"`
	Leader   string `json:"leader"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Sort     int    `json:"sort"`
	Enable   uint8  `json:"enable" gorm:"default:1;comment:0禁用1启动;"`
}

type SysDept struct {
	SysDeptOpt
	CreatedAt jtime     `json:"createdAt"`
	Children  []SysDept `json:"children,omitempty" gorm:"-"`
}

func (o *SysDept) TableName() string {
	return "t_sysdept"
}
