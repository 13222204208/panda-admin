// Package product Product控制器
package product

import (
	"context"

	v1 "server/app/admin/api/product/v1"
	"server/app/admin/internal/logic/product"
)
// GetProductList 获取Product列表
func (c *ControllerV1) GetProductList(ctx context.Context, req *v1.GetProductListReq) (res *v1.GetProductListRes, err error) {
	res, err = product.New().GetProductList(ctx, *req)
	return
}
// CreateProduct 创建Product
func (c *ControllerV1) CreateProduct(ctx context.Context, req *v1.CreateProductReq) (res *v1.CreateProductRes, err error) {
	res, err = product.New().CreateProduct(ctx, *req)
	return
}
// UpdateProduct 更新Product
func (c *ControllerV1) UpdateProduct(ctx context.Context, req *v1.UpdateProductReq) (res *v1.UpdateProductRes, err error) {
	res, err = product.New().UpdateProduct(ctx, *req)
	return
}
// DeleteProduct 删除Product
func (c *ControllerV1) DeleteProduct(ctx context.Context, req *v1.DeleteProductReq) (res *v1.DeleteProductRes, err error) {
	res, err = product.New().DeleteProduct(ctx, *req)
	return
}
// BatchDeleteProduct 批量删除Product
func (c *ControllerV1) BatchDeleteProduct(ctx context.Context, req *v1.BatchDeleteProductReq) (res *v1.BatchDeleteProductRes, err error) {
	res, err = product.New().BatchDeleteProduct(ctx, *req)
	return
}