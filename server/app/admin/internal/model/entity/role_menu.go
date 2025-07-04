// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleMenu is the golang structure for table role_menu.
type RoleMenu struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键ID"` // 主键ID
	RoleId    uint64      `json:"roleId"    orm:"role_id"    description:"角色ID"` // 角色ID
	MenuId    uint64      `json:"menuId"    orm:"menu_id"    description:"菜单ID"` // 菜单ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
}
