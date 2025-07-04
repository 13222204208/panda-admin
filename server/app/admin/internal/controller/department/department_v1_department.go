package department

import (
	"context"

	v1 "server/app/admin/api/department/v1"
	"server/app/admin/internal/logic/department"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res, err = department.New().GetList(ctx, *req)
	return
}

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	res, err = department.New().Create(ctx, *req)
	return
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = department.New().Update(ctx, *req)
	return
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = department.New().Delete(ctx, *req)
	return
}
