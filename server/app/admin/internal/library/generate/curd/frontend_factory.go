package curd

import (
	"context"
	"server/app/admin/internal/consts"
)

// FrontendGeneratorFactory 前端生成器工厂
type FrontendGeneratorFactory struct{}

// NewFrontendGeneratorFactory 创建前端生成器工厂
func NewFrontendGeneratorFactory() *FrontendGeneratorFactory {
	return &FrontendGeneratorFactory{}
}

// CreateFrontendConfig 创建前端配置
func (f *FrontendGeneratorFactory) CreateFrontendConfig(
	tableName, tableComment, entityName, packageName, moduleName string,
	columns []Column,
	options *consts.GenerateOptions,
	webRoot string,
) FrontendConfig {
	return FrontendConfig{
		GenerateConfig: GenerateConfig{
			TableName:    tableName,
			TableComment: tableComment,
			EntityName:   entityName,
			PackageName:  packageName,
			ModuleName:   moduleName,
			Columns:      columns,
			Options:      options,
		},
		WebRoot: webRoot,
	}
}

// QuickGenerate 快速生成前端文件
func (f *FrontendGeneratorFactory) QuickGenerate(
	ctx context.Context,
	tableName, tableComment, entityName, packageName, moduleName string,
	columns []Column,
	options *consts.GenerateOptions,
) error {
	config := f.CreateFrontendConfig(
		tableName, tableComment, entityName, packageName, moduleName,
		columns, options, "", // 使用默认web根目录
	)

	generator := NewFrontendGenerator()
	return generator.GenerateAll(ctx, config)
}

// QuickGenerateWithCustomWebRoot 使用自定义web根目录快速生成
func (f *FrontendGeneratorFactory) QuickGenerateWithCustomWebRoot(
	ctx context.Context,
	tableName, tableComment, entityName, packageName, moduleName, webRoot string,
	columns []Column,
	options *consts.GenerateOptions,
) error {
	config := f.CreateFrontendConfig(
		tableName, tableComment, entityName, packageName, moduleName,
		columns, options, webRoot,
	)

	generator := NewFrontendGenerator()
	return generator.GenerateAll(ctx, config)
}