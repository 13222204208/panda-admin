package role

import (
	"context"

	v1 "server/app/admin/api/role/v1"
	"server/app/admin/internal/logic/role"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res, err = role.New().GetList(ctx, *req)
	return
}

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	res, err = role.New().Create(ctx, *req)
	return
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = role.New().Update(ctx, *req)
	return
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = role.New().Delete(ctx, *req)
	return
}
func (c *ControllerV1) GetAll(ctx context.Context, req *v1.GetAllReq) (res *v1.GetAllRes, err error) {
	res, err = role.New().GetAll(ctx, *req)
	return
}
func (c *ControllerV1) AssignMenus(ctx context.Context, req *v1.AssignMenusReq) (res *v1.AssignMenusRes, err error) {
	err = role.New().AssignMenus(ctx, *req)
	return
}
func (c *ControllerV1) GetRoleMenuIds(ctx context.Context, req *v1.GetRoleMenuIdsReq) (res *v1.GetRoleMenuIdsRes, err error) {
	res, err = role.New().GetRoleMenuIds(ctx, *req)
	return
}
