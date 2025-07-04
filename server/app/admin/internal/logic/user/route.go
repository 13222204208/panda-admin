package user

import (
	"context"
	"fmt"
	v1 "server/app/admin/api/user/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/middleware"
	"server/app/admin/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

const (
	// DeveloperUsername 开发者用户名
	DeveloperUsername = "developer"
	// MenuTypeMenu 菜单类型
	MenuTypeMenu = 0
)

// GetUserRoutes 获取用户路由权限
func (s *sUser) GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (*v1.GetUserRoutesRes, error) {
	// 从上下文中获取用户ID
	userID, ok := ctx.Value(middleware.CtxUserID).(uint64)
	if !ok {
		return nil, gerror.New("无法获取用户ID")
	}

	// 根据用户ID查询用户信息
	var user entity.User
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userID).Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户信息失败")
	}
	if user.Id == 0 {
		return nil, gerror.New("用户不存在")
	}

	// 如果是 developer 用户，返回所有路由
	if user.Username == DeveloperUsername {
		routes, err := s.getAllRoutes(ctx)
		if err != nil {
			return nil, err
		}
		return &v1.GetUserRoutesRes{
			Routes: routes,
		}, nil
	}

	// 获取用户的菜单权限并转换为路由格式
	userMenus, err := s.getUserMenus(ctx, userID)
	if err != nil {
		return nil, gerror.Wrap(err, "获取用户菜单权限失败")
	}

	// 获取用户角色信息
	userRoles, err := s.getUserRoles(ctx, userID)
	if err != nil {
		return nil, gerror.Wrap(err, "获取用户角色失败")
	}

	routes := s.convertMenusToRoutes(userMenus, userRoles)
	fmt.Println("参数:", routes)
	return &v1.GetUserRoutesRes{
		Routes: routes,
	}, nil
}

// getAllRoutes 函数中的调用（developer用户）
func (s *sUser) getAllRoutes(ctx context.Context) ([]v1.RouteInfo, error) {
	var menus []entity.Menu

	// 查询所有菜单，按排序字段排序
	err := dao.Menu.Ctx(ctx).
		Where(dao.Menu.Columns().MenuType, 0). // 只查询菜单类型（0菜单）
		OrderAsc(dao.Menu.Columns().Rank).
		OrderAsc(dao.Menu.Columns().Id).
		Scan(&menus)
	if err != nil {
		return nil, gerror.Wrap(err, "查询菜单列表失败")
	}

	// 转换菜单数据为路由格式
	// 对于developer用户，可以传入空的角色列表或所有角色
	routes := s.convertMenusToRoutes(menus, []entity.Role{})

	return routes, nil
}

// getUserMenus 获取用户菜单权限
func (s *sUser) getUserMenus(ctx context.Context, userID uint64) ([]entity.Menu, error) {
	// 获取用户角色
	userRoles, err := s.getUserRoles(ctx, userID)
	if err != nil {
		return nil, gerror.Wrap(err, "获取用户角色失败")
	}

	// 获取角色对应的菜单权限
	menus, err := s.getRolePermissions(ctx, userRoles)
	if err != nil {
		return nil, gerror.Wrap(err, "获取角色菜单权限失败")
	}

	return menus, nil
}

// internalRouteInfo 内部路由信息（包含构建树结构需要的字段）
type internalRouteInfo struct {
	v1.RouteInfo
	Id       uint64
	ParentId uint64
}

// convertMenusToRoutes 将菜单转换为路由格式
func (s *sUser) convertMenusToRoutes(menus []entity.Menu, userRoles []entity.Role) []v1.RouteInfo {
	if len(menus) == 0 {
		return []v1.RouteInfo{}
	}

	// 提取角色编码列表，如果包含admin角色则返回admin
	roles := make([]string, 0, len(userRoles))
	hasAdmin := true
	for _, role := range userRoles {
		if role.Code == "admin" {
			hasAdmin = true
		}
		roles = append(roles, role.Code)
	}

	// 如果需要固定返回admin
	if hasAdmin {
		roles = []string{"admin", "common", "developer"}
	}

	// 先将菜单转换为内部路由结构
	internalRoutes := make([]internalRouteInfo, 0, len(menus))
	for _, menu := range menus {
		if menu.MenuType == MenuTypeMenu {
			internalRoute := internalRouteInfo{
				RouteInfo: v1.RouteInfo{
					Path:      menu.Path,
					Name:      menu.Name,
					Component: menu.Component,
					Meta: v1.RouteMeta{
						Icon:      menu.Icon,
						Title:     menu.Title,
						Rank:      menu.Rank,
						KeepAlive: menu.KeepAlive == 1,
						FrameSrc:  menu.FrameSrc,
						HiddenTag: menu.HiddenTag == 1,
						FixedTag:  menu.FixedTag == 1,
						ShowLink:  menu.ShowLink == 1,
						// ShowParent:      menu.ShowParent == 1,
						EnterTransition: menu.EnterTransition,
						LeaveTransition: menu.LeaveTransition,
						ActivePath:      menu.ActivePath,
						Roles:           roles, // 填充用户角色列表
					},
				},
				Id:       menu.Id,
				ParentId: menu.ParentId,
			}
			internalRoutes = append(internalRoutes, internalRoute)
		}
	}

	// 构建路由树
	routeTree := s.buildInternalRouteTree(internalRoutes, 0)
	return routeTree
}

// buildInternalRouteTree 构建内部路由树
func (s *sUser) buildInternalRouteTree(routes []internalRouteInfo, parentID uint64) []v1.RouteInfo {
	tree := make([]v1.RouteInfo, 0)

	for i := range routes {
		if routes[i].ParentId == parentID {
			// 递归构建子路由
			childRoutes := s.buildInternalRouteTree(routes, routes[i].Id)
			if len(childRoutes) > 0 {
				routes[i].RouteInfo.Children = childRoutes
			}
			tree = append(tree, routes[i].RouteInfo)
		}
	}

	return tree
}
