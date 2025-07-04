// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Dict is the golang structure for table dict.
type Dict struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键ID"`                        // 主键ID
	Title     string      `json:"title"     orm:"title"      description:"字典名称"`                        // 字典名称
	DictType  string      `json:"dictType"  orm:"dict_type"  description:"字典类型（如 sex、status、job_type）"` // 字典类型（如 sex、status、job_type）
	DictLabel string      `json:"dictLabel" orm:"dict_label" description:"字典标签（如 男、启用）"`                // 字典标签（如 男、启用）
	DictValue string      `json:"dictValue" orm:"dict_value" description:"字典值（如 1、enabled）"`            // 字典值（如 1、enabled）
	Sort      int         `json:"sort"      orm:"sort"       description:"排序值（升序）"`                     // 排序值（升序）
	Status    int         `json:"status"    orm:"status"     description:"状态（1启用，0禁用）"`                 // 状态（1启用，0禁用）
	Remark    string      `json:"remark"    orm:"remark"     description:"备注说明"`                        // 备注说明
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                        // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                        // 更新时间
}
