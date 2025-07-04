package v1

import "github.com/gogf/gf/v2/frame/g"

// GetUserRoutesReq 获取用户路由权限请求参数
type GetUserRoutesReq struct {
	g.Meta `path:"/user/routes" method:"get" tags:"用户管理" summary:"获取用户路由权限"`
}

// GetUserRoutesRes 获取用户路由权限返回参数
type GetUserRoutesRes struct {
	Routes []RouteInfo `json:"routes" dc:"路由列表"`
}

// RouteInfo 路由信息
type RouteInfo struct {
	Path      string      `json:"path" dc:"路由路径"`
	Name      string      `json:"name,omitempty" dc:"路由名称"`
	Meta      RouteMeta   `json:"meta" dc:"路由元信息"`
	Component string      `json:"component" dc:"组件路径"`
	Children  []RouteInfo `json:"children,omitempty" dc:"子路由"`
}

// RouteMeta 路由元信息
type RouteMeta struct {
	Icon            string   `json:"icon,omitempty" dc:"图标"`
	Title           string   `json:"title" dc:"标题"`
	Rank            int      `json:"rank,omitempty" dc:"排序"`
	Roles           []string `json:"roles,omitempty" dc:"角色权限"`
	KeepAlive       bool     `json:"keepAlive,omitempty" dc:"是否缓存"`
	FrameSrc        string   `json:"frameSrc,omitempty" dc:"iframe链接"`
	HiddenTag       bool     `json:"hiddenTag,omitempty" dc:"是否隐藏标签"`
	FixedTag        bool     `json:"fixedTag,omitempty" dc:"是否固定标签"`
	ShowLink        bool     `json:"showLink" dc:"是否显示链接"`
	ShowParent      bool     `json:"showParent,omitempty" dc:"是否显示父级"`
	EnterTransition string   `json:"enterTransition,omitempty" dc:"进场动画"`
	LeaveTransition string   `json:"leaveTransition,omitempty" dc:"离场动画"`
	ActivePath      string   `json:"activePath,omitempty" dc:"激活路径"`
}
