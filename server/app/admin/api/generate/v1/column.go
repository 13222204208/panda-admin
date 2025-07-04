package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FormModeOption 表单模式选项
type FormModeOption struct {
	Label string `json:"label" dc:"显示名称"`
	Value string `json:"value" dc:"选项值"`
}

// FormRoleOption 表单验证选项
type FormRoleOption struct {
	Label string `json:"label" dc:"显示名称"`
	Value string `json:"value" dc:"选项值"`
}

// WhereModeOption 查询条件选项
type WhereModeOption struct {
	Label string `json:"label" dc:"显示名称"`
	Value string `json:"value" dc:"选项值"`
}

// GetColumnConfigOptionsReq 获取字段配置选项请求
type GetColumnConfigOptionsReq struct {
	g.Meta `path:"/generate/column/options" method:"get" tags:"代码生成" summary:"获取字段配置选项"`
}

// GetColumnConfigOptionsRes 获取字段配置选项响应
type GetColumnConfigOptionsRes struct {
	FormModes        []FormModeOption     `json:"formModes" dc:"表单模式选项"`
	FormValidations  []FormRoleOption     `json:"formValidations" dc:"表单验证选项"`
	WhereModes       []WhereModeOption    `json:"whereModes" dc:"查询条件选项"`
}