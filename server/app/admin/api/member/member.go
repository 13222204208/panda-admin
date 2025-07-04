// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package member

import (
	"context"

	"server/app/admin/api/member/v1"
)

type IMemberV1 interface {
	CreateMember(ctx context.Context, req *v1.CreateMemberReq) (res *v1.CreateMemberRes, err error)
	UpdateMember(ctx context.Context, req *v1.UpdateMemberReq) (res *v1.UpdateMemberRes, err error)
	DeleteMember(ctx context.Context, req *v1.DeleteMemberReq) (res *v1.DeleteMemberRes, err error)
	BatchDeleteMember(ctx context.Context, req *v1.BatchDeleteMemberReq) (res *v1.BatchDeleteMemberRes, err error)
	GetMemberList(ctx context.Context, req *v1.GetMemberListReq) (res *v1.GetMemberListRes, err error)
}
