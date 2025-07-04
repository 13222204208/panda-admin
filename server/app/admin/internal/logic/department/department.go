package department

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "server/app/admin/api/department/v1"
	"server/app/admin/internal/dao"
)

type sDepartment struct{}

func New() *sDepartment {
	return &sDepartment{}
}

// Create 创建部门
func (s *sDepartment) Create(ctx context.Context, in v1.CreateReq) (out *v1.CreateRes, err error) {
	out = &v1.CreateRes{}

	// 验证必填字段
	if in.Name == nil || *in.Name == "" {
		return nil, gerror.New("部门名称不能为空")
	}

	// 检查部门名称在同级下是否唯一
	parentId := uint64(0)
	if in.ParentId != nil {
		parentId = *in.ParentId
	}
	count, err := dao.Department.Ctx(ctx).Where(dao.Department.Columns().Name, *in.Name).Where(dao.Department.Columns().ParentId, parentId).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询部门名称失败")
	}
	if count > 0 {
		return nil, gerror.Newf("同级下部门名称 %s 已存在", *in.Name)
	}

	// 构建插入数据
	data := g.Map{
		dao.Department.Columns().ParentId: parentId,
		dao.Department.Columns().Name:     *in.Name,
		dao.Department.Columns().Status:   gconv.Int(in.Status),
	}

	// 可选字段
	if in.Principal != nil {
		data[dao.Department.Columns().Principal] = *in.Principal
	}
	if in.Phone != nil {
		data[dao.Department.Columns().Phone] = *in.Phone
	}
	if in.Email != nil {
		data[dao.Department.Columns().Email] = *in.Email
	}
	if in.Sort != nil {
		data[dao.Department.Columns().Sort] = *in.Sort
	}
	if in.Remark != nil {
		data[dao.Department.Columns().Remark] = *in.Remark
	}

	// 插入数据
	id, err := dao.Department.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建部门失败")
	}

	out.Id = uint64(id)
	return
}

// GetList 获取部门列表
func (s *sDepartment) GetList(ctx context.Context, in v1.GetListReq) (out *v1.GetListRes, err error) {
	out = &v1.GetListRes{}

	// 构建查询条件
	m := dao.Department.Ctx(ctx)
	if in.ParentId != nil {
		m = m.Where(dao.Department.Columns().ParentId, *in.ParentId)
	}
	if in.Name != "" {
		m = m.WhereLike(dao.Department.Columns().Name, "%"+in.Name+"%")
	}
	if in.Status != nil {
		m = m.Where(dao.Department.Columns().Status, *in.Status)
	}

	// 查询列表（按排序号和创建时间排序，取消分页）
	err = m.OrderAsc(dao.Department.Columns().Sort).OrderAsc(dao.Department.Columns().CreatedAt).Scan(&out.List)
	if err != nil {
		return nil, gerror.Wrap(err, "查询部门列表失败")
	}
	return
}

// Update 更新部门
func (s *sDepartment) Update(ctx context.Context, in v1.UpdateReq) (err error) {
	// 检查部门是否存在
	count, err := dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询部门失败")
	}
	if count == 0 {
		return gerror.Newf("部门ID %d 不存在", in.Id)
	}

	// 动态构建更新数据
	updateData := g.Map{}

	// 检查并添加需要更新的字段
	if in.ParentId != nil {
		// 检查是否会形成循环引用
		if *in.ParentId == in.Id {
			return gerror.New("不能将部门设置为自己的子部门")
		}
		updateData[dao.Department.Columns().ParentId] = *in.ParentId
	}

	if in.Name != nil {
		if *in.Name == "" {
			return gerror.New("部门名称不能为空")
		}
		// 检查同级下名称唯一性
		parentId := uint64(0)
		if in.ParentId != nil {
			parentId = *in.ParentId
		} else {
			// 获取当前部门的父级ID
			var currentDept struct {
				ParentId uint64 `json:"parentId"`
			}
			err = dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, in.Id).Scan(&currentDept)
			if err != nil {
				return gerror.Wrap(err, "查询当前部门信息失败")
			}
			parentId = currentDept.ParentId
		}
		count, err = dao.Department.Ctx(ctx).Where(dao.Department.Columns().Name, *in.Name).Where(dao.Department.Columns().ParentId, parentId).WhereNot(dao.Department.Columns().Id, in.Id).Count()
		if err != nil {
			return gerror.Wrap(err, "查询部门名称失败")
		}
		if count > 0 {
			return gerror.Newf("同级下部门名称 %s 已存在", *in.Name)
		}
		updateData[dao.Department.Columns().Name] = *in.Name
	}

	if in.Principal != nil {
		updateData[dao.Department.Columns().Principal] = *in.Principal
	}

	if in.Phone != nil {
		updateData[dao.Department.Columns().Phone] = *in.Phone
	}

	if in.Email != nil {
		updateData[dao.Department.Columns().Email] = *in.Email
	}

	if in.Sort != nil {
		updateData[dao.Department.Columns().Sort] = *in.Sort
	}

	if in.Status != nil {
		updateData[dao.Department.Columns().Status] = *in.Status
	}

	if in.Remark != nil {
		updateData[dao.Department.Columns().Remark] = *in.Remark
	}

	// 检查是否有字段需要更新
	if len(updateData) == 0 {
		return gerror.New("没有需要更新的字段")
	}

	// 更新数据
	_, err = dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, in.Id).Data(updateData).Update()
	return gerror.Wrap(err, "更新部门失败")
}

// Delete 删除部门
func (s *sDepartment) Delete(ctx context.Context, in v1.DeleteReq) (err error) {
	// 检查部门是否存在
	count, err := dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询部门失败")
	}
	if count == 0 {
		return gerror.Newf("部门ID %d 不存在", in.Id)
	}

	// 检查是否有子部门
	childCount, err := dao.Department.Ctx(ctx).Where(dao.Department.Columns().ParentId, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "查询子部门失败")
	}
	if childCount > 0 {
		return gerror.New("该部门下还有子部门，无法删除")
	}

	// 删除数据
	_, err = dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, in.Id).Delete()
	return gerror.Wrap(err, "删除部门失败")
}
