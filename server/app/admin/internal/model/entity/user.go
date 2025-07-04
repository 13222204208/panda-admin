// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id           uint64      `json:"id"           orm:"id"            description:"主键ID"`          // 主键ID
	Title        string      `json:"title"        orm:"title"         description:"职位名称"`          // 职位名称
	DepartmentId uint64      `json:"departmentId" orm:"department_id" description:"所属部门ID"`        // 所属部门ID
	Nickname     string      `json:"nickname"     orm:"nickname"      description:"昵称"`            // 昵称
	Username     string      `json:"username"     orm:"username"      description:"用户名"`           // 用户名
	Password     string      `json:"password"     orm:"password"      description:"密码（加密存储）"`      // 密码（加密存储）
	Avatar       string      `json:"avatar"       orm:"avatar"        description:"头像"`            // 头像
	Phone        string      `json:"phone"        orm:"phone"         description:"联系电话"`          // 联系电话
	Email        string      `json:"email"        orm:"email"         description:"邮箱地址"`          // 邮箱地址
	Sex          int         `json:"sex"          orm:"sex"           description:"性别（0未知，1男，2女）"` // 性别（0未知，1男，2女）
	Status       int         `json:"status"       orm:"status"        description:"状态（1启用，0禁用）"`   // 状态（1启用，0禁用）
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`            // 备注
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`          // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`          // 更新时间
}
