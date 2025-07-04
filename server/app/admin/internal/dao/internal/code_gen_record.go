// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CodeGenRecordDao is the data access object for the table code_gen_record.
type CodeGenRecordDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  CodeGenRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// CodeGenRecordColumns defines and stores column names for the table code_gen_record.
type CodeGenRecordColumns struct {
	Id           string // 主键ID
	TableName    string // 数据表名称
	TableComment string // 表注释
	PackageName  string // 生成的Go包名
	ModuleName   string // 模块名（例如 system、user）
	Options      string // 配置选项
	Columns      string // 表字段
	Status       string // 生成状态（1成功，0失败）代码是否已生成
	CreatedAt    string // 生成时间
	UpdatedAt    string // 更新时间
}

// codeGenRecordColumns holds the columns for the table code_gen_record.
var codeGenRecordColumns = CodeGenRecordColumns{
	Id:           "id",
	TableName:    "table_name",
	TableComment: "table_comment",
	PackageName:  "package_name",
	ModuleName:   "module_name",
	Options:      "options",
	Columns:      "columns",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewCodeGenRecordDao creates and returns a new DAO object for table data access.
func NewCodeGenRecordDao(handlers ...gdb.ModelHandler) *CodeGenRecordDao {
	return &CodeGenRecordDao{
		group:    "default",
		table:    "code_gen_record",
		columns:  codeGenRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CodeGenRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CodeGenRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CodeGenRecordDao) Columns() CodeGenRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CodeGenRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CodeGenRecordDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CodeGenRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
