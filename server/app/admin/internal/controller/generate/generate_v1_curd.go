package generate

import (
	"context"

	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/logic/generate"
)

func (c *ControllerV1) CodeGenRecord(ctx context.Context, req *v1.CodeGenRecordReq) (res *v1.CodeGenRecordRes, err error) {
	return generate.New().CodeGenRecord(ctx, *req)
}
