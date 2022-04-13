package model

type OprVehicleOpt struct {
	Id            uint   `json:"deviceId" gorm:"primary_key"`
	DeviceNo      string `json:"deviceNo"`
	DeviceName    string `json:"deviceName"`
	Icon          string `json:"icon"`
	ChlCount      int    `json:"chlCount"`
	ChlNames      string `json:"chlNames" gorm:"comment:通道别名,隔开;"`
	IoCount       int    `json:"ioCount"`
	IoNames       string `json:"ioNames" gorm:"comment:io别名,隔开;"`
	OrganizeId    uint   `json:"organizeId"`                      // 分组Id
	OrganizeGuid  string `json:"organizeGuid" gorm:"default:'';"` // 所属组织Guid
	StationGuid   string `json:"stationGuid"`                     // 工作站guid
	EffectiveTime string `json:"effectiveTime"`
	Details       string `json:"details"`
}

type OprVehicle struct {
	OprVehicleOpt
	Type           string `json:"type" gorm:"type:varchar(20);"`
	Guid           string `json:"guid" gorm:"type:varchar(64);"`
	Version        string `json:"version" gorm:"type:varchar(20);"`
	Online         bool   `json:"online"`
	LastOnlineTime string `json:"lastOnlineTime"`
	CreatedAt      jtime  `json:"createTime"`
	UpdatedAt      jtime  `json:"updateTime"`
}

// TableName 表名
func (s *OprVehicle) TableName() string {
	return "t_oprvehicle"
}
