// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package attachment

import (
	"context"

	"server/app/admin/api/attachment/v1"
)

type IAttachmentV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error)
	Download(ctx context.Context, req *v1.DownloadReq) (res *v1.DownloadRes, err error)
}
