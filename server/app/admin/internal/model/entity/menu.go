// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id              uint64      `json:"id"              orm:"id"               description:"主键ID"`                       // 主键ID
	MenuType        int         `json:"menuType"        orm:"menu_type"        description:"菜单类型（0菜单、1 iframe、2外链、3按钮）"` // 菜单类型（0菜单、1 iframe、2外链、3按钮）
	ParentId        uint64      `json:"parentId"        orm:"parent_id"        description:"父级菜单ID"`                     // 父级菜单ID
	Title           string      `json:"title"           orm:"title"            description:"菜单名称"`                       // 菜单名称
	Name            string      `json:"name"            orm:"name"             description:"路由名称（必须唯一）"`                 // 路由名称（必须唯一）
	Path            string      `json:"path"            orm:"path"             description:"路由路径"`                       // 路由路径
	Component       string      `json:"component"       orm:"component"        description:"组件路径"`                       // 组件路径
	Rank            int         `json:"rank"            orm:"rank"             description:"菜单排序（home 的 rank 应为 0）"`     // 菜单排序（home 的 rank 应为 0）
	Redirect        string      `json:"redirect"        orm:"redirect"         description:"重定向地址"`                      // 重定向地址
	Icon            string      `json:"icon"            orm:"icon"             description:"菜单图标"`                       // 菜单图标
	ExtraIcon       string      `json:"extraIcon"       orm:"extra_icon"       description:"右侧额外图标"`                     // 右侧额外图标
	EnterTransition string      `json:"enterTransition" orm:"enter_transition" description:"进场动画"`                       // 进场动画
	LeaveTransition string      `json:"leaveTransition" orm:"leave_transition" description:"离场动画"`                       // 离场动画
	ActivePath      string      `json:"activePath"      orm:"active_path"      description:"激活菜单的 path"`                 // 激活菜单的 path
	Auths           string      `json:"auths"           orm:"auths"            description:"权限标识（按钮级别权限）"`               // 权限标识（按钮级别权限）
	FrameSrc        string      `json:"frameSrc"        orm:"frame_src"        description:"iframe 链接地址"`                // iframe 链接地址
	FrameLoading    int         `json:"frameLoading"    orm:"frame_loading"    description:"iframe 页面是否首次加载显示动画"`        // iframe 页面是否首次加载显示动画
	KeepAlive       int         `json:"keepAlive"       orm:"keep_alive"       description:"是否缓存该页面"`                    // 是否缓存该页面
	HiddenTag       int         `json:"hiddenTag"       orm:"hidden_tag"       description:"是否禁止添加到标签页"`                 // 是否禁止添加到标签页
	FixedTag        int         `json:"fixedTag"        orm:"fixed_tag"        description:"是否固定在标签页中"`                  // 是否固定在标签页中
	ShowLink        int         `json:"showLink"        orm:"show_link"        description:"是否在菜单中显示该项"`                 // 是否在菜单中显示该项
	ShowParent      int         `json:"showParent"      orm:"show_parent"      description:"是否显示父级菜单"`                   // 是否显示父级菜单
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`                       // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"修改时间"`                       // 修改时间
}
