// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/user/v1"
)

type (
	ILogin interface {
		// GetCaptcha 获取验证码
		GetCaptcha(ctx context.Context, req v1.CaptchaReq) (res *v1.CaptchaRes, err error)
	}
	IUser interface {
		// Login 用户登录
		Login(ctx context.Context, in v1.LoginReq) (out *v1.LoginRes, err error)
		// RefreshToken 刷新令牌
		RefreshToken(ctx context.Context, in v1.RefreshTokenReq) (out *v1.RefreshTokenRes, err error)
		// GetUserRoutes 获取用户路由权限
		GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (*v1.GetUserRoutesRes, error)
		// Create 创建用户
		Create(ctx context.Context, in v1.CreateReq) (out *v1.CreateRes, err error)
		// GetList 获取用户列表
		GetList(ctx context.Context, in v1.GetListReq) (out *v1.GetListRes, err error)
		// Update 更新用户
		Update(ctx context.Context, in v1.UpdateReq) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, in v1.DeleteReq) (err error)
		// GetDetail 获取用户详情
		GetDetail(ctx context.Context, in v1.GetDetailReq) (out *v1.GetDetailRes, err error)
		// ResetPassword 重置密码
		ResetPassword(ctx context.Context, in v1.ResetPasswordReq) (err error)
		// BatchDelete 批量删除用户
		BatchDelete(ctx context.Context, in v1.BatchDeleteReq) (err error)
		// GetRoleIds 获取用户对应的角色ID列表
		GetRoleIds(ctx context.Context, in v1.GetRoleIdsReq) (out *v1.GetRoleIdsRes, err error)
		// AssignRoles 分配用户角色
		AssignRoles(ctx context.Context, in v1.AssignRolesReq) (err error)
		// UploadAvatar 上传用户头像
		UploadAvatar(ctx context.Context, req v1.UploadAvatarReq) (*v1.UploadAvatarRes, error)
	}
)

var (
	localLogin ILogin
	localUser  IUser
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
