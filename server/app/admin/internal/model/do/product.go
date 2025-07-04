// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure of table product for DAO operations like Where/Data.
type Product struct {
	g.Meta    `orm:"table:product, do:true"`
	Id        interface{} // 商品ID
	Name      interface{} // 商品名称
	Price     interface{} // 商品价格
	Stock     interface{} // 商品库存
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
