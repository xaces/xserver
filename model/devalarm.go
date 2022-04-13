package model

// DevAlarm 报警
type DevAlarm struct {
	Guid      string `json:"guid" gorm:"primary_key;index"`
	DeviceNo  string `json:"deviceNo" gorm:"type:varchar(24);"`
	DTU       string `json:"dtu"`
	AlarmType int    `json:"alarmType"`                          // 类型
	StartData string `json:"startData"`                          // gps信息 json 字符串
	EndData   string `json:"endData"`                            // gps信息 json 字符串
	StartTime string `json:"startTime" gorm:"type:varchar(20);"` // 开始时间
	EndTime   string `json:"endTime" gorm:"type:varchar(20);"`   // 结束时间
}

// TableName 表名
func (s *DevAlarm) TableName() string {
	return "t_devalarm"
}

type DevAlarmHistory DevAlarm

func (s *DevAlarmHistory) TableName() string {
	return "t_devalarmhistory"
}

type lnType int

const (
	AlarmLinkUnknow      = lnType(0x00)
	AlarmLinkDev         = lnType(0x01)
	AlarmLinkFtpFile     = lnType(0x02)
	AlarmLinkStorageFile = lnType(0x03)
)

// 报警关联信息
type DevAlarmDetails struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	DeviceNo string `json:"deviceNo" gorm:"type:varchar(24);"`
	DTU      string `json:"dtu"`
	Guid     string `json:"guid"`     // guid和DevAlarm.guid用来关联
	Flag     uint8  `json:"flag"`     // 0-实时 1-补传
	Status   int    `json:"status"`   // 0-开始 1--结束 2--报警中
	LinkType lnType `json:"linkType"` // 数据类型
	Data     string `json:"data"`
}

// TableName 表名
func (s *DevAlarmDetails) TableName() string {
	return "t_devalarmdetails"
}

// 报警关联信息
type DevAlarmFile struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	DeviceNo  string `json:"deviceNo" gorm:"type:varchar(24);"`
	DTU       string `json:"dtu"`
	AlarmType int    `json:"alarmType"`
	Guid      string `json:"guid"` // guid和DevAlarm.guid用来关联
	LinkType  lnType `json:"linkType"`
	Channel   int    `json:"channel"`
	Size      int    `json:"size"`
	Duration  int    `json:"duration"`
	FileType  int    `json:"fileType"`
	Name      string `json:"name"`
}

// TableName 表名
func (s *DevAlarmFile) TableName() string {
	return "t_devalarmfile"
}