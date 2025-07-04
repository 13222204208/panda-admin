// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuDao is the data access object for the table menu.
type MenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MenuColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MenuColumns defines and stores column names for the table menu.
type MenuColumns struct {
	Id              string // 主键ID
	MenuType        string // 菜单类型（0菜单、1 iframe、2外链、3按钮）
	ParentId        string // 父级菜单ID
	Title           string // 菜单名称
	Name            string // 路由名称（必须唯一）
	Path            string // 路由路径
	Component       string // 组件路径
	Rank            string // 菜单排序（home 的 rank 应为 0）
	Redirect        string // 重定向地址
	Icon            string // 菜单图标
	ExtraIcon       string // 右侧额外图标
	EnterTransition string // 进场动画
	LeaveTransition string // 离场动画
	ActivePath      string // 激活菜单的 path
	Auths           string // 权限标识（按钮级别权限）
	FrameSrc        string // iframe 链接地址
	FrameLoading    string // iframe 页面是否首次加载显示动画
	KeepAlive       string // 是否缓存该页面
	HiddenTag       string // 是否禁止添加到标签页
	FixedTag        string // 是否固定在标签页中
	ShowLink        string // 是否在菜单中显示该项
	ShowParent      string // 是否显示父级菜单
	CreatedAt       string // 创建时间
	UpdatedAt       string // 修改时间
}

// menuColumns holds the columns for the table menu.
var menuColumns = MenuColumns{
	Id:              "id",
	MenuType:        "menu_type",
	ParentId:        "parent_id",
	Title:           "title",
	Name:            "name",
	Path:            "path",
	Component:       "component",
	Rank:            "rank",
	Redirect:        "redirect",
	Icon:            "icon",
	ExtraIcon:       "extra_icon",
	EnterTransition: "enter_transition",
	LeaveTransition: "leave_transition",
	ActivePath:      "active_path",
	Auths:           "auths",
	FrameSrc:        "frame_src",
	FrameLoading:    "frame_loading",
	KeepAlive:       "keep_alive",
	HiddenTag:       "hidden_tag",
	FixedTag:        "fixed_tag",
	ShowLink:        "show_link",
	ShowParent:      "show_parent",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao(handlers ...gdb.ModelHandler) *MenuDao {
	return &MenuDao{
		group:    "default",
		table:    "menu",
		columns:  menuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
