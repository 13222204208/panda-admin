package generate

import (
	"context"
	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetCodeGenRecordList 获取代码生成记录列表
func (s *sGenerate) GetCodeGenRecordList(ctx context.Context, req v1.GetCodeGenRecordListReq) (res *v1.GetCodeGenRecordListRes, err error) {
	res = &v1.GetCodeGenRecordListRes{}

	// 构建查询条件
	m := dao.CodeGenRecord.Ctx(ctx)

	if req.TableName != "" {
		m = m.WhereLike(dao.CodeGenRecord.Columns().TableName, "%"+req.TableName+"%")
	}
	if req.ModuleName != "" {
		m = m.WhereLike(dao.CodeGenRecord.Columns().ModuleName, "%"+req.ModuleName+"%")
	}
	if req.Status != nil {
		m = m.Where(dao.CodeGenRecord.Columns().Status, *req.Status)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询代码生成记录总数失败")
	}
	res.Total = total

	// 分页查询
	var list []entity.CodeGenRecord
	err = m.Page(req.CurrentPage, req.PageSize).
		OrderDesc(dao.CodeGenRecord.Columns().CreatedAt).
		Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询代码生成记录列表失败")
	}

	// 转换数据格式
	if err := gconv.Scan(list, &res.List); err != nil {
		return nil, gerror.Wrap(err, "数据转换失败")
	}

	res.CurrentPage = req.CurrentPage
	return res, nil
}

// GetCodeGenRecordDetail 获取代码生成记录详情
func (s *sGenerate) GetCodeGenRecordDetail(ctx context.Context, req v1.GetCodeGenRecordDetailReq) (res *v1.GetCodeGenRecordDetailRes, err error) {
	res = &v1.GetCodeGenRecordDetailRes{}

	// 查询记录详情
	var record entity.CodeGenRecord
	err = dao.CodeGenRecord.Ctx(ctx).
		Where(dao.CodeGenRecord.Columns().Id, req.Id).
		Scan(&record)
	if err != nil {
		return nil, gerror.Wrap(err, "查询代码生成记录详情失败")
	}

	// 检查记录是否存在
	if record.Id == 0 {
		return nil, gerror.New("代码生成记录不存在")
	}

	// 转换数据格式
	if err := gconv.Scan(record, res); err != nil {
		return nil, gerror.Wrap(err, "数据转换失败")
	}

	return res, nil
}

// DeleteCodeGenRecord 删除代码生成记录
func (s *sGenerate) DeleteCodeGenRecord(ctx context.Context, req v1.DeleteCodeGenRecordReq) (res *v1.DeleteCodeGenRecordRes, err error) {
	res = &v1.DeleteCodeGenRecordRes{}

	// 检查记录是否存在
	count, err := dao.CodeGenRecord.Ctx(ctx).Where(dao.CodeGenRecord.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询代码生成记录失败")
	}
	if count == 0 {
		return nil, gerror.New("代码生成记录不存在")
	}

	// 删除记录
	_, err = dao.CodeGenRecord.Ctx(ctx).Where(dao.CodeGenRecord.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除代码生成记录失败")
	}

	return
}

// UpdateCodeGenRecord 更新代码生成记录
func (s *sGenerate) UpdateCodeGenRecord(ctx context.Context, req v1.UpdateCodeGenRecordReq) (res *v1.UpdateCodeGenRecordRes, err error) {
	res = &v1.UpdateCodeGenRecordRes{}

	// 检查记录是否存在
	count, err := dao.CodeGenRecord.Ctx(ctx).Where(dao.CodeGenRecord.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询代码生成记录失败")
	}
	if count == 0 {
		return nil, gerror.New("代码生成记录不存在")
	}

	// 更新记录
	_, err = dao.CodeGenRecord.Ctx(ctx).
		Where(dao.CodeGenRecord.Columns().Id, req.Id).
		Data(g.Map{
			dao.CodeGenRecord.Columns().TableName:    req.TableName,
			dao.CodeGenRecord.Columns().TableComment: req.TableComment,
			dao.CodeGenRecord.Columns().PackageName:  req.PackageName,
			dao.CodeGenRecord.Columns().ModuleName:   req.ModuleName,
			dao.CodeGenRecord.Columns().Options:      req.Options,
			dao.CodeGenRecord.Columns().Columns:      req.Columns,
		}).
		Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新代码生成记录失败")
	}

	return
}
