package model

type SysLoginLog struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	BusinessType   string `json:"businessType"`
	RequestMethod  string `json:"requestMethod"`
	Method         string `json:"method"`
	OperateURL     string `json:"operateUrl"`
	OperateAddress string `json:"operateAddress"`
	RequestBody    string `json:"requestBody"`
	Success        bool   `json:"success"`
	LoggingType    string `json:"loggingType"`
	ErrorMsg       string `json:"errorMsg"`
	SystemOs       string `json:"systemOs"`
	CreateTime     string `json:"createTime"`
	OperateName    string `json:"operateName"`
	Browser        string `json:"browser"`
}
