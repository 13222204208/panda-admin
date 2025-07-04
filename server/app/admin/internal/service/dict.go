// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/dict/v1"
)

type (
	IDict interface {
		// GetList 获取字典列表
		GetList(ctx context.Context, in v1.GetListReq) (out *v1.GetListRes, err error)
		// Update 更新字典
		Update(ctx context.Context, in v1.UpdateReq) (out *v1.UpdateRes, err error)
		// Delete 删除字典
		Delete(ctx context.Context, in v1.DeleteReq) (out *v1.DeleteRes, err error)
		// BatchDelete 批量删除字典
		BatchDelete(ctx context.Context, in v1.BatchDeleteReq) (out *v1.BatchDeleteRes, err error)
		// BatchCreate 批量创建字典
		BatchCreate(ctx context.Context, in v1.BatchCreateReq) (out *v1.BatchCreateRes, err error)
		// GetOptions 根据字典类型获取字典选项
		GetOptions(ctx context.Context, in v1.GetOptionsReq) (out *v1.GetOptionsRes, err error)
		// GetDistinctTypes 获取不重复的字典类型和标题
		GetDistinctTypes(ctx context.Context, in v1.GetDistinctTypesReq) (out *v1.GetDistinctTypesRes, err error)
	}
)

var (
	localDict IDict
)

func Dict() IDict {
	if localDict == nil {
		panic("implement not found for interface IDict, forgot register?")
	}
	return localDict
}

func RegisterDict(i IDict) {
	localDict = i
}
