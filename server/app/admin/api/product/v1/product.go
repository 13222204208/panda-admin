package v1

import (
	"server/app/admin/api/common/page"
	"server/app/admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductCommon 商品信息表公共字段
type ProductCommon struct {
	Name *string `json:"name,omitempty"  dc:"商品名称"`
	Description *string `json:"description,omitempty"  dc:"商品描述"`
	Price *string `json:"price,omitempty"  dc:"商品价格"`
	Stock *string `json:"stock,omitempty"  dc:"库存数量"`
	MainImage *string `json:"main_image,omitempty"  dc:"商品主图URL"`
	CarouselImages *string `json:"carousel_images,omitempty"  dc:"商品轮播图URL数组(JSON格式)"`
	Attachment *string `json:"attachment,omitempty"  dc:"附件文件URL"`
	Status *string `json:"status,omitempty"  dc:"商品状态(1:上架,0:下架)"`
}


// CreateProductReq 创建商品信息表请求
type CreateProductReq struct {
    g.Meta `path:"/product" method:"post" tags:"商品信息表" summary:"创建商品信息表"`
    ProductCommon
}

// CreateProductRes 创建商品信息表响应
type CreateProductRes struct {}



// UpdateProductReq 更新商品信息表请求
type UpdateProductReq struct {
    g.Meta `path:"/product/{id}" method:"put" tags:"商品信息表" summary:"更新商品信息表"`
    Id uint64 `json:"id" v:"required#请输入ID" dc:"ID"`
    ProductCommon
}

// UpdateProductRes 更新商品信息表响应
type UpdateProductRes struct {}



// DeleteProductReq 删除商品信息表请求
type DeleteProductReq struct {
    g.Meta `path:"/product/{id}" method:"delete" tags:"商品信息表" summary:"删除商品信息表"`
    Id uint64 `json:"id" v:"required#请输入ID" dc:"ID"`
}

// DeleteProductRes 删除商品信息表响应
type DeleteProductRes struct {}



// BatchDeleteProductReq 批量删除商品信息表请求
type BatchDeleteProductReq struct {
    g.Meta `path:"/product/batch" method:"delete" tags:"商品信息表" summary:"批量删除商品信息表"`
    Ids []uint64 `json:"ids" v:"required#请输入ID列表" dc:"ID列表"`
}

// BatchDeleteProductRes 批量删除商品信息表响应
type BatchDeleteProductRes struct {}



// GetProductListReq 获取商品信息表列表请求
type GetProductListReq struct {
    g.Meta `path:"/product" method:"get" tags:"商品信息表" summary:"获取商品信息表列表"`
    page.ReqPage
    Name string `json:"name" dc:"商品名称"`
}

// GetProductListRes 获取商品信息表列表响应
type GetProductListRes struct {
    List []*entity.Product `json:"list" dc:"商品信息表列表"`
    page.ResPage
}


