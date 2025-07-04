// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure for table member.
type Member struct {
	Id        int64       `json:"id"        orm:"id"         description:"会员ID"` // 会员ID
	Username  string      `json:"username"  orm:"username"   description:"用户名"`  // 用户名
	Email     string      `json:"email"     orm:"email"      description:"邮箱"`   // 邮箱
	Mobile    string      `json:"mobile"    orm:"mobile"     description:"手机号"`  // 手机号
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
}
