package model

// SysDept 操作请求(获取/修改/更新)
type SysDept struct {
	Id       uint64 `json:"deptId" gorm:"primary_key"`
	ParentId uint64 `json:"parentId"`
	DeptName string `json:"deptName"`
	Address  string `json:"address"`
	Sort     int    `json:"sort"`
	Enable   uint8  `json:"enable" gorm:"default:1;comment:0禁用1启动;"`
}

func (o *SysDept) TableName() string {
	return "t_sysdept"
}
