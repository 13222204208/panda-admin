package v1

import "github.com/gogf/gf/v2/frame/g"

type CodeGenRecordReq struct {
	g.Meta       `path:"/generate/record/{id}" method:"post" tags:"代码生成" summary:"执行代码生成"`
	Id           uint64 `json:"id" v:"required#请输入记录ID" dc:"记录ID"`
	TableName    string `json:"tableName" v:"required#请输入表名称" dc:"数据表名称"`
	TableComment string `json:"tableComment" dc:"表注释"`
	PackageName  string `json:"packageName" v:"required#请输入包名" dc:"生成的Go包名"`
	ModuleName   string `json:"moduleName" v:"required#请输入模块名" dc:"模块名"`
	Options      string `json:"options" dc:"配置选项"`
	Columns      string `json:"columns" dc:"表字段"`
}

// UpdateCodeGenRecordRes 更新代码生成记录响应
type CodeGenRecordRes struct{}
