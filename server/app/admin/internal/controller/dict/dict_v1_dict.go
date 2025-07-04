package dict

import (
	"context"

	v1 "server/app/admin/api/dict/v1"
	"server/app/admin/internal/logic/dict"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res, err = dict.New().GetList(ctx, *req)
	return
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	res, err = dict.New().Update(ctx, *req)
	return
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	res, err = dict.New().Delete(ctx, *req)
	return
}

func (c *ControllerV1) BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error) {
	res, err = dict.New().BatchDelete(ctx, *req)
	return
}

func (c *ControllerV1) BatchCreate(ctx context.Context, req *v1.BatchCreateReq) (res *v1.BatchCreateRes, err error) {
	res, err = dict.New().BatchCreate(ctx, *req)
	return
}

func (c *ControllerV1) GetOptions(ctx context.Context, req *v1.GetOptionsReq) (res *v1.GetOptionsRes, err error) {
	res, err = dict.New().GetOptions(ctx, *req)
	return
}
func (c *ControllerV1) GetDistinctTypes(ctx context.Context, req *v1.GetDistinctTypesReq) (res *v1.GetDistinctTypesRes, err error) {
	return dict.New().GetDistinctTypes(ctx, *req)
}
