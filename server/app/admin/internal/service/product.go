// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/product/v1"
)

type (
	IProduct interface {
		// GetProductList 获取商品信息表列表
		GetProductList(ctx context.Context, in v1.GetProductListReq) (out *v1.GetProductListRes, err error)
		// CreateProduct 创建商品信息表
		CreateProduct(ctx context.Context, in v1.CreateProductReq) (out *v1.CreateProductRes, err error)
		// UpdateProduct 更新商品信息表
		UpdateProduct(ctx context.Context, in v1.UpdateProductReq) (out *v1.UpdateProductRes, err error)
		// DeleteProduct 删除商品信息表
		DeleteProduct(ctx context.Context, in v1.DeleteProductReq) (out *v1.DeleteProductRes, err error)
		// BatchDeleteProduct 批量删除商品信息表
		BatchDeleteProduct(ctx context.Context, in v1.BatchDeleteProductReq) (out *v1.BatchDeleteProductRes, err error)
	}
)

var (
	localProduct IProduct
)

func Product() IProduct {
	if localProduct == nil {
		panic("implement not found for interface IProduct, forgot register?")
	}
	return localProduct
}

func RegisterProduct(i IProduct) {
	localProduct = i
}
