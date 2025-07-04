// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrdersDao is the data access object for the table orders.
type OrdersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OrdersColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OrdersColumns defines and stores column names for the table orders.
type OrdersColumns struct {
	Id          string // 订单ID
	OrderNo     string // 订单编号
	UserId      string // 用户ID
	ProductId   string // 商品ID
	ProductName string // 商品名称
	Quantity    string // 购买数量
	Price       string // 商品单价
	TotalAmount string // 订单总金额
	Status      string // 订单状态(0:待支付,1:已支付,2:已发货,3:已完成,4:已取消)
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// ordersColumns holds the columns for the table orders.
var ordersColumns = OrdersColumns{
	Id:          "id",
	OrderNo:     "order_no",
	UserId:      "user_id",
	ProductId:   "product_id",
	ProductName: "product_name",
	Quantity:    "quantity",
	Price:       "price",
	TotalAmount: "total_amount",
	Status:      "status",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewOrdersDao creates and returns a new DAO object for table data access.
func NewOrdersDao(handlers ...gdb.ModelHandler) *OrdersDao {
	return &OrdersDao{
		group:    "default",
		table:    "orders",
		columns:  ordersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrdersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrdersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrdersDao) Columns() OrdersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrdersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrdersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OrdersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
