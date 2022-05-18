package model

type SysFileOpt struct {
	ID         uint   `json:"fileId" gorm:"primary_key"`
	Name       string `json:"fileName"`
	Path       string `json:"filePath"`
	Type       string `json:"fileType"`
	Size       int64  `json:"fileSize"`
	Desc       string `json:"fileDesc"`
	TargetDate string `json:"targetDate"`
}

// SysFile
type SysFile struct {
	SysFileOpt
	CreatedAt jtime  `json:"createdAt"`
	CreatedBy string `json:"createdBy" gorm:"comment:创建者;"`
}

func (o *SysFile) TableName() string {
	return "t_sysfile"
}
