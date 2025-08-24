package product

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"server/app/admin/api/common/page"
	v1 "server/app/admin/api/product/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/model/entity"
)

type sProduct struct{}

func New() *sProduct {
	return &sProduct{}
}


// GetProductList 获取商品信息表列表
func (s *sProduct) GetProductList(ctx context.Context, in v1.GetProductListReq) (out *v1.GetProductListRes, err error) {
	out = &v1.GetProductListRes{}

	m := dao.Product.Ctx(ctx)

	// 构建查询条件
	if in.Name != "" {
		m = m.WhereLike(dao.Product.Columns().Name, "%"+in.Name+"%")
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询商品信息表总数失败")
	}

	// 分页查询
	// 初始化为空切片，确保返回空数组而不是null
	list := make([]*entity.Product, 0)
	err = m.Page(in.CurrentPage, in.PageSize).
		OrderDesc(dao.Product.Columns().CreatedAt).
		Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询商品信息表列表失败")
	}

	out.ResPage = page.ResPage{
		Total:       int(total),
		CurrentPage: in.CurrentPage,
	}
	out.List = list
	return
}



// CreateProduct 创建商品信息表
func (s *sProduct) CreateProduct(ctx context.Context, in v1.CreateProductReq) (out *v1.CreateProductRes, err error) {
	out = &v1.CreateProductRes{}

	// 构建插入数据
	data := g.Map{}
	if in.Name != nil {
		data[dao.Product.Columns().Name] = *in.Name
	}
	if in.Description != nil {
		data[dao.Product.Columns().Description] = *in.Description
	}
	if in.Price != nil {
		data[dao.Product.Columns().Price] = *in.Price
	}
	if in.Stock != nil {
		data[dao.Product.Columns().Stock] = *in.Stock
	}
	if in.MainImage != nil {
		data[dao.Product.Columns().MainImage] = *in.MainImage
	}
	if in.CarouselImages != nil {
		data[dao.Product.Columns().CarouselImages] = *in.CarouselImages
	}
	if in.Attachment != nil {
		data[dao.Product.Columns().Attachment] = *in.Attachment
	}
	if in.Status != nil {
		data[dao.Product.Columns().Status] = *in.Status
	}

	// 插入数据
	_, err = dao.Product.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建商品信息表失败")
	}
	return
}



// UpdateProduct 更新商品信息表
func (s *sProduct) UpdateProduct(ctx context.Context, in v1.UpdateProductReq) (out *v1.UpdateProductRes, err error) {
	out = &v1.UpdateProductRes{}

	// 检查商品信息表是否存在
	count, err := dao.Product.Ctx(ctx).Where(dao.Product.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询商品信息表失败")
	}
	if count == 0 {
		return nil, gerror.New("商品信息表不存在")
	}

	// 动态构建更新数据
	updateData := g.Map{}
	if in.Name != nil {
		updateData[dao.Product.Columns().Name] = *in.Name
	}
	if in.Description != nil {
		updateData[dao.Product.Columns().Description] = *in.Description
	}
	if in.Price != nil {
		updateData[dao.Product.Columns().Price] = *in.Price
	}
	if in.Stock != nil {
		updateData[dao.Product.Columns().Stock] = *in.Stock
	}
	if in.MainImage != nil {
		updateData[dao.Product.Columns().MainImage] = *in.MainImage
	}
	if in.CarouselImages != nil {
		updateData[dao.Product.Columns().CarouselImages] = *in.CarouselImages
	}
	if in.Attachment != nil {
		updateData[dao.Product.Columns().Attachment] = *in.Attachment
	}
	if in.Status != nil {
		updateData[dao.Product.Columns().Status] = *in.Status
	}

	// 更新数据
	_, err = dao.Product.Ctx(ctx).
		Where(dao.Product.Columns().Id, in.Id).
		Data(updateData).
		Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新商品信息表失败")
	}

	return
}



// DeleteProduct 删除商品信息表
func (s *sProduct) DeleteProduct(ctx context.Context, in v1.DeleteProductReq) (out *v1.DeleteProductRes, err error) {
	out = &v1.DeleteProductRes{}

	// 检查商品信息表是否存在
	count, err := dao.Product.Ctx(ctx).Where(dao.Product.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询商品信息表失败")
	}
	if count == 0 {
		return nil, gerror.New("商品信息表不存在")
	}

	// 删除数据
	_, err = dao.Product.Ctx(ctx).Where(dao.Product.Columns().Id, in.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除商品信息表失败")
	}

	return
}



// BatchDeleteProduct 批量删除商品信息表
func (s *sProduct) BatchDeleteProduct(ctx context.Context, in v1.BatchDeleteProductReq) (out *v1.BatchDeleteProductRes, err error) {
	out = &v1.BatchDeleteProductRes{}

	if len(in.Ids) == 0 {
		return nil, gerror.New("请选择要删除的商品信息表")
	}

	// 批量删除
	_, err = dao.Product.Ctx(ctx).WhereIn(dao.Product.Columns().Id, in.Ids).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除商品信息表失败")
	}

	return
}
