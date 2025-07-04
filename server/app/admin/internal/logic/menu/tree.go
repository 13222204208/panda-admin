package menu

import (
	"context"
	v1 "server/app/admin/api/menu/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

// GetTree 获取菜单树状结构
func (s *sMenu) GetTree(ctx context.Context, req v1.GetTreeReq) (*v1.GetTreeRes, error) {
	var menus []entity.Menu

	// 查询所有显示的菜单（showLink = 1表示显示）
	err := dao.Menu.Ctx(ctx).
		Where(dao.Menu.Columns().ShowLink, 1). // 过滤隐藏的菜单
		OrderAsc(dao.Menu.Columns().Rank).
		OrderAsc(dao.Menu.Columns().Id).
		Scan(&menus)
	if err != nil {
		return nil, gerror.Wrap(err, "查询菜单列表失败")
	}

	// 构建菜单树，从parentId=0开始
	children := buildMenuTree(menus, 0)

	// 创建最外层主菜单，id为0
	mainMenu := v1.MenuTreeNode{
		Id:       0,
		Title:    "主菜单",
		Children: children,
	}

	return &v1.GetTreeRes{
		Tree: []v1.MenuTreeNode{mainMenu},
	}, nil
}

// buildMenuTree 构建菜单树
func buildMenuTree(menus []entity.Menu, parentId uint64) []v1.MenuTreeNode {
	var tree []v1.MenuTreeNode

	for _, menu := range menus {
		if menu.ParentId == parentId {
			node := v1.MenuTreeNode{
				Id:    menu.Id,
				Title: menu.Title,
			}

			// 递归查找子菜单
			children := buildMenuTree(menus, menu.Id)
			if len(children) > 0 {
				node.Children = children
			}

			tree = append(tree, node)
		}
	}

	return tree
}
