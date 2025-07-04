// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Department is the golang structure of table department for DAO operations like Where/Data.
type Department struct {
	g.Meta    `orm:"table:department, do:true"`
	Id        interface{} // 主键ID
	ParentId  interface{} // 父级部门ID
	Name      interface{} // 部门名称
	Principal interface{} // 负责人名称
	Phone     interface{} // 联系电话
	Email     interface{} // 邮箱地址
	Sort      interface{} // 排序号
	Status    interface{} // 状态（1启用，0禁用）
	Remark    interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
