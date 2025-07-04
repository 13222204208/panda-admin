package generate

import (
	"context"

	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/logic/generate"
)

func (c *ControllerV1) GenerateSql(ctx context.Context, req *v1.GenerateSqlReq) (res *v1.GenerateSqlRes, err error) {
	return generate.New().GenerateSql(ctx, *req)
}
func (c *ControllerV1) ExecuteSql(ctx context.Context, req *v1.ExecuteSqlReq) (res *v1.ExecuteSqlRes, err error) {
	return generate.New().ExecuteSql(ctx, *req)
}
