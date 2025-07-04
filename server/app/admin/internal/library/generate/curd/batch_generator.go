package curd

import (
	"context"
	"fmt"
)

// BatchGenerator 批量生成器
type BatchGenerator struct {
	generator *CurdGenerator
	templates map[string]string // 模板类型 -> 模板路径
}

// func NewBatchGenerator() *BatchGenerator {
// 	return &BatchGenerator{
// 		generator: NewCurdGenerator(),
// 		templates: map[string]string{
// 			"api":        "/Users/ypp/开源/panda/server/app/admin/resource/generate/curd/api.go.template",
// 			"controller": "/Users/ypp/开源/panda/server/app/admin/resource/generate/curd/controller.go.template",
// 			// 可以添加更多模板类型
// 		},
// 	}
// }

// GenerateAll 生成所有类型的文件
func (bg *BatchGenerator) GenerateAll(ctx context.Context, config GenerateConfig) error {
	for templateType, templatePath := range bg.templates {
		outputPath := bg.getOutputPath(templateType, config)
		templateData := bg.generator.prepareTemplateData(config)

		if err := bg.generator.generateFromTemplate(ctx, templatePath, outputPath, templateData); err != nil {
			return fmt.Errorf("生成 %s 失败: %v", templateType, err)
		}
	}
	return nil
}

// getOutputPath 根据模板类型获取输出路径
func (bg *BatchGenerator) getOutputPath(templateType string, config GenerateConfig) string {
	switch templateType {
	case "api":
		return fmt.Sprintf("/Users/ypp/开源/panda/server/app/admin/api/%s/v1/%s.go",
			config.PackageName, config.ModuleName)
	case "controller":
		return fmt.Sprintf("/Users/ypp/开源/panda/server/app/admin/internal/controller/%s/%s.go",
			config.PackageName, config.ModuleName)
	default:
		return ""
	}
}
