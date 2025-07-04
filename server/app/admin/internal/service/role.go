// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/role/v1"
)

type (
	IRole interface {
		// Create 创建角色
		Create(ctx context.Context, in v1.CreateReq) (out *v1.CreateRes, err error)
		// GetList 获取角色列表
		GetList(ctx context.Context, in v1.GetListReq) (out *v1.GetListRes, err error)
		// Update 更新角色
		Update(ctx context.Context, in v1.UpdateReq) (err error)
		// Delete 删除角色
		Delete(ctx context.Context, in v1.DeleteReq) (err error)
		// GetAll 获取所有角色列表（不分页）
		GetAll(ctx context.Context, in v1.GetAllReq) (out *v1.GetAllRes, err error)
		// AssignMenus 分配角色菜单权限
		AssignMenus(ctx context.Context, in v1.AssignMenusReq) (err error)
		// GetRoleMenuIds 获取角色菜单ID列表
		GetRoleMenuIds(ctx context.Context, in v1.GetRoleMenuIdsReq) (out *v1.GetRoleMenuIdsRes, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
