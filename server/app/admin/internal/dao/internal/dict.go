// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DictDao is the data access object for the table dict.
type DictDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DictColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DictColumns defines and stores column names for the table dict.
type DictColumns struct {
	Id        string // 主键ID
	Title     string // 字典名称
	DictType  string // 字典类型（如 sex、status、job_type）
	DictLabel string // 字典标签（如 男、启用）
	DictValue string // 字典值（如 1、enabled）
	Sort      string // 排序值（升序）
	Status    string // 状态（1启用，0禁用）
	Remark    string // 备注说明
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// dictColumns holds the columns for the table dict.
var dictColumns = DictColumns{
	Id:        "id",
	Title:     "title",
	DictType:  "dict_type",
	DictLabel: "dict_label",
	DictValue: "dict_value",
	Sort:      "sort",
	Status:    "status",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewDictDao creates and returns a new DAO object for table data access.
func NewDictDao(handlers ...gdb.ModelHandler) *DictDao {
	return &DictDao{
		group:    "default",
		table:    "dict",
		columns:  dictColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DictDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DictDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DictDao) Columns() DictColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DictDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DictDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DictDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
