package attachment

import (
	"context"

	v1 "server/app/admin/api/attachment/v1"
	"server/app/admin/internal/logic/attachment"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	return attachment.New().GetList(ctx, req)
}
func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	return attachment.New().Upload(ctx, req)
}
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return attachment.New().Update(ctx, req)
}
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return attachment.New().Delete(ctx, req)
}
func (c *ControllerV1) BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error) {
	return attachment.New().BatchDelete(ctx, req)
}
func (c *ControllerV1) Download(ctx context.Context, req *v1.DownloadReq) (res *v1.DownloadRes, err error) {
	return attachment.New().Download(ctx, req)
}
