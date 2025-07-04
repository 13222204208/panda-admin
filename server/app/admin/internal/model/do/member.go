// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure of table member for DAO operations like Where/Data.
type Member struct {
	g.Meta    `orm:"table:member, do:true"`
	Id        interface{} // 会员ID
	Username  interface{} // 用户名
	Email     interface{} // 邮箱
	Mobile    interface{} // 手机号
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
