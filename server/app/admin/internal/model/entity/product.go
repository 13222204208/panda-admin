// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Product is the golang structure for table product.
type Product struct {
	Id             int64       `json:"id"             orm:"id"              description:"商品ID"`               // 商品ID
	Name           string      `json:"name"           orm:"name"            description:"商品名称"`               // 商品名称
	Description    string      `json:"description"    orm:"description"     description:"商品描述"`               // 商品描述
	Price          float64     `json:"price"          orm:"price"           description:"商品价格"`               // 商品价格
	Stock          int         `json:"stock"          orm:"stock"           description:"库存数量"`               // 库存数量
	MainImage      string      `json:"mainImage"      orm:"main_image"      description:"商品主图URL"`            // 商品主图URL
	CarouselImages string      `json:"carouselImages" orm:"carousel_images" description:"商品轮播图URL数组(JSON格式)"` // 商品轮播图URL数组(JSON格式)
	Attachment     string      `json:"attachment"     orm:"attachment"      description:"附件文件URL"`            // 附件文件URL
	Status         int         `json:"status"         orm:"status"          description:"商品状态(1:上架,0:下架)"`    // 商品状态(1:上架,0:下架)
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`               // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`               // 更新时间
}
