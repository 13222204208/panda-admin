// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dict

import (
	"context"

	"server/app/admin/api/dict/v1"
)

type IDictV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error)
	BatchCreate(ctx context.Context, req *v1.BatchCreateReq) (res *v1.BatchCreateRes, err error)
	GetOptions(ctx context.Context, req *v1.GetOptionsReq) (res *v1.GetOptionsRes, err error)
	GetDistinctTypes(ctx context.Context, req *v1.GetDistinctTypesReq) (res *v1.GetDistinctTypesRes, err error)
}
