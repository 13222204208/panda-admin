// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Test is the golang structure of table test for DAO operations like Where/Data.
type Test struct {
	g.Meta    `orm:"table:test, do:true"`
	Id        interface{} // 用户ID，主键
	Username  interface{} // 用户名
	Email     interface{} // 电子邮箱
	Mobile    interface{} // 手机号码
	Password  interface{} // 密码(加密存储)
	Status    interface{} // 状态(0-禁用,1-正常)
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
