package model

type SysFileOpt struct {
	Id         uint64 `json:"fileId" gorm:"primary_key"`
	FileName   string `json:"fileName"`
	FilePath   string `json:"filePath"`
	FileType   string `json:"roleType"`
	FileSize   string `json:"fileSize"`
	FileDesc   string `json:"fileDesc"`
	TargetDate string `json:"targetDate"`
}

// SysFile
type SysFile struct {
	SysFileOpt
	CreatedAt jtime  `json:"createTime" gorm:"column:created_time;"`
	CreatedBy string `json:"createBy" gorm:"comment:创建者;"`
}

func (o *SysFile) TableName() string {
	return "t_sysfile"
}
