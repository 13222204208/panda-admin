// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Dict is the golang structure of table dict for DAO operations like Where/Data.
type Dict struct {
	g.Meta    `orm:"table:dict, do:true"`
	Id        interface{} // 主键ID
	Title     interface{} // 字典名称
	DictType  interface{} // 字典类型（如 sex、status、job_type）
	DictLabel interface{} // 字典标签（如 男、启用）
	DictValue interface{} // 字典值（如 1、enabled）
	Sort      interface{} // 排序值（升序）
	Status    interface{} // 状态（1启用，0禁用）
	Remark    interface{} // 备注说明
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
