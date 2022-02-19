package model

// SysNotice
type SysNotice struct {
	Id         uint64 `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Sender     string `json:"sender"`
	SenderName string `json:"senderName"`
	Accept     string `json:"accept"`
	AcceptName string `json:"acceptName"`
	Type       string `json:"type"`
	CreatedAt  jtime  `json:"createTime" gorm:"column:created_time;"`
}

func (o *SysNotice) TableName() string {
	return "t_sysnotice"
}
