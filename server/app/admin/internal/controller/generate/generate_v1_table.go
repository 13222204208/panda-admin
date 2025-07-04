package generate

import (
	"context"

	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/logic/generate"
)

func (c *ControllerV1) GetTables(ctx context.Context, req *v1.GetTablesReq) (res *v1.GetTablesRes, err error) {
	return generate.New().GetTables(ctx, *req)
}
func (c *ControllerV1) ImportTables(ctx context.Context, req *v1.ImportTablesReq) (res *v1.ImportTablesRes, err error) {
	return generate.New().ImportTables(ctx, *req)
}
func (c *ControllerV1) GetTablesWithColumns(ctx context.Context, req *v1.GetTablesWithColumnsReq) (res *v1.GetTablesWithColumnsRes, err error) {
	return generate.New().GetTablesWithColumns(ctx, *req)
}
func (c *ControllerV1) GetTableColumns(ctx context.Context, req *v1.GetTableColumnsReq) (res *v1.GetTableColumnsRes, err error) {
	return generate.New().GetTableColumns(ctx, *req)
}
