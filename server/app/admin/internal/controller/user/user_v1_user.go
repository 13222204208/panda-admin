package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/internal/logic/user"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res, err = user.New().GetList(ctx, *req)
	return
}

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	res, err = user.New().Create(ctx, *req)
	return
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = user.New().Update(ctx, *req)
	return
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = user.New().Delete(ctx, *req)
	return
}

func (c *ControllerV1) GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error) {
	res, err = user.New().GetDetail(ctx, *req)
	return
}

func (c *ControllerV1) ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error) {
	err = user.New().ResetPassword(ctx, *req)
	return
}
func (c *ControllerV1) BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error) {
	err = user.New().BatchDelete(ctx, *req)
	return
}

func (c *ControllerV1) GetRoleIds(ctx context.Context, req *v1.GetRoleIdsReq) (res *v1.GetRoleIdsRes, err error) {
	res, err = user.New().GetRoleIds(ctx, *req)
	return
}
func (c *ControllerV1) AssignRoles(ctx context.Context, req *v1.AssignRolesReq) (res *v1.AssignRolesRes, err error) {
	err = user.New().AssignRoles(ctx, *req)
	return
}

func (c *ControllerV1) UploadAvatar(ctx context.Context, req *v1.UploadAvatarReq) (res *v1.UploadAvatarRes, err error) {
	res, err = user.New().UploadAvatar(ctx, *req)
	return
}
