package role

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "server/app/admin/api/role/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/model/entity"
)

type sRole struct{}

func New() *sRole {
	return &sRole{}
}

// Create 创建角色
func (s *sRole) Create(ctx context.Context, in v1.CreateReq) (out *v1.CreateRes, err error) {
	out = &v1.CreateRes{}

	// 检查编码唯一性并插入数据
	if count, _ := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Code, *in.Code).Count(); err != nil {
		return nil, gerror.Wrap(err, "查询角色编码失败")
	} else if count > 0 {
		return nil, gerror.Newf("角色编码 %s 已存在", *in.Code)
	}

	// 插入数据
	data := g.Map{
		dao.Role.Columns().Name:   *in.Name,
		dao.Role.Columns().Code:   *in.Code,
		dao.Role.Columns().Status: gconv.Int(in.Status),
	}
	if in.Remark != nil {
		data[dao.Role.Columns().Remark] = *in.Remark
	}

	id, err := dao.Role.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建角色失败")
	}

	out.Id = uint64(id)
	return
}

// GetList 获取角色列表
func (s *sRole) GetList(ctx context.Context, in v1.GetListReq) (out *v1.GetListRes, err error) {
	out = &v1.GetListRes{}

	// 构建查询条件
	m := dao.Role.Ctx(ctx)
	if in.Name != "" {
		m = m.WhereLike(dao.Role.Columns().Name, "%"+in.Name+"%")
	}
	if in.Code != "" {
		m = m.Where(dao.Role.Columns().Code, in.Code)
	}
	if in.Status != nil {
		m = m.Where(dao.Role.Columns().Status, in.Status)
	}

	// 查询总数
	out.Total, err = m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色总数失败")
	}

	// 查询列表
	err = m.Page(in.CurrentPage, in.PageSize).Scan(&out.List)
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色列表失败")
	}

	out.CurrentPage = in.CurrentPage
	return
}

// Update 更新角色
func (s *sRole) Update(ctx context.Context, in v1.UpdateReq) (err error) {
	// 检查角色是否存在
	count, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询角色失败")
	}
	if count == 0 {
		return gerror.Newf("角色ID %d 不存在", in.Id)
	}

	// 动态构建更新数据
	updateData := g.Map{}

	// 检查并添加需要更新的字段
	if in.Name != nil {
		if *in.Name == "" {
			return gerror.New("角色名称不能为空")
		}
		updateData[dao.Role.Columns().Name] = *in.Name
	}

	if in.Code != nil {
		if *in.Code == "" {
			return gerror.New("角色编码不能为空")
		}
		// 检查更新后的角色编码是否已存在（排除当前角色）
		count, err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Code, *in.Code).WhereNot(dao.Role.Columns().Id, in.Id).Count()
		if err != nil {
			return gerror.Wrap(err, "查询角色编码失败")
		}
		if count > 0 {
			return gerror.Newf("角色编码 %s 已存在", *in.Code)
		}
		updateData[dao.Role.Columns().Code] = *in.Code
	}

	if in.Status != nil {
		updateData[dao.Role.Columns().Status] = *in.Status
	}

	if in.Remark != nil {
		// 备注允许为空字符串
		updateData[dao.Role.Columns().Remark] = *in.Remark
	}

	// 检查是否有字段需要更新
	if len(updateData) == 0 {
		return gerror.New("没有需要更新的字段")
	}

	// 更新数据
	_, err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.Id).Data(updateData).Update()
	return gerror.Wrap(err, "更新角色失败")
}

// Delete 删除角色
func (s *sRole) Delete(ctx context.Context, in v1.DeleteReq) (err error) {
	// 检查角色是否存在
	count, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询角色失败")
	}
	if count == 0 {
		return gerror.Newf("角色ID %d 不存在", in.Id)
	}

	// 删除数据
	_, err = dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.Id).Delete()
	return gerror.Wrap(err, "删除角色失败")
}

// GetAll 获取所有角色列表（不分页）
func (s *sRole) GetAll(ctx context.Context, in v1.GetAllReq) (out *v1.GetAllRes, err error) {
	out = &v1.GetAllRes{}

	// 构建查询条件
	m := dao.Role.Ctx(ctx)
	// 只查询启用的角色
	m = m.Where(dao.Role.Columns().Status, 1)

	// 查询所有角色，按创建时间排序
	err = m.OrderAsc(dao.Role.Columns().CreatedAt).Scan(&out.List)
	if err != nil {
		return nil, gerror.Wrap(err, "查询所有角色列表失败")
	}

	return
}

// AssignMenus 分配角色菜单权限
func (s *sRole) AssignMenus(ctx context.Context, in v1.AssignMenusReq) (err error) {
	// 检查角色是否存在
	count, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询角色失败")
	}
	if count == 0 {
		return gerror.Newf("角色ID %d 不存在", in.Id)
	}

	// 验证菜单ID是否有效
	if len(in.MenuIds) > 0 {
		menuCount, err := dao.Menu.Ctx(ctx).WhereIn(dao.Menu.Columns().Id, in.MenuIds).Count()
		if err != nil {
			return gerror.Wrap(err, "查询菜单失败")
		}
		if menuCount != len(in.MenuIds) {
			return gerror.New("存在无效的菜单ID")
		}
	}

	// 使用事务更新角色菜单关联
	err = dao.RoleMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除原有的角色菜单关联
		_, err := dao.RoleMenu.Ctx(ctx).Where(dao.RoleMenu.Columns().RoleId, in.Id).Delete()
		if err != nil {
			return gerror.Wrap(err, "删除原有角色菜单关联失败")
		}

		// 插入新的角色菜单关联
		if len(in.MenuIds) > 0 {
			var insertData []g.Map
			for _, menuId := range in.MenuIds {
				insertData = append(insertData, g.Map{
					dao.RoleMenu.Columns().RoleId: in.Id,
					dao.RoleMenu.Columns().MenuId: menuId,
				})
			}
			_, err = dao.RoleMenu.Ctx(ctx).Data(insertData).Insert()
			if err != nil {
				return gerror.Wrap(err, "插入角色菜单关联失败")
			}
		}

		return nil
	})

	return gerror.Wrap(err, "分配角色菜单权限失败")
}

// GetRoleMenuIds 获取角色菜单ID列表
func (s *sRole) GetRoleMenuIds(ctx context.Context, in v1.GetRoleMenuIdsReq) (out *v1.GetRoleMenuIdsRes, err error) {
	out = &v1.GetRoleMenuIdsRes{}

	// 检查角色是否存在
	count, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色失败")
	}
	if count == 0 {
		return nil, gerror.Newf("角色ID %d 不存在", in.Id)
	}

	// 查询角色关联的菜单ID列表
	var roleMenus []entity.RoleMenu
	err = dao.RoleMenu.Ctx(ctx).Where(dao.RoleMenu.Columns().RoleId, in.Id).Scan(&roleMenus)
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色菜单关联失败")
	}

	// 提取菜单ID列表
	for _, roleMenu := range roleMenus {
		out.MenuIds = append(out.MenuIds, roleMenu.MenuId)
	}

	return
}
