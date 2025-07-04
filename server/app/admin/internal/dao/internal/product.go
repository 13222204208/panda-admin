// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductDao is the data access object for the table product.
type ProductDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ProductColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ProductColumns defines and stores column names for the table product.
type ProductColumns struct {
	Id        string // 商品ID
	Name      string // 商品名称
	Price     string // 商品价格
	Stock     string // 商品库存
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// productColumns holds the columns for the table product.
var productColumns = ProductColumns{
	Id:        "id",
	Name:      "name",
	Price:     "price",
	Stock:     "stock",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewProductDao creates and returns a new DAO object for table data access.
func NewProductDao(handlers ...gdb.ModelHandler) *ProductDao {
	return &ProductDao{
		group:    "default",
		table:    "product",
		columns:  productColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProductDao) Columns() ProductColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProductDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
