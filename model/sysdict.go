package model

// SysDictDataOpt 操作请求(修改/更新)
type SysDictDataOpt struct {
	ID        uint   `json:"dataId" gorm:"primary_key"`
	Enable    string `json:"enable" gorm:"size:1;default:'0';comment:帐号状态(0正常 1停用);"`
	Label     string `json:"dataLabel"` // 字典标签
	Value     string `json:"dataValue"` // 字典键值
	TypeCode  string `json:"typeCode"`  // 字典类型
	IsDefault string `json:"isDefault"` // 是否默认（Y是 N否）
	CSSClass  string `json:"cssClass"`  // 样式属性（其他样式扩展）
	ListClass string `json:"listClass"` // 表格字典样式
}

type SysDictData struct {
	SysDictDataOpt
	CreatedAt jtime `json:"createdAt"`
}

func (o *SysDictData) TableName() string {
	return "t_sysdictdata"
}

// SysDictTypeOpt 操作请求(修改/更新)
type SysDictTypeOpt struct {
	ID          uint   `json:"typeId" gorm:"primary_key"`
	Enable      string `json:"enable" gorm:"size:1;default:'0';comment:帐号状态(0正常 1停用);"`
	Name        string `json:"typeName"`
	Code        string `json:"typeCode"`
	Description string `json:"description"`
}
type SysDictType struct {
	SysDictTypeOpt
	CreatedAt jtime `json:"createdAt"`
}

func (o *SysDictType) TableName() string {
	return "t_sysdicttype"
}
