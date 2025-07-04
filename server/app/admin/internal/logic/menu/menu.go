package menu

import (
	"context"
	v1 "server/app/admin/api/menu/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMenu struct{}

func New() *sMenu {
	return &sMenu{}
}

// GetList 获取菜单列表（不分页，获取全部菜单）
func (s *sMenu) GetList(ctx context.Context, req v1.GetListReq) (*v1.GetListRes, error) {
	var (
		menus []entity.Menu
		query = dao.Menu.Ctx(ctx)
	)

	// 构建查询条件
	if req.Title != "" {
		query = query.WhereLike(dao.Menu.Columns().Title, "%"+req.Title+"%")
	}
	if req.MenuType != nil {
		query = query.Where(dao.Menu.Columns().MenuType, *req.MenuType)
	}
	if req.ParentId != nil {
		query = query.Where(dao.Menu.Columns().ParentId, *req.ParentId)
	}

	// 按排序字段和ID排序
	err := query.OrderAsc(dao.Menu.Columns().Rank).OrderAsc(dao.Menu.Columns().Id).Scan(&menus)
	if err != nil {
		return nil, gerror.Wrap(err, "查询菜单列表失败")
	}

	// 转换为响应格式
	var list []v1.MenuInfo
	if err := gconv.Scan(menus, &list); err != nil {
		return nil, gerror.Wrap(err, "数据转换失败")
	}

	return &v1.GetListRes{
		List: list,
	}, nil
}

// Create 创建菜单
func (s *sMenu) Create(ctx context.Context, req v1.CreateReq) (*v1.CreateRes, error) {
	// 检查路由名称是否唯一
	if req.Name != nil && *req.Name != "" {
		count, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Name, *req.Name).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "检查路由名称失败")
		}
		if count > 0 {
			return nil, gerror.New("路由名称已存在")
		}
	}

	// 检查父级菜单是否存在
	if req.ParentId != nil && *req.ParentId > 0 {
		count, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, *req.ParentId).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "检查父级菜单失败")
		}
		if count == 0 {
			return nil, gerror.New("父级菜单不存在")
		}
	}

	// 插入菜单数据
	id, err := dao.Menu.Ctx(ctx).Data(req).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建菜单失败")
	}

	return &v1.CreateRes{
		Id: uint64(id),
	}, nil
}

// Update 更新菜单
func (s *sMenu) Update(ctx context.Context, req v1.UpdateReq) (*v1.UpdateRes, error) {
	// 检查菜单是否存在
	count, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "检查菜单失败")
	}
	if count == 0 {
		return nil, gerror.New("菜单不存在")
	}

	// 检查路由名称是否唯一（排除当前菜单）
	if req.Name != nil && *req.Name != "" {
		count, err = dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Name, *req.Name).WhereNot(dao.Menu.Columns().Id, req.Id).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "检查路由名称失败")
		}
		if count > 0 {
			return nil, gerror.New("路由名称已存在")
		}
	}

	// 检查父级菜单是否存在（不能设置自己为父级）
	if req.ParentId != nil && *req.ParentId > 0 {
		if *req.ParentId == req.Id {
			return nil, gerror.New("不能设置自己为父级菜单")
		}
		count, err = dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, *req.ParentId).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "检查父级菜单失败")
		}
		if count == 0 {
			return nil, gerror.New("父级菜单不存在")
		}
	}

	// 更新菜单数据
	_, err = dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, req.Id).Data(req.MenuCommon).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新菜单失败")
	}

	return &v1.UpdateRes{}, nil
}

// Delete 删除菜单
func (s *sMenu) Delete(ctx context.Context, req v1.DeleteReq) (*v1.DeleteRes, error) {
	// 检查菜单是否存在
	count, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "检查菜单失败")
	}
	if count == 0 {
		return nil, gerror.New("菜单不存在")
	}

	// 检查是否有子菜单
	childCount, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().ParentId, req.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "检查子菜单失败")
	}
	if childCount > 0 {
		return nil, gerror.New("存在子菜单，无法删除")
	}

	// 删除菜单
	_, err = dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除菜单失败")
	}

	return &v1.DeleteRes{}, nil
}
