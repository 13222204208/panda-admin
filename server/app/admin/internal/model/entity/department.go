// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Department is the golang structure for table department.
type Department struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键ID"`        // 主键ID
	ParentId  uint64      `json:"parentId"  orm:"parent_id"  description:"父级部门ID"`      // 父级部门ID
	Name      string      `json:"name"      orm:"name"       description:"部门名称"`        // 部门名称
	Principal string      `json:"principal" orm:"principal"  description:"负责人名称"`       // 负责人名称
	Phone     string      `json:"phone"     orm:"phone"      description:"联系电话"`        // 联系电话
	Email     string      `json:"email"     orm:"email"      description:"邮箱地址"`        // 邮箱地址
	Sort      int         `json:"sort"      orm:"sort"       description:"排序号"`         // 排序号
	Status    int         `json:"status"    orm:"status"     description:"状态（1启用，0禁用）"` // 状态（1启用，0禁用）
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`          // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`        // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`        // 更新时间
}
