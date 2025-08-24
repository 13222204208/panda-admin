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
	g.Meta         `orm:"table:product, do:true"`
	Id             interface{} // 商品ID
	Name           interface{} // 商品名称
	Description    interface{} // 商品描述
	Price          interface{} // 商品价格
	Stock          interface{} // 库存数量
	MainImage      interface{} // 商品主图URL
	CarouselImages interface{} // 商品轮播图URL数组(JSON格式)
	Attachment     interface{} // 附件文件URL
	Status         interface{} // 商品状态(1:上架,0:下架)
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
