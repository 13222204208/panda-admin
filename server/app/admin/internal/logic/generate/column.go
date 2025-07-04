package generate

import (
	"context"
	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/library/generate"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetColumnConfigOptions 获取字段配置选项
func (s *sGenerate) GetColumnConfigOptions(ctx context.Context, req v1.GetColumnConfigOptionsReq) (res *v1.GetColumnConfigOptionsRes, err error) {
	res = &v1.GetColumnConfigOptionsRes{}

	// 获取表单模式选项
	formModeOptions := generate.GetFormModeOptions()
	if err := gconv.Scan(formModeOptions, &res.FormModes); err != nil {
		return nil, err
	}

	// 获取表单验证选项
	formValidationOptions := generate.GetFormValidationOptions()
	if err := gconv.Scan(formValidationOptions, &res.FormValidations); err != nil {
		return nil, err
	}

	// 获取查询条件选项
	whereModeOptions := generate.GetWhereModeOptions()
	if err := gconv.Scan(whereModeOptions, &res.WhereModes); err != nil {
		return nil, err
	}

	return res, nil
}
