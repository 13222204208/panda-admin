package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetTreeReq 获取菜单树请求参数
type GetTreeReq struct {
	g.Meta `path:"/menu/tree" method:"get" tags:"菜单管理" summary:"获取菜单树"`
}

// GetTreeRes 获取菜单树返回参数
type GetTreeRes struct {
	Tree []MenuTreeNode `json:"tree" dc:"菜单树"`
}

// MenuTreeNode 菜单树节点
type MenuTreeNode struct {
	Id       uint64           `json:"id" dc:"菜单ID"`
	Title    string           `json:"title" dc:"菜单标题"`
	Children []MenuTreeNode   `json:"children,omitempty" dc:"子菜单"`
}