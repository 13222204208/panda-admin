// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/member/v1"
)

type (
	IMember interface {
		// GetMemberList 获取会员信息表列表
		GetMemberList(ctx context.Context, in v1.GetMemberListReq) (out *v1.GetMemberListRes, err error)
		// CreateMember 创建会员信息表
		CreateMember(ctx context.Context, in v1.CreateMemberReq) (out *v1.CreateMemberRes, err error)
		// UpdateMember 更新会员信息表
		UpdateMember(ctx context.Context, in v1.UpdateMemberReq) (out *v1.UpdateMemberRes, err error)
		// DeleteMember 删除会员信息表
		DeleteMember(ctx context.Context, in v1.DeleteMemberReq) (out *v1.DeleteMemberRes, err error)
		// BatchDeleteMember 批量删除会员信息表
		BatchDeleteMember(ctx context.Context, in v1.BatchDeleteMemberReq) (out *v1.BatchDeleteMemberRes, err error)
	}
)

var (
	localMember IMember
)

func Member() IMember {
	if localMember == nil {
		panic("implement not found for interface IMember, forgot register?")
	}
	return localMember
}

func RegisterMember(i IMember) {
	localMember = i
}
