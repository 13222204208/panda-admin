package curd

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

// FrontendGenerator 前端代码生成器
type FrontendGenerator struct {
	generator *CurdGenerator
}

// NewFrontendGenerator 创建新的前端生成器
func NewFrontendGenerator() *FrontendGenerator {
	return &FrontendGenerator{
		generator: NewCurdGenerator(),
	}
}

// FrontendConfig 前端生成配置
type FrontendConfig struct {
	GenerateConfig
	WebRoot string `json:"web_root"` // 前端项目根目录
}

// GenerateAll 生成所有前端文件
func (fg *FrontendGenerator) GenerateAll(ctx context.Context, config FrontendConfig) error {
	// 生成API文件
	if err := fg.GenerateAPI(ctx, config); err != nil {
		return fmt.Errorf("生成API文件失败: %v", err)
	}

	// 生成主页面文件
	if err := fg.GenerateIndexPage(ctx, config); err != nil {
		return fmt.Errorf("生成主页面文件失败: %v", err)
	}

	// 生成表单组件
	if err := fg.GenerateFormComponent(ctx, config); err != nil {
		return fmt.Errorf("生成表单组件失败: %v", err)
	}

	// 生成工具文件
	if err := fg.GenerateUtils(ctx, config); err != nil {
		return fmt.Errorf("生成工具文件失败: %v", err)
	}

	g.Log().Infof(ctx, "前端文件生成完成: %s", config.ModuleName)
	return nil
}

// GenerateAPI 生成TypeScript API文件
func (fg *FrontendGenerator) GenerateAPI(ctx context.Context, config FrontendConfig) error {
	templatePath := fg.getTemplatePath("api.ts.template")
	outputPath := fg.getAPIOutputPath(config)
	templateData := fg.generator.prepareTemplateData(config.GenerateConfig)

	return fg.generator.generateFromTemplate(ctx, templatePath, outputPath, templateData)
}

// GenerateIndexPage 生成主页面Vue文件
func (fg *FrontendGenerator) GenerateIndexPage(ctx context.Context, config FrontendConfig) error {
	templatePath := fg.getTemplatePath("index.vue.template")
	outputPath := fg.getIndexPageOutputPath(config)
	templateData := fg.generator.prepareTemplateData(config.GenerateConfig)

	return fg.generator.generateFromTemplate(ctx, templatePath, outputPath, templateData)
}

// GenerateFormComponent 生成表单组件
func (fg *FrontendGenerator) GenerateFormComponent(ctx context.Context, config FrontendConfig) error {
	templatePath := fg.getTemplatePath("form/index.vue.template")
	outputPath := fg.getFormComponentOutputPath(config)
	templateData := fg.generator.prepareTemplateData(config.GenerateConfig)

	return fg.generator.generateFromTemplate(ctx, templatePath, outputPath, templateData)
}

// GenerateUtils 生成工具文件
func (fg *FrontendGenerator) GenerateUtils(ctx context.Context, config FrontendConfig) error {
	// 生成types.ts
	if err := fg.generateTypesFile(ctx, config); err != nil {
		return err
	}
	g.Log().Infof(ctx, "生成utils/types.ts文件完成: %v", config.Options.BatchDelete)
	// 生成hook.tsx
	if err := fg.generateHookFile(ctx, config); err != nil {
		return err
	}

	// 生成rule.ts
	if err := fg.generateRuleFile(ctx, config); err != nil {
		return err
	}

	return nil
}

// generateTypesFile 生成types.ts文件
func (fg *FrontendGenerator) generateTypesFile(ctx context.Context, config FrontendConfig) error {
	templatePath := fg.getTemplatePath("utils/types.ts.template")
	outputPath := fg.getTypesOutputPath(config)
	templateData := fg.generator.prepareTemplateData(config.GenerateConfig)

	return fg.generator.generateFromTemplate(ctx, templatePath, outputPath, templateData)
}

// generateHookFile 生成hook.tsx文件
func (fg *FrontendGenerator) generateHookFile(ctx context.Context, config FrontendConfig) error {
	templatePath := fg.getTemplatePath("utils/hook.tsx.template")
	outputPath := fg.getHookOutputPath(config)
	templateData := fg.generator.prepareTemplateData(config.GenerateConfig)

	return fg.generator.generateFromTemplate(ctx, templatePath, outputPath, templateData)
}

// generateRuleFile 生成rule.ts文件
func (fg *FrontendGenerator) generateRuleFile(ctx context.Context, config FrontendConfig) error {
	templatePath := fg.getTemplatePath("utils/rule.ts.template")
	outputPath := fg.getRuleOutputPath(config)
	templateData := fg.generator.prepareTemplateData(config.GenerateConfig)

	return fg.generator.generateFromTemplate(ctx, templatePath, outputPath, templateData)
}

// getTemplatePath 获取模板文件路径
func (fg *FrontendGenerator) getTemplatePath(templateName string) string {
	return fg.generator.getTemplatePath(templateName)
}

// getWebRoot 获取前端项目根目录
func (fg *FrontendGenerator) getWebRoot(config FrontendConfig) string {
	if config.WebRoot != "" {
		return config.WebRoot
	}
	// 默认前端项目路径
	projectRoot := fg.generator.getProjectRoot()
	return filepath.Join(filepath.Dir(projectRoot), "web")
}

// getAPIOutputPath 获取API文件输出路径
func (fg *FrontendGenerator) getAPIOutputPath(config FrontendConfig) string {
	webRoot := fg.getWebRoot(config)
	return filepath.Join(webRoot, "src", "api", config.ModuleName+".ts")
}

// getIndexPageOutputPath 获取主页面输出路径
func (fg *FrontendGenerator) getIndexPageOutputPath(config FrontendConfig) string {
	webRoot := fg.getWebRoot(config)
	return filepath.Join(webRoot, "src", "views", config.PackageName, "index.vue")
}

// getFormComponentOutputPath 获取表单组件输出路径
func (fg *FrontendGenerator) getFormComponentOutputPath(config FrontendConfig) string {
	webRoot := fg.getWebRoot(config)
	return filepath.Join(webRoot, "src", "views", config.PackageName, "form", "index.vue")
}

// getTypesOutputPath 获取types文件输出路径
func (fg *FrontendGenerator) getTypesOutputPath(config FrontendConfig) string {
	webRoot := fg.getWebRoot(config)
	return filepath.Join(webRoot, "src", "views", config.PackageName, "utils", "types.ts")
}

// getHookOutputPath 获取hook文件输出路径
func (fg *FrontendGenerator) getHookOutputPath(config FrontendConfig) string {
	webRoot := fg.getWebRoot(config)
	return filepath.Join(webRoot, "src", "views", config.PackageName, "utils", "hook.tsx")
}

// getRuleOutputPath 获取rule文件输出路径
func (fg *FrontendGenerator) getRuleOutputPath(config FrontendConfig) string {
	webRoot := fg.getWebRoot(config)
	return filepath.Join(webRoot, "src", "views", config.PackageName, "utils", "rule.ts")
}

// GenerateSingleFile 生成单个前端文件
func (fg *FrontendGenerator) GenerateSingleFile(ctx context.Context, config FrontendConfig, fileType string) error {
	switch strings.ToLower(fileType) {
	case "api":
		return fg.GenerateAPI(ctx, config)
	case "index", "page":
		return fg.GenerateIndexPage(ctx, config)
	case "form":
		return fg.GenerateFormComponent(ctx, config)
	case "types":
		return fg.generateTypesFile(ctx, config)
	case "hook":
		return fg.generateHookFile(ctx, config)
	case "rule":
		return fg.generateRuleFile(ctx, config)
	default:
		return fmt.Errorf("不支持的文件类型: %s", fileType)
	}
}

// ValidateConfig 验证前端生成配置
func (fg *FrontendGenerator) ValidateConfig(config FrontendConfig) error {
	if config.ModuleName == "" {
		return fmt.Errorf("模块名不能为空")
	}
	if config.PackageName == "" {
		return fmt.Errorf("包名不能为空")
	}
	if config.EntityName == "" {
		return fmt.Errorf("实体名不能为空")
	}
	if len(config.Columns) == 0 {
		return fmt.Errorf("字段配置不能为空")
	}
	return nil
}

// GetGeneratedFiles 获取将要生成的文件列表
func (fg *FrontendGenerator) GetGeneratedFiles(config FrontendConfig) []string {
	return []string{
		fg.getAPIOutputPath(config),
		fg.getIndexPageOutputPath(config),
		fg.getFormComponentOutputPath(config),
		fg.getTypesOutputPath(config),
		fg.getHookOutputPath(config),
		fg.getRuleOutputPath(config),
	}
}
