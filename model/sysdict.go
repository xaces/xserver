package model

// SysDictDataOpt 操作请求(修改/更新)
type SysDictDataOpt struct {
	Id        uint64 `json:"dictCode" gorm:"primary_key"`
	Status    string `json:"status" gorm:"size:1;default:'0';comment:帐号状态(0正常 1停用);"`
	Remark    string `json:"remark"`
	DictSort  int    `json:"dictSort"`  // 字典排序
	DictLabel string `json:"dictLabel"` // 字典标签
	DictValue string `json:"dictValue"` // 字典键值
	DictType  string `json:"dictType"`  // 字典类型
	CSSClass  string `json:"cssClass"`  // 样式属性（其他样式扩展）
	ListClass string `json:"listClass"` // 表格字典样式
	IsDefault string `json:"isDefault"` // 是否默认（Y是 N否）
}

type SysDictData struct {
	SysDictDataOpt
	CreatedAt jtime `json:"createTime" gorm:"column:created_time;"`
}

func (o *SysDictData) TableName() string {
	return "t_sysdictdata"
}

// SysDictTypeOpt 操作请求(修改/更新)
type SysDictTypeOpt struct {
	Id       uint64 `json:"dictId" gorm:"primary_key"`
	Status   string `json:"status" gorm:"size:1;default:'0';comment:帐号状态(0正常 1停用);"`
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Remark   string `json:"remark"`
}
type SysDictType struct {
	SysDictTypeOpt
	CreatedAt jtime `json:"createTime" gorm:"column:created_time;"`
}

func (o *SysDictType) TableName() string {
	return "t_sysdicttype"
}
