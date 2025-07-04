// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleMenu is the golang structure of table role_menu for DAO operations like Where/Data.
type RoleMenu struct {
	g.Meta    `orm:"table:role_menu, do:true"`
	Id        interface{} // 主键ID
	RoleId    interface{} // 角色ID
	MenuId    interface{} // 菜单ID
	CreatedAt *gtime.Time // 创建时间
}
