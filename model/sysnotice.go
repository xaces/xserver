package model

// SysNoticeOpt
type SysNoticeOpt struct {
	Id         uint   `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Sender     string `json:"sender"`
	SenderName string `json:"senderName"`
	Accept     string `json:"accept"`
	AcceptName string `json:"acceptName"`
	Type       string `json:"type"`
}

// SysNotice
type SysNotice struct {
	SysNoticeOpt
	CreatedAt jtime `json:"createTime" gorm:"column:created_time;"`
}

func (o *SysNotice) TableName() string {
	return "t_sysnotice"
}
