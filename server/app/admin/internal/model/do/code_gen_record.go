// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeGenRecord is the golang structure of table code_gen_record for DAO operations like Where/Data.
type CodeGenRecord struct {
	g.Meta       `orm:"table:code_gen_record, do:true"`
	Id           interface{} // 主键ID
	TableName    interface{} // 数据表名称
	TableComment interface{} // 表注释
	PackageName  interface{} // 生成的Go包名
	ModuleName   interface{} // 模块名（例如 system、user）
	Options      interface{} // 配置选项
	Columns      interface{} // 表字段
	Status       interface{} // 生成状态（1成功，0失败）代码是否已生成
	CreatedAt    *gtime.Time // 生成时间
	UpdatedAt    *gtime.Time // 更新时间
}
