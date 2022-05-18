package model

type SysTationOpt struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"` // 名称
	Scheme       string `json:"scheme"`
	Host         string `json:"host" gorm:"comment:地址;"`    //
	Max          int    `json:"max" gorm:"comment:最大接入数目;"` //
	Details      string `json:"details"`
	OrganizeGUID string `json:"organizeGuid" gorm:"default:'';"`
	Enable       bool   `json:"enable" gorm:"default:1;comment:0禁用 1启动;"`
	Status       uint8  `json:"status" gorm:"-;comment:状态;"`     //
	Access       int    `json:"access" gorm:"-;comment:当前接入数目;"` //
	Logins       int    `json:"logins" gorm:"-;comment:当前在线数目;"` //
}

type SysTation struct {
	SysTationOpt
	GUID      string `json:"guid"`
	CreatedAt jtime  `json:"createdAt"`
}

func (o *SysTation) TableName() string {
	return "t_systation"
}
