// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure of table orders for DAO operations like Where/Data.
type Orders struct {
	g.Meta      `orm:"table:orders, do:true"`
	Id          interface{} // 订单ID
	OrderNo     interface{} // 订单编号
	UserId      interface{} // 用户ID
	ProductId   interface{} // 商品ID
	ProductName interface{} // 商品名称
	Quantity    interface{} // 购买数量
	Price       interface{} // 商品单价
	TotalAmount interface{} // 订单总金额
	Status      interface{} // 订单状态(0:待支付,1:已支付,2:已发货,3:已完成,4:已取消)
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
