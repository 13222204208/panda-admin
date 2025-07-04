package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuCommon 菜单公共字段
type MenuCommon struct {
	MenuType        *int    `json:"menuType,omitempty" v:"in:0,1,2,3#菜单类型只能是0,1,2,3" dc:"菜单类型（0菜单、1 iframe、2外链、3按钮）"`
	ParentId        *uint64 `json:"parentId,omitempty" dc:"父级菜单ID"`
	Title           *string `json:"title,omitempty" v:"required#请输入菜单名称" dc:"菜单名称"`
	Name            *string `json:"name,omitempty" v:"required#请输入路由名称" dc:"路由名称（必须唯一）"`
	Path            *string `json:"path,omitempty" dc:"路由路径"`
	Component       *string `json:"component,omitempty" dc:"组件路径"`
	Rank            *int    `json:"rank,omitempty" dc:"菜单排序（home 的 rank 应为 0）"`
	Redirect        *string `json:"redirect,omitempty" dc:"重定向地址"`
	Icon            *string `json:"icon,omitempty" dc:"菜单图标"`
	ExtraIcon       *string `json:"extraIcon,omitempty" dc:"右侧额外图标"`
	EnterTransition *string `json:"enterTransition,omitempty" dc:"进场动画"`
	LeaveTransition *string `json:"leaveTransition,omitempty" dc:"离场动画"`
	ActivePath      *string `json:"activePath,omitempty" dc:"激活菜单的 path"`
	Auths           *string `json:"auths,omitempty" dc:"权限标识（按钮级别权限）"`
	FrameSrc        *string `json:"frameSrc,omitempty" dc:"iframe 链接地址"`
	FrameLoading    *int    `json:"frameLoading,omitempty" v:"in:0,1#frameLoading只能是0或1" dc:"iframe 页面是否首次加载显示动画"`
	KeepAlive       *int    `json:"keepAlive,omitempty" v:"in:0,1#keepAlive只能是0或1" dc:"是否缓存该页面"`
	HiddenTag       *int    `json:"hiddenTag,omitempty" v:"in:0,1#hiddenTag只能是0或1" dc:"是否禁止添加到标签页"`
	FixedTag        *int    `json:"fixedTag,omitempty" v:"in:0,1#fixedTag只能是0或1" dc:"是否固定在标签页中"`
	ShowLink        *int    `json:"showLink,omitempty" v:"in:0,1#showLink只能是0或1" dc:"是否在菜单中显示该项"`
	ShowParent      *int    `json:"showParent,omitempty" v:"in:0,1#showParent只能是0或1" dc:"是否显示父级菜单"`
}

// GetListReq 获取菜单列表请求参数（不分页，获取全部菜单）
type GetListReq struct {
	g.Meta   `path:"/menu" method:"get" tags:"菜单管理" summary:"获取菜单列表"`
	Title    string  `json:"title" dc:"菜单名称"`
	MenuType *int    `json:"menuType" v:"in:0,1,2,3#菜单类型只能是0,1,2,3" dc:"菜单类型（0菜单、1 iframe、2外链、3按钮）"`
	ParentId *uint64 `json:"parentId" dc:"父级菜单ID"`
}

// GetListRes 获取菜单列表返回参数
type GetListRes struct {
	List []MenuInfo `json:"list" dc:"菜单列表"`
}

// CreateReq 创建菜单请求参数
type CreateReq struct {
	g.Meta `path:"/menu" method:"post" tags:"菜单管理" summary:"创建菜单"`
	MenuCommon
}

// CreateRes 创建菜单返回参数
type CreateRes struct {
	Id uint64 `json:"id" dc:"菜单ID"`
}

// UpdateReq 更新菜单请求参数
type UpdateReq struct {
	g.Meta `path:"/menu/{id}" method:"put" tags:"菜单管理" summary:"更新菜单"`
	Id     uint64 `json:"id" v:"required#请输入菜单ID" dc:"菜单ID"`
	MenuCommon
}

// UpdateRes 更新菜单返回参数
type UpdateRes struct{}

// DeleteReq 删除菜单请求参数
type DeleteReq struct {
	g.Meta `path:"/menu/{id}" method:"delete" tags:"菜单管理" summary:"删除菜单"`
	Id     uint64 `json:"id" v:"required#请输入菜单ID" dc:"菜单ID"`
}

// DeleteRes 删除菜单返回参数
type DeleteRes struct{}

// MenuInfo 菜单信息
type MenuInfo struct {
	Id              uint64      `json:"id" dc:"主键ID"`
	MenuType        int         `json:"menuType" dc:"菜单类型（0菜单、1 iframe、2外链、3按钮）"`
	ParentId        uint64      `json:"parentId" dc:"父级菜单ID"`
	Title           string      `json:"title" dc:"菜单名称"`
	Name            string      `json:"name" dc:"路由名称（必须唯一）"`
	Path            string      `json:"path" dc:"路由路径"`
	Component       string      `json:"component" dc:"组件路径"`
	Rank            int         `json:"rank" dc:"菜单排序（home 的 rank 应为 0）"`
	Redirect        string      `json:"redirect" dc:"重定向地址"`
	Icon            string      `json:"icon" dc:"菜单图标"`
	ExtraIcon       string      `json:"extraIcon" dc:"右侧额外图标"`
	EnterTransition string      `json:"enterTransition" dc:"进场动画"`
	LeaveTransition string      `json:"leaveTransition" dc:"离场动画"`
	ActivePath      string      `json:"activePath" dc:"激活菜单的 path"`
	Auths           string      `json:"auths" dc:"权限标识（按钮级别权限）"`
	FrameSrc        string      `json:"frameSrc" dc:"iframe 链接地址"`
	FrameLoading    int         `json:"frameLoading" dc:"iframe 页面是否首次加载显示动画"`
	KeepAlive       int         `json:"keepAlive" dc:"是否缓存该页面"`
	HiddenTag       int         `json:"hiddenTag" dc:"是否禁止添加到标签页"`
	FixedTag        int         `json:"fixedTag" dc:"是否固定在标签页中"`
	ShowLink        int         `json:"showLink" dc:"是否在菜单中显示该项"`
	ShowParent      int         `json:"showParent" dc:"是否显示父级菜单"`
	CreatedAt       *gtime.Time `json:"createdAt" dc:"创建时间"`
}
