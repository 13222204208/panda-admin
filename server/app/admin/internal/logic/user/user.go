package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/middleware"
	"server/app/admin/internal/model/do"
	"server/app/admin/internal/model/entity"
	"server/utility"
)

type sUser struct{}

func New() *sUser {
	return &sUser{}
}

// Create 创建用户
func (s *sUser) Create(ctx context.Context, in v1.CreateReq) (out *v1.CreateRes, err error) {
	out = &v1.CreateRes{}

	// 验证必填字段
	if in.Username == nil || *in.Username == "" {
		return nil, gerror.New("用户名不能为空")
	}

	// 检查用户名唯一性
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, *in.Username).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户名失败")
	}
	if count > 0 {
		return nil, gerror.Newf("用户名 %s 已存在", *in.Username)
	}

	// 检查邮箱唯一性（如果提供）
	if in.Email != nil && *in.Email != "" {
		count, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Email, *in.Email).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "查询邮箱失败")
		}
		if count > 0 {
			return nil, gerror.Newf("邮箱 %s 已存在", *in.Email)
		}
	}

	// 检查手机号唯一性（如果提供）
	if in.Phone != nil && *in.Phone != "" {
		count, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Phone, *in.Phone).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "查询手机号失败")
		}
		if count > 0 {
			return nil, gerror.Newf("手机号 %s 已存在", *in.Phone)
		}
	}
	//密码是必填项,且要大于等于6位
	if in.Password == nil || *in.Password == "" || len(*in.Password) < 6 {
		return nil, gerror.New("密码不能为空且长度至少为6位")
	}

	// 使用 bcrypt 加密密码
	hashedPassword, err := utility.EncryptPassword(*in.Password)
	if err != nil {
		return nil, gerror.Wrap(err, "密码加密失败")
	}

	// 构建插入数据
	data := g.Map{
		dao.User.Columns().Username: *in.Username,
		dao.User.Columns().Password: hashedPassword,
		dao.User.Columns().Status:   gconv.Int(gutil.GetOrDefaultAny(in.Status, 1)), // 默认启用
	}

	// 可选字段
	if in.Title != nil {
		data[dao.User.Columns().Title] = *in.Title
	}
	if in.DepartmentId != nil {
		data[dao.User.Columns().DepartmentId] = *in.DepartmentId
	}
	if in.Nickname != nil {
		data[dao.User.Columns().Nickname] = *in.Nickname
	}
	if in.Phone != nil {
		data[dao.User.Columns().Phone] = *in.Phone
	}
	if in.Email != nil {
		data[dao.User.Columns().Email] = *in.Email
	}
	if in.Sex != nil {
		data[dao.User.Columns().Sex] = *in.Sex
	}
	if in.Remark != nil {
		data[dao.User.Columns().Remark] = *in.Remark
	}

	// 插入用户数据
	id, err := dao.User.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建用户失败")
	}

	out.Id = uint64(id)
	return
}

// GetList 获取用户列表
func (s *sUser) GetList(ctx context.Context, in v1.GetListReq) (out *v1.GetListRes, err error) {
	out = &v1.GetListRes{}

	// 构建查询条件
	m := dao.User.Ctx(ctx)

	if in.DepartmentId != nil {
		m = m.Where(dao.User.Columns().DepartmentId, *in.DepartmentId)
	}
	if in.Username != "" {
		m = m.WhereLike(dao.User.Columns().Username, "%"+in.Username+"%")
	}
	if in.Phone != "" {
		m = m.WhereLike(dao.User.Columns().Phone, "%"+in.Phone+"%")
	}
	if in.Status != nil {
		m = m.Where(dao.User.Columns().Status, *in.Status)
	}

	// 获取总数
	totalCount, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户总数失败")
	}
	out.Total = totalCount

	// 分页查询
	offset := (in.CurrentPage - 1) * in.PageSize
	err = m.Limit(offset, in.PageSize).OrderDesc(dao.User.Columns().CreatedAt).Scan(&out.List)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户列表失败")
	}

	// 填充部门名称
	for i := range out.List {
		if out.List[i].DepartmentId > 0 {
			var dept v1.Dept
			err = dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, out.List[i].DepartmentId).Scan(&dept)
			if err == nil {
				out.List[i].Dept.Id = dept.Id
				out.List[i].Dept.Name = dept.Name
			}
		}
	}

	return
}

// Update 更新用户
func (s *sUser) Update(ctx context.Context, in v1.UpdateReq) (err error) {
	// 检查用户是否存在
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询用户失败")
	}
	if count == 0 {
		return gerror.Newf("用户ID %d 不存在", in.Id)
	}

	// 动态构建更新数据
	updateData := g.Map{}

	// 检查并添加需要更新的字段
	if in.Username != nil {
		if *in.Username == "" {
			return gerror.New("用户名不能为空")
		}
		// 检查用户名唯一性
		count, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Username, *in.Username).WhereNot(dao.User.Columns().Id, in.Id).Count()
		if err != nil {
			return gerror.Wrap(err, "查询用户名失败")
		}
		if count > 0 {
			return gerror.Newf("用户名 %s 已存在", *in.Username)
		}
		updateData[dao.User.Columns().Username] = *in.Username
	}

	if in.Email != nil {
		if *in.Email != "" {
			// 检查邮箱唯一性
			count, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Email, *in.Email).WhereNot(dao.User.Columns().Id, in.Id).Count()
			if err != nil {
				return gerror.Wrap(err, "查询邮箱失败")
			}
			if count > 0 {
				return gerror.Newf("邮箱 %s 已存在", *in.Email)
			}
		}
		updateData[dao.User.Columns().Email] = *in.Email
	}

	if in.Phone != nil {
		if *in.Phone != "" {
			// 检查手机号唯一性
			count, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Phone, *in.Phone).WhereNot(dao.User.Columns().Id, in.Id).Count()
			if err != nil {
				return gerror.Wrap(err, "查询手机号失败")
			}
			if count > 0 {
				return gerror.Newf("手机号 %s 已存在", *in.Phone)
			}
		}
		updateData[dao.User.Columns().Phone] = *in.Phone
	}

	// 其他字段更新
	if in.Title != nil {
		updateData[dao.User.Columns().Title] = *in.Title
	}
	if in.DepartmentId != nil {
		updateData[dao.User.Columns().DepartmentId] = *in.DepartmentId
	}
	if in.Nickname != nil {
		updateData[dao.User.Columns().Nickname] = *in.Nickname
	}
	if in.Sex != nil {
		updateData[dao.User.Columns().Sex] = *in.Sex
	}
	if in.Status != nil {
		updateData[dao.User.Columns().Status] = *in.Status
	}
	if in.Remark != nil {
		updateData[dao.User.Columns().Remark] = *in.Remark
	}

	// 更新用户数据
	if len(updateData) > 0 {
		_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Data(updateData).Update()
		if err != nil {
			return gerror.Wrap(err, "更新用户失败")
		}
	}

	return
}

// Delete 删除用户
func (s *sUser) Delete(ctx context.Context, in v1.DeleteReq) (err error) {
	// 检查用户是否存在
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询用户失败")
	}
	if count == 0 {
		return gerror.Newf("用户ID %d 不存在", in.Id)
	}

	// 删除用户角色关联
	_, err = dao.UserRole.Ctx(ctx).Where(dao.UserRole.Columns().UserId, in.Id).Delete()
	if err != nil {
		return gerror.Wrap(err, "删除用户角色关联失败")
	}

	// 删除用户数据
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Delete()
	return gerror.Wrap(err, "删除用户失败")
}

// GetDetail 获取用户详情
func (s *sUser) GetDetail(ctx context.Context, in v1.GetDetailReq) (out *v1.GetDetailRes, err error) {
	out = &v1.GetDetailRes{}
	//根据token获取用户id
	userID, ok := ctx.Value(middleware.CtxUserID).(uint64)
	if !ok {
		return nil, gerror.New("无法获取用户ID")
	}
	// 检查用户是否存在

	// 查询用户基本信息
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userID).Scan(&out.UserInfo)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户信息失败")
	}
	if out.UserInfo.Id == 0 {
		return nil, gerror.Newf("用户ID %d 不存在", userID)
	}

	return
}

// ResetPassword 重置密码
func (s *sUser) ResetPassword(ctx context.Context, in v1.ResetPasswordReq) (err error) {
	// 检查用户是否存在
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询用户失败")
	}
	if count == 0 {
		return gerror.Newf("用户ID %d 不存在", in.Id)
	}
	fmt.Println("参数:", in)
	//如果oldPassword不为空,则需要验证旧密码是否正确
	if in.OldPassword != "" {
		var user entity.User
		err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Scan(&user)
		if err != nil {
			return gerror.Wrap(err, "查询用户信息失败")
		}
		if user.Id == 0 {
			return gerror.Newf("用户ID %d 不存在", in.Id)
		}
		// 验证旧密码是否正确
		if err = utility.ComparePassword(user.Password, in.OldPassword); err != nil {
			return gerror.New("原密码不正确")
		}
		fmt.Println("验证密码:", err)
	}

	// 使用 bcrypt 加密密码
	hashedPassword, err := utility.EncryptPassword(in.Password)
	if err != nil {
		return gerror.Wrap(err, "密码加密失败")
	}

	// 更新密码
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Data(g.Map{
		dao.User.Columns().Password: hashedPassword,
	}).Update()
	return gerror.Wrap(err, "重置密码失败")
}

// updateUserRoles 更新用户角色关联
func (s *sUser) updateUserRoles(ctx context.Context, userId uint64, roleIds []uint64) error {
	// 删除现有角色关联
	_, err := dao.UserRole.Ctx(ctx).Where(dao.UserRole.Columns().UserId, userId).Delete()
	if err != nil {
		return gerror.Wrap(err, "删除现有角色关联失败")
	}

	// 添加新的角色关联
	if len(roleIds) > 0 {
		var userRoles []do.UserRole
		for _, roleId := range roleIds {
			userRoles = append(userRoles, do.UserRole{
				UserId: userId,
				RoleId: roleId,
			})
		}
		_, err = dao.UserRole.Ctx(ctx).Data(userRoles).Insert()
		if err != nil {
			return gerror.Wrap(err, "添加角色关联失败")
		}
	}

	return nil
}

// BatchDelete 批量删除用户
func (s *sUser) BatchDelete(ctx context.Context, in v1.BatchDeleteReq) (err error) {
	// 参数验证
	if len(in.Ids) == 0 {
		return gerror.New("请选择要删除的用户")
	}

	// 检查用户是否存在
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id+" IN (?)", in.Ids).Count()
	if err != nil {
		return gerror.Wrap(err, "查询用户失败")
	}
	if count != len(in.Ids) {
		return gerror.New("部分用户不存在，无法批量删除")
	}

	// 使用事务确保数据一致性
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除用户角色关联
		_, err = tx.Model(dao.UserRole.Table()).Ctx(ctx).Where(dao.UserRole.Columns().UserId+" IN (?)", in.Ids).Delete()
		if err != nil {
			return gerror.Wrap(err, "删除用户角色关联失败")
		}

		// 批量删除用户数据
		_, err = tx.Model(dao.User.Table()).Ctx(ctx).Where(dao.User.Columns().Id+" IN (?)", in.Ids).Delete()
		if err != nil {
			return gerror.Wrap(err, "批量删除用户失败")
		}

		return nil
	})
}

// GetRoleIds 获取用户对应的角色ID列表
func (s *sUser) GetRoleIds(ctx context.Context, in v1.GetRoleIdsReq) (out *v1.GetRoleIdsRes, err error) {
	out = &v1.GetRoleIdsRes{}

	// 查询用户角色关联表，获取角色ID列表
	roleIdArray, err := dao.UserRole.Ctx(ctx).
		Where(dao.UserRole.Columns().UserId, in.UserId).
		Fields(dao.UserRole.Columns().RoleId).
		Array()
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户角色ID列表失败")
	}

	// 转换为uint64切片
	var roleIds []uint64
	for _, v := range roleIdArray {
		if roleId := v.Uint64(); roleId > 0 {
			roleIds = append(roleIds, roleId)
		}
	}

	out.RoleIds = roleIds
	return
}

// AssignRoles 分配用户角色
func (s *sUser) AssignRoles(ctx context.Context, in v1.AssignRolesReq) (err error) {
	// 参数验证
	if in.UserId == 0 {
		return gerror.New("用户ID不能为空")
	}

	// 检查用户是否存在
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.UserId).Count()
	if err != nil {
		return gerror.Wrap(err, "查询用户失败")
	}
	if count == 0 {
		return gerror.New("用户不存在")
	}

	// 验证角色ID是否有效
	if len(in.RoleIds) > 0 {
		roleCount, err := dao.Role.Ctx(ctx).WhereIn(dao.Role.Columns().Id, in.RoleIds).Count()
		if err != nil {
			return gerror.Wrap(err, "查询角色失败")
		}
		if roleCount != len(in.RoleIds) {
			return gerror.New("存在无效的角色ID")
		}
	}

	// 更新用户角色关联
	err = s.updateUserRoles(ctx, in.UserId, in.RoleIds)
	if err != nil {
		return gerror.Wrap(err, "分配用户角色失败")
	}

	return nil
}

// UploadAvatar 上传用户头像
func (s *sUser) UploadAvatar(ctx context.Context, req v1.UploadAvatarReq) (*v1.UploadAvatarRes, error) {
	// 检查用户是否存在
	var user *entity.User
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, req.Id).Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户失败")
	}
	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	// 使用公共工具保存头像
	avatarUrl, err := utility.SaveBase64Avatar(ctx, req.Avatar, uint64(req.Id))
	if err != nil {
		return nil, gerror.Wrap(err, "保存头像失败")
	}

	// 更新用户头像
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, req.Id).Update(g.Map{
		dao.User.Columns().Avatar: avatarUrl,
	})
	if err != nil {
		return nil, gerror.Wrap(err, "更新用户头像失败")
	}

	// 返回结果
	return &v1.UploadAvatarRes{
		AvatarUrl: avatarUrl,
	}, nil
}
