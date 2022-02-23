package model

// SysDept
type SysDeptOpt struct {
	Id       uint64 `json:"deptId" gorm:"primary_key"`
	ParentId uint64 `json:"parentId"`
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
	CreatedAt jtime `json:"createTime" gorm:"column:created_time;"`
}

func (o *SysDept) TableName() string {
	return "t_sysdept"
}
