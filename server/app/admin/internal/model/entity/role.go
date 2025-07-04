// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键ID"`        // 主键ID
	Name      string      `json:"name"      orm:"name"       description:"角色名称"`        // 角色名称
	Code      string      `json:"code"      orm:"code"       description:"角色编码（唯一）"`    // 角色编码（唯一）
	Status    int         `json:"status"    orm:"status"     description:"状态（1启用，0禁用）"` // 状态（1启用，0禁用）
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`          // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`        // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`        // 更新时间
}
