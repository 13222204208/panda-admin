// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/menu/v1"
)

type (
	IMenu interface {
		// GetList 获取菜单列表（不分页，获取全部菜单）
		GetList(ctx context.Context, req v1.GetListReq) (*v1.GetListRes, error)
		// Create 创建菜单
		Create(ctx context.Context, req v1.CreateReq) (*v1.CreateRes, error)
		// Update 更新菜单
		Update(ctx context.Context, req v1.UpdateReq) (*v1.UpdateRes, error)
		// Delete 删除菜单
		Delete(ctx context.Context, req v1.DeleteReq) (*v1.DeleteRes, error)
		// GetTree 获取菜单树状结构
		GetTree(ctx context.Context, req v1.GetTreeReq) (*v1.GetTreeRes, error)
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
