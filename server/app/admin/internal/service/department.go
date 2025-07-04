// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/department/v1"
)

type (
	IDepartment interface {
		// Create 创建部门
		Create(ctx context.Context, in v1.CreateReq) (out *v1.CreateRes, err error)
		// GetList 获取部门列表
		GetList(ctx context.Context, in v1.GetListReq) (out *v1.GetListRes, err error)
		// Update 更新部门
		Update(ctx context.Context, in v1.UpdateReq) (err error)
		// Delete 删除部门
		Delete(ctx context.Context, in v1.DeleteReq) (err error)
	}
)

var (
	localDepartment IDepartment
)

func Department() IDepartment {
	if localDepartment == nil {
		panic("implement not found for interface IDepartment, forgot register?")
	}
	return localDepartment
}

func RegisterDepartment(i IDepartment) {
	localDepartment = i
}
