// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeGenRecord is the golang structure for table code_gen_record.
type CodeGenRecord struct {
	Id           uint64      `json:"id"           orm:"id"            description:"主键ID"`                 // 主键ID
	TableName    string      `json:"tableName"    orm:"table_name"    description:"数据表名称"`                // 数据表名称
	TableComment string      `json:"tableComment" orm:"table_comment" description:"表注释"`                  // 表注释
	PackageName  string      `json:"packageName"  orm:"package_name"  description:"生成的Go包名"`              // 生成的Go包名
	ModuleName   string      `json:"moduleName"   orm:"module_name"   description:"模块名（例如 system、user）"`  // 模块名（例如 system、user）
	Options      string      `json:"options"      orm:"options"       description:"配置选项"`                 // 配置选项
	Columns      string      `json:"columns"      orm:"columns"       description:"表字段"`                  // 表字段
	Status       int         `json:"status"       orm:"status"        description:"生成状态（1成功，0失败）代码是否已生成"` // 生成状态（1成功，0失败）代码是否已生成
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"生成时间"`                 // 生成时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`                 // 更新时间
}
