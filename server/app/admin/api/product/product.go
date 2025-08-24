// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package product

import (
	"context"

	"server/app/admin/api/product/v1"
)

type IProductV1 interface {
	CreateProduct(ctx context.Context, req *v1.CreateProductReq) (res *v1.CreateProductRes, err error)
	UpdateProduct(ctx context.Context, req *v1.UpdateProductReq) (res *v1.UpdateProductRes, err error)
	DeleteProduct(ctx context.Context, req *v1.DeleteProductReq) (res *v1.DeleteProductRes, err error)
	BatchDeleteProduct(ctx context.Context, req *v1.BatchDeleteProductReq) (res *v1.BatchDeleteProductRes, err error)
	GetProductList(ctx context.Context, req *v1.GetProductListReq) (res *v1.GetProductListRes, err error)
}
