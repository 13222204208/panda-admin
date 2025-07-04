// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Test is the golang structure for table test.
type Test struct {
	Id        int64       `json:"id"        orm:"id"         description:"用户ID，主键"`       // 用户ID，主键
	Username  string      `json:"username"  orm:"username"   description:"用户名"`           // 用户名
	Email     string      `json:"email"     orm:"email"      description:"电子邮箱"`          // 电子邮箱
	Mobile    string      `json:"mobile"    orm:"mobile"     description:"手机号码"`          // 手机号码
	Password  string      `json:"password"  orm:"password"   description:"密码(加密存储)"`      // 密码(加密存储)
	Status    int         `json:"status"    orm:"status"     description:"状态(0-禁用,1-正常)"` // 状态(0-禁用,1-正常)
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`          // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`          // 更新时间
}
