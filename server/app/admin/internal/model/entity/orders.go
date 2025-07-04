// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure for table orders.
type Orders struct {
	Id          uint64      `json:"id"          orm:"id"           description:"订单ID"`                                // 订单ID
	OrderNo     string      `json:"orderNo"     orm:"order_no"     description:"订单编号"`                                // 订单编号
	UserId      uint64      `json:"userId"      orm:"user_id"      description:"用户ID"`                                // 用户ID
	ProductId   uint64      `json:"productId"   orm:"product_id"   description:"商品ID"`                                // 商品ID
	ProductName string      `json:"productName" orm:"product_name" description:"商品名称"`                                // 商品名称
	Quantity    uint        `json:"quantity"    orm:"quantity"     description:"购买数量"`                                // 购买数量
	Price       float64     `json:"price"       orm:"price"        description:"商品单价"`                                // 商品单价
	TotalAmount float64     `json:"totalAmount" orm:"total_amount" description:"订单总金额"`                               // 订单总金额
	Status      uint        `json:"status"      orm:"status"       description:"订单状态(0:待支付,1:已支付,2:已发货,3:已完成,4:已取消)"` // 订单状态(0:待支付,1:已支付,2:已发货,3:已完成,4:已取消)
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:"创建时间"`                                // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"  description:"更新时间"`                                // 更新时间
}
