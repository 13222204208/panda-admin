// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:role, do:true"`
	Id        interface{} // 主键ID
	Name      interface{} // 角色名称
	Code      interface{} // 角色编码（唯一）
	Status    interface{} // 状态（1启用，0禁用）
	Remark    interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
