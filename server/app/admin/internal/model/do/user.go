// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:user, do:true"`
	Id           interface{} // 主键ID
	Title        interface{} // 职位名称
	DepartmentId interface{} // 所属部门ID
	Nickname     interface{} // 昵称
	Username     interface{} // 用户名
	Password     interface{} // 密码（加密存储）
	Avatar       interface{} // 头像
	Phone        interface{} // 联系电话
	Email        interface{} // 邮箱地址
	Sex          interface{} // 性别（0未知，1男，2女）
	Status       interface{} // 状态（1启用，0禁用）
	Remark       interface{} // 备注
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
