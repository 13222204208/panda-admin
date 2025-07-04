// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/app/admin/api/generate/v1"
)

type (
	IGenerate interface {
		// GetColumnConfigOptions 获取字段配置选项
		GetColumnConfigOptions(ctx context.Context, req v1.GetColumnConfigOptionsReq) (res *v1.GetColumnConfigOptionsRes, err error)
		// 代码生成
		// 代码生成 - 同时生成后端和前端
		CodeGenRecord(ctx context.Context, req v1.CodeGenRecordReq) (res *v1.CodeGenRecordRes, err error)
		// 批量生成多个模块的后端和前端代码
		BatchGenerateFullStack(ctx context.Context, requests []v1.CodeGenRecordReq) error
		// GetCodeGenRecordList 获取代码生成记录列表
		GetCodeGenRecordList(ctx context.Context, req v1.GetCodeGenRecordListReq) (res *v1.GetCodeGenRecordListRes, err error)
		// GetCodeGenRecordDetail 获取代码生成记录详情
		GetCodeGenRecordDetail(ctx context.Context, req v1.GetCodeGenRecordDetailReq) (res *v1.GetCodeGenRecordDetailRes, err error)
		// DeleteCodeGenRecord 删除代码生成记录
		DeleteCodeGenRecord(ctx context.Context, req v1.DeleteCodeGenRecordReq) (res *v1.DeleteCodeGenRecordRes, err error)
		// UpdateCodeGenRecord 更新代码生成记录
		UpdateCodeGenRecord(ctx context.Context, req v1.UpdateCodeGenRecordReq) (res *v1.UpdateCodeGenRecordRes, err error)
		// GenerateSql 根据提示词生成SQL语句
		GenerateSql(ctx context.Context, req v1.GenerateSqlReq) (res *v1.GenerateSqlRes, err error)
		// ExecuteSql 执行SQL语句（仅允许创建表语句）
		ExecuteSql(ctx context.Context, req v1.ExecuteSqlReq) (res *v1.ExecuteSqlRes, err error)
		// GetTables 获取数据库中所有表信息
		GetTables(ctx context.Context, req v1.GetTablesReq) (res *v1.GetTablesRes, err error)
		// ImportTables 导入表信息到代码生成记录
		ImportTables(ctx context.Context, req v1.ImportTablesReq) (res *v1.ImportTablesRes, err error)
		GetTablesWithColumns(ctx context.Context, req v1.GetTablesWithColumnsReq) (res *v1.GetTablesWithColumnsRes, err error)
		// GetTableColumns 根据表名获取字段信息
		GetTableColumns(ctx context.Context, req v1.GetTableColumnsReq) (res *v1.GetTableColumnsRes, err error)
	}
)

var (
	localGenerate IGenerate
)

func Generate() IGenerate {
	if localGenerate == nil {
		panic("implement not found for interface IGenerate, forgot register?")
	}
	return localGenerate
}

func RegisterGenerate(i IGenerate) {
	localGenerate = i
}
