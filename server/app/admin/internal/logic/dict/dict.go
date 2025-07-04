package dict

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"

	"server/app/admin/api/common/page"
	v1 "server/app/admin/api/dict/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/model/entity"
)

type sDict struct{}

func New() *sDict {
	return &sDict{}
}

// GetList 获取字典列表
func (s *sDict) GetList(ctx context.Context, in v1.GetListReq) (out *v1.GetListRes, err error) {
	out = &v1.GetListRes{}

	m := dao.Dict.Ctx(ctx)

	// 构建查询条件
	if in.DictType != "" {
		m = m.WhereLike(dao.Dict.Columns().DictType, "%"+in.DictType+"%")
	}
	if in.DictLabel != "" {
		m = m.WhereLike(dao.Dict.Columns().DictLabel, "%"+in.DictLabel+"%")
	}
	if in.Status != nil {
		m = m.Where(dao.Dict.Columns().Status, *in.Status)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典总数失败")
	}

	// 分页查询
	var list []entity.Dict
	err = m.Page(in.CurrentPage, in.PageSize).
		OrderDesc(dao.Dict.Columns().CreatedAt).
		Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典列表失败")
	}

	// 转换数据格式
	var dictList []v1.DictInfo
	for _, item := range list {
		dictList = append(dictList, v1.DictInfo{
			Id:        item.Id,
			Title:     item.Title,
			DictType:  item.DictType,
			DictLabel: item.DictLabel,
			DictValue: item.DictValue,
			Sort:      item.Sort,
			Status:    item.Status,
			Remark:    item.Remark,
			CreatedAt: item.CreatedAt,
		})
	}

	out.ResPage = page.ResPage{
		Total:       int(total),
		CurrentPage: in.CurrentPage,
	}
	out.List = dictList
	return
}

// Update 更新字典
func (s *sDict) Update(ctx context.Context, in v1.UpdateReq) (out *v1.UpdateRes, err error) {
	out = &v1.UpdateRes{}

	// 检查字典是否存在
	count, err := dao.Dict.Ctx(ctx).Where(dao.Dict.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典失败")
	}
	if count == 0 {
		return nil, gerror.New("字典不存在")
	}

	// 检查字典值是否已存在（排除当前记录）
	if in.DictType != nil && in.DictValue != nil {
		count, err = dao.Dict.Ctx(ctx).
			Where(dao.Dict.Columns().DictType, *in.DictType).
			Where(dao.Dict.Columns().DictValue, *in.DictValue).
			WhereNot(dao.Dict.Columns().Id, in.Id).
			Count()
		if err != nil {
			return nil, gerror.Wrap(err, "查询字典值失败")
		}
		if count > 0 {
			return nil, gerror.New("该字典类型下的字典值已存在")
		}
	}

	// 动态构建更新数据
	updateData := g.Map{}

	// 检查并添加需要更新的字段
	if in.Title != nil {
		updateData[dao.Dict.Columns().Title] = *in.Title
	}
	if in.DictType != nil {
		updateData[dao.Dict.Columns().DictType] = *in.DictType
	}
	if in.DictLabel != nil {
		updateData[dao.Dict.Columns().DictLabel] = *in.DictLabel
	}
	if in.DictValue != nil {
		updateData[dao.Dict.Columns().DictValue] = *in.DictValue
	}
	if in.Sort != nil {
		updateData[dao.Dict.Columns().Sort] = *in.Sort
	}
	if in.Status != nil {
		updateData[dao.Dict.Columns().Status] = *in.Status
	}
	if in.Remark != nil {
		updateData[dao.Dict.Columns().Remark] = *in.Remark
	}

	// 更新数据
	_, err = dao.Dict.Ctx(ctx).
		Where(dao.Dict.Columns().Id, in.Id).
		Data(updateData).
		Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新字典失败")
	}

	return
}

// Delete 删除字典
func (s *sDict) Delete(ctx context.Context, in v1.DeleteReq) (out *v1.DeleteRes, err error) {
	out = &v1.DeleteRes{}

	// 检查字典是否存在
	count, err := dao.Dict.Ctx(ctx).Where(dao.Dict.Columns().Id, in.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典失败")
	}
	if count == 0 {
		return nil, gerror.New("字典不存在")
	}

	// 删除数据
	_, err = dao.Dict.Ctx(ctx).Where(dao.Dict.Columns().Id, in.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除字典失败")
	}

	return
}

// BatchDelete 批量删除字典
func (s *sDict) BatchDelete(ctx context.Context, in v1.BatchDeleteReq) (out *v1.BatchDeleteRes, err error) {
	out = &v1.BatchDeleteRes{}

	if len(in.Ids) == 0 {
		return nil, gerror.New("请选择要删除的字典")
	}

	// 批量删除
	_, err = dao.Dict.Ctx(ctx).WhereIn(dao.Dict.Columns().Id, in.Ids).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除字典失败")
	}

	return
}

// BatchCreate 批量创建字典
func (s *sDict) BatchCreate(ctx context.Context, in v1.BatchCreateReq) (out *v1.BatchCreateRes, err error) {
	out = &v1.BatchCreateRes{}

	if len(in.DictItems) == 0 {
		return nil, gerror.New("请添加字典项")
	}

	// 检查字典值是否重复
	valueMap := make(map[string]bool)
	for _, item := range in.DictItems {
		if valueMap[item.DictValue] {
			return nil, gerror.New(fmt.Sprintf("字典值 %s 重复", item.DictValue))
		}
		valueMap[item.DictValue] = true
	}

	// 检查数据库中是否已存在相同的字典值
	var existValues []string
	for _, item := range in.DictItems {
		existValues = append(existValues, item.DictValue)
	}

	var existRecords []entity.Dict
	err = dao.Dict.Ctx(ctx).
		Where(dao.Dict.Columns().DictType, in.DictType).
		WhereIn(dao.Dict.Columns().DictValue, existValues).
		Scan(&existRecords)
	if err != nil {
		return nil, gerror.Wrap(err, "查询已存在字典值失败")
	}

	if len(existRecords) > 0 {
		return nil, gerror.New(fmt.Sprintf("字典值 %s 已存在", existRecords[0].DictValue))
	}

	// 使用事务批量插入
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, item := range in.DictItems {
			// 构建插入数据
			data := g.Map{
				dao.Dict.Columns().Title:     in.Title,
				dao.Dict.Columns().DictType:  in.DictType,
				dao.Dict.Columns().DictLabel: item.DictLabel,
				dao.Dict.Columns().DictValue: item.DictValue,
				dao.Dict.Columns().Sort:      gconv.Int(gutil.GetOrDefaultAny(item.Sort, 0)),
				dao.Dict.Columns().Status:    1, // 默认启用
			}

			// 可选字段
			if item.Remark != nil {
				data[dao.Dict.Columns().Remark] = *item.Remark
			}

			_, err = dao.Dict.Ctx(ctx).TX(tx).Data(data).Insert()
			if err != nil {
				return gerror.Wrap(err, "批量插入字典失败")
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return
}

// GetOptions 根据字典类型获取字典选项
func (s *sDict) GetOptions(ctx context.Context, in v1.GetOptionsReq) (out *v1.GetOptionsRes, err error) {
	out = &v1.GetOptionsRes{}

	var list []entity.Dict
	err = dao.Dict.Ctx(ctx).
		Where(dao.Dict.Columns().DictType, in.DictType).
		Where(dao.Dict.Columns().Status, 1). // 只获取启用的字典
		OrderAsc(dao.Dict.Columns().Sort).
		OrderAsc(dao.Dict.Columns().Id).
		Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询字典选项失败")
	}

	// 转换数据格式
	var options []v1.DictOption
	for _, item := range list {
		options = append(options, v1.DictOption{
			Label: item.DictLabel,
			Value: item.DictValue,
		})
	}

	out.Options = options
	return
}

// GetDistinctTypes 获取不重复的字典类型和标题
func (s *sDict) GetDistinctTypes(ctx context.Context, in v1.GetDistinctTypesReq) (out *v1.GetDistinctTypesRes, err error) {
	out = &v1.GetDistinctTypesRes{}

	// 使用DISTINCT查询获取不重复的title和dictType组合
	var results []struct {
		Title    string `json:"title"`
		DictType string `json:"dictType"`
	}

	err = dao.Dict.Ctx(ctx).
		Fields("DISTINCT title, dict_type").
		Where("title IS NOT NULL AND title != ''").
		Where("dict_type IS NOT NULL AND dict_type != ''").
		OrderAsc("dict_type, title").
		Scan(&results)

	if err != nil {
		return nil, gerror.Wrap(err, "查询不重复的字典类型和标题失败")
	}

	// 转换数据格式
	var types []v1.DictTypeInfo
	for _, item := range results {
		types = append(types, v1.DictTypeInfo{
			Title:    item.Title,
			DictType: item.DictType,
		})
	}

	out.Types = types
	return
}
