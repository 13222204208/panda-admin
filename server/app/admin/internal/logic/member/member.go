package member

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"server/app/admin/api/common/page"
	v1 "server/app/admin/api/member/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/model/entity"
)

type sMember struct{}

func New() *sMember {
	return &sMember{}
}


// GetMemberList 获取会员信息表列表
func (s *sMember) GetMemberList(ctx context.Context, in v1.GetMemberListReq) (out *v1.GetMemberListRes, err error) {
	out = &v1.GetMemberListRes{}

	m := dao.Member.Ctx(ctx)

	// 构建查询条件
	if in.Username != "" {
		m = m.WhereLike(dao.Member.Columns().Username, "%"+in.Username+"%")
	}
	if in.Email != "" {
		m = m.WhereLike(dao.Member.Columns().Email, "%"+in.Email+"%")
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询会员信息表总数失败")
	}

	// 分页查询
	// 初始化为空切片，确保返回空数组而不是null
	list := make([]*entity.Member, 0)
	err = m.Page(in.CurrentPage, in.PageSize).
		OrderDesc(dao.Member.Columns().CreatedAt).
		Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询会员信息表列表失败")
	}

	out.ResPage = page.ResPage{
		Total:       int(total),
		CurrentPage: in.CurrentPage,
	}
	out.List = list
	return
}



// CreateMember 创建会员信息表
func (s *sMember) CreateMember(ctx context.Context, in v1.CreateMemberReq) (out *v1.CreateMemberRes, err error) {
	out = &v1.CreateMemberRes{}
	// 验证必填字段
	if in.Username == nil || *in.Username == "" {
		return nil, gerror.New("用户名不能为空")
	}
	// 验证必填字段
	if in.Mobile == nil || *in.Mobile == "" {
		return nil, gerror.New("手机号不能为空")
	}

	// 构建插入数据
	data := g.Map{}
	if in.Username != nil {
		data[dao.Member.Columns().Username] = *in.Username
	}
	if in.Email != nil {
		data[dao.Member.Columns().Email] = *in.Email
	}
	if in.Mobile != nil {
		data[dao.Member.Columns().Mobile] = *in.Mobile
	}

	// 插入数据
	_, err = dao.Member.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建会员信息表失败")
	}
	return
}



// UpdateMember 更新会员信息表
func (s *sMember) UpdateMember(ctx context.Context, in v1.UpdateMemberReq) (out *v1.UpdateMemberRes, err error) {
	out = &v1.UpdateMemberRes{}

	// 检查会员信息表是否存在
	count, err := dao.Member.Ctx(ctx).Where(dao.Member.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询会员信息表失败")
	}
	if count == 0 {
		return nil, gerror.New("会员信息表不存在")
	}

	// 动态构建更新数据
	updateData := g.Map{}
	if in.Username != nil {
		updateData[dao.Member.Columns().Username] = *in.Username
	}
	if in.Email != nil {
		updateData[dao.Member.Columns().Email] = *in.Email
	}
	if in.Mobile != nil {
		updateData[dao.Member.Columns().Mobile] = *in.Mobile
	}

	// 更新数据
	_, err = dao.Member.Ctx(ctx).
		Where(dao.Member.Columns().Id, in.Id).
		Data(updateData).
		Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新会员信息表失败")
	}

	return
}



// DeleteMember 删除会员信息表
func (s *sMember) DeleteMember(ctx context.Context, in v1.DeleteMemberReq) (out *v1.DeleteMemberRes, err error) {
	out = &v1.DeleteMemberRes{}

	// 检查会员信息表是否存在
	count, err := dao.Member.Ctx(ctx).Where(dao.Member.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询会员信息表失败")
	}
	if count == 0 {
		return nil, gerror.New("会员信息表不存在")
	}

	// 删除数据
	_, err = dao.Member.Ctx(ctx).Where(dao.Member.Columns().Id, in.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除会员信息表失败")
	}

	return
}



// BatchDeleteMember 批量删除会员信息表
func (s *sMember) BatchDeleteMember(ctx context.Context, in v1.BatchDeleteMemberReq) (out *v1.BatchDeleteMemberRes, err error) {
	out = &v1.BatchDeleteMemberRes{}

	if len(in.Ids) == 0 {
		return nil, gerror.New("请选择要删除的会员信息表")
	}

	// 批量删除
	_, err = dao.Member.Ctx(ctx).WhereIn(dao.Member.Columns().Id, in.Ids).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除会员信息表失败")
	}

	return
}
