// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"server/app/admin/api/user/v1"
)

type IUserV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
	Captcha(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error)
	GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (res *v1.GetUserRoutesRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error)
	BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error)
	GetRoleIds(ctx context.Context, req *v1.GetRoleIdsReq) (res *v1.GetRoleIdsRes, err error)
	AssignRoles(ctx context.Context, req *v1.AssignRolesReq) (res *v1.AssignRolesRes, err error)
	UploadAvatar(ctx context.Context, req *v1.UploadAvatarReq) (res *v1.UploadAvatarRes, err error)
}
