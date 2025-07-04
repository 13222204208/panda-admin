// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DepartmentDao is the data access object for the table department.
type DepartmentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DepartmentColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DepartmentColumns defines and stores column names for the table department.
type DepartmentColumns struct {
	Id        string // 主键ID
	ParentId  string // 父级部门ID
	Name      string // 部门名称
	Principal string // 负责人名称
	Phone     string // 联系电话
	Email     string // 邮箱地址
	Sort      string // 排序号
	Status    string // 状态（1启用，0禁用）
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// departmentColumns holds the columns for the table department.
var departmentColumns = DepartmentColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	Principal: "principal",
	Phone:     "phone",
	Email:     "email",
	Sort:      "sort",
	Status:    "status",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewDepartmentDao creates and returns a new DAO object for table data access.
func NewDepartmentDao(handlers ...gdb.ModelHandler) *DepartmentDao {
	return &DepartmentDao{
		group:    "default",
		table:    "department",
		columns:  departmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DepartmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DepartmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DepartmentDao) Columns() DepartmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DepartmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DepartmentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DepartmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
