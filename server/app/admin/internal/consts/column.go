package consts

// ColumnConfig 字段配置信息
type ColumnConfig struct {
	IsEdit          bool   `json:"isEdit" dc:"是否可编辑"`
	IsList          bool   `json:"isList" dc:"是否在列表显示"`
	IsQuery         bool   `json:"isQuery" dc:"是否查询字段"`
	DictType        string `json:"dictType" dc:"字典类型"`
	HtmlType        string `json:"htmlType" dc:"HTML类型"`
	IsUnique        bool   `json:"isUnique" dc:"是否唯一"`
	FieldName       string `json:"fieldName" dc:"字段名称"`
	QueryType       string `json:"queryType" dc:"查询类型"`
	IsRequired      bool   `json:"isRequired" dc:"是否必填"`
	FieldComment    string `json:"fieldComment" dc:"字段注释"`
	ValidationRules string `json:"validationRules" dc:"验证规则"`
}
