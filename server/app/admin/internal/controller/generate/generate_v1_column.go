package generate

import (
	"context"

	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/logic/generate"
)

func (c *ControllerV1) GetColumnConfigOptions(ctx context.Context, req *v1.GetColumnConfigOptionsReq) (res *v1.GetColumnConfigOptionsRes, err error) {
	return generate.New().GetColumnConfigOptions(ctx, *req)
}
