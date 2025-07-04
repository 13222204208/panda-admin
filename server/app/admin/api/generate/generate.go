// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package generate

import (
	"context"

	"server/app/admin/api/generate/v1"
)

type IGenerateV1 interface {
	GetColumnConfigOptions(ctx context.Context, req *v1.GetColumnConfigOptionsReq) (res *v1.GetColumnConfigOptionsRes, err error)
	CodeGenRecord(ctx context.Context, req *v1.CodeGenRecordReq) (res *v1.CodeGenRecordRes, err error)
	GetCodeGenRecordList(ctx context.Context, req *v1.GetCodeGenRecordListReq) (res *v1.GetCodeGenRecordListRes, err error)
	GetCodeGenRecordDetail(ctx context.Context, req *v1.GetCodeGenRecordDetailReq) (res *v1.GetCodeGenRecordDetailRes, err error)
	DeleteCodeGenRecord(ctx context.Context, req *v1.DeleteCodeGenRecordReq) (res *v1.DeleteCodeGenRecordRes, err error)
	UpdateCodeGenRecord(ctx context.Context, req *v1.UpdateCodeGenRecordReq) (res *v1.UpdateCodeGenRecordRes, err error)
	GenerateSql(ctx context.Context, req *v1.GenerateSqlReq) (res *v1.GenerateSqlRes, err error)
	ExecuteSql(ctx context.Context, req *v1.ExecuteSqlReq) (res *v1.ExecuteSqlRes, err error)
	GetTables(ctx context.Context, req *v1.GetTablesReq) (res *v1.GetTablesRes, err error)
	ImportTables(ctx context.Context, req *v1.ImportTablesReq) (res *v1.ImportTablesRes, err error)
	GetTablesWithColumns(ctx context.Context, req *v1.GetTablesWithColumnsReq) (res *v1.GetTablesWithColumnsRes, err error)
	GetTableColumns(ctx context.Context, req *v1.GetTableColumnsReq) (res *v1.GetTableColumnsRes, err error)
}
