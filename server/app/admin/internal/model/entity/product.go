// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure for table product.
type Product struct {
	Id        int64       `json:"id"        orm:"id"         description:"商品ID"` // 商品ID
	Name      string      `json:"name"      orm:"name"       description:"商品名称"` // 商品名称
	Price     float64     `json:"price"     orm:"price"      description:"商品价格"` // 商品价格
	Stock     int         `json:"stock"     orm:"stock"      description:"商品库存"` // 商品库存
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
}
