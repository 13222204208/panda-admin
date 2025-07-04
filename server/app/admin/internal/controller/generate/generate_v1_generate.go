package generate

import (
	"context"

	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/logic/generate"
)

func (c *ControllerV1) GetCodeGenRecordList(ctx context.Context, req *v1.GetCodeGenRecordListReq) (res *v1.GetCodeGenRecordListRes, err error) {
	return generate.New().GetCodeGenRecordList(ctx, *req)
}
func (c *ControllerV1) GetCodeGenRecordDetail(ctx context.Context, req *v1.GetCodeGenRecordDetailReq) (res *v1.GetCodeGenRecordDetailRes, err error) {
	return generate.New().GetCodeGenRecordDetail(ctx, *req)
}
func (c *ControllerV1) DeleteCodeGenRecord(ctx context.Context, req *v1.DeleteCodeGenRecordReq) (res *v1.DeleteCodeGenRecordRes, err error) {
	return generate.New().DeleteCodeGenRecord(ctx, *req)
}
func (c *ControllerV1) UpdateCodeGenRecord(ctx context.Context, req *v1.UpdateCodeGenRecordReq) (res *v1.UpdateCodeGenRecordRes, err error) {
	return generate.New().UpdateCodeGenRecord(ctx, *req)
}
