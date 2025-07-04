// Package member Member控制器
package member

import (
	"context"

	v1 "server/app/admin/api/member/v1"
	"server/app/admin/internal/logic/member"
)
// GetMemberList 获取Member列表
func (c *ControllerV1) GetMemberList(ctx context.Context, req *v1.GetMemberListReq) (res *v1.GetMemberListRes, err error) {
	res, err = member.New().GetMemberList(ctx, *req)
	return
}
// CreateMember 创建Member
func (c *ControllerV1) CreateMember(ctx context.Context, req *v1.CreateMemberReq) (res *v1.CreateMemberRes, err error) {
	res, err = member.New().CreateMember(ctx, *req)
	return
}
// UpdateMember 更新Member
func (c *ControllerV1) UpdateMember(ctx context.Context, req *v1.UpdateMemberReq) (res *v1.UpdateMemberRes, err error) {
	res, err = member.New().UpdateMember(ctx, *req)
	return
}
// DeleteMember 删除Member
func (c *ControllerV1) DeleteMember(ctx context.Context, req *v1.DeleteMemberReq) (res *v1.DeleteMemberRes, err error) {
	res, err = member.New().DeleteMember(ctx, *req)
	return
}
// BatchDeleteMember 批量删除Member
func (c *ControllerV1) BatchDeleteMember(ctx context.Context, req *v1.BatchDeleteMemberReq) (res *v1.BatchDeleteMemberRes, err error) {
	res, err = member.New().BatchDeleteMember(ctx, *req)
	return
}