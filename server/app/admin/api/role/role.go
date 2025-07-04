// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"server/app/admin/api/role/v1"
)

type IRoleV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GetAll(ctx context.Context, req *v1.GetAllReq) (res *v1.GetAllRes, err error)
	AssignMenus(ctx context.Context, req *v1.AssignMenusReq) (res *v1.AssignMenusRes, err error)
	GetRoleMenuIds(ctx context.Context, req *v1.GetRoleMenuIdsReq) (res *v1.GetRoleMenuIdsRes, err error)
}
