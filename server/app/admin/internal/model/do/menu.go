// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta          `orm:"table:menu, do:true"`
	Id              interface{} // 主键ID
	MenuType        interface{} // 菜单类型（0菜单、1 iframe、2外链、3按钮）
	ParentId        interface{} // 父级菜单ID
	Title           interface{} // 菜单名称
	Name            interface{} // 路由名称（必须唯一）
	Path            interface{} // 路由路径
	Component       interface{} // 组件路径
	Rank            interface{} // 菜单排序（home 的 rank 应为 0）
	Redirect        interface{} // 重定向地址
	Icon            interface{} // 菜单图标
	ExtraIcon       interface{} // 右侧额外图标
	EnterTransition interface{} // 进场动画
	LeaveTransition interface{} // 离场动画
	ActivePath      interface{} // 激活菜单的 path
	Auths           interface{} // 权限标识（按钮级别权限）
	FrameSrc        interface{} // iframe 链接地址
	FrameLoading    interface{} // iframe 页面是否首次加载显示动画
	KeepAlive       interface{} // 是否缓存该页面
	HiddenTag       interface{} // 是否禁止添加到标签页
	FixedTag        interface{} // 是否固定在标签页中
	ShowLink        interface{} // 是否在菜单中显示该项
	ShowParent      interface{} // 是否显示父级菜单
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 修改时间
}
