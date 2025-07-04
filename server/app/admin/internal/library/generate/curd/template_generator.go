package curd

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"server/app/admin/internal/consts"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
)

// Column 字段配置结构体
type Column struct {
	ColumnName     string `json:"columnName"`     // 数据库字段名
	ColumnComment  string `json:"columnComment"`  // 字段注释
	GoField        string `json:"goField"`        // Go字段名
	GoType         string `json:"goType"`         // Go类型
	JsonField      string `json:"jsonField"`      // JSON字段名
	IsRequired     bool   `json:"isRequired"`     // 是否必填
	IsQuery        bool   `json:"isQuery"`        // 是否用于查询
	IsList         bool   `json:"isList"`         // 是否显示在列表
	IsPointer      bool   `json:"isPointer"`      // 是否指针类型
	IsUnique       bool   `json:"isUnique"`       // 是否唯一
	ValidationRule string `json:"validationRule"` // 验证规则
}

// CurdGenerator CURD代码生成器
type CurdGenerator struct {
	view *gview.View
}

// NewCurdGenerator 创建新的CURD生成器
func NewCurdGenerator() *CurdGenerator {
	return &CurdGenerator{
		view: gview.New(),
	}
}

// getProjectRoot 获取项目根目录
func (cg *CurdGenerator) getProjectRoot() string {
	// 获取当前文件的绝对路径
	_, currentFile, _, _ := runtime.Caller(0)
	// 从当前文件路径向上查找到server目录
	currentDir := filepath.Dir(currentFile)

	// 向上查找直到找到包含go.mod的目录
	for {
		if gfile.Exists(filepath.Join(currentDir, "go.mod")) {
			return currentDir
		}
		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			// 已经到达根目录，未找到go.mod
			break
		}
		currentDir = parent
	}

	// 如果未找到go.mod，返回当前工作目录
	if workDir := gfile.Pwd(); workDir != "" {
		return workDir
	}

	// 最后的备选方案
	return "/Users/ypp/开源/panda/server"
}

// getTemplatePath 获取模板文件路径
func (cg *CurdGenerator) getTemplatePath(templateName string) string {
	projectRoot := cg.getProjectRoot()
	return filepath.Join(projectRoot, "app", "admin", "resource", "generate", "curd", templateName)
}

// getOutputPath 获取输出文件路径
func (cg *CurdGenerator) getOutputPath(packageName, moduleName string) string {
	projectRoot := cg.getProjectRoot()
	return filepath.Join(projectRoot, "app", "admin", "api", packageName, "v1", moduleName+".go")
}

// GenerateAPI 生成API文件
func (cg *CurdGenerator) GenerateAPI(ctx context.Context, config GenerateConfig) error {
	// 动态获取模板文件路径
	templatePath := cg.getTemplatePath("api.go.template")
	fmt.Println(templatePath, "templatePath路径")
	// 动态获取输出文件路径
	outputPath := cg.getOutputPath(config.PackageName, config.ModuleName)
	// 准备模板数据
	templateData := cg.prepareTemplateData(config)
	// 生成代码
	return cg.generateFromTemplate(ctx, templatePath, outputPath, templateData)
}

// GenerateConfig 生成配置
type GenerateConfig struct {
	TableName    string                  `json:"table_name"`    // 表名
	TableComment string                  `json:"table_comment"` // 表注释
	EntityName   string                  `json:"entity_name"`   // 实体名称
	PackageName  string                  `json:"package_name"`  // 包名
	ModuleName   string                  `json:"module_name"`   // 模块名
	Columns      []Column                `json:"columns"`       // 字段配置
	Options      *consts.GenerateOptions `json:"options"`       // 生成选项
}

// prepareTemplateData 准备模板数据
func (cg *CurdGenerator) prepareTemplateData(config GenerateConfig) g.Map {
	return g.Map{
		"EntityName":   config.EntityName,
		"TableComment": config.TableComment,
		"ModuleName":   config.ModuleName,
		"PackageName":  config.PackageName,
		"Columns":      config.Columns,
		"Options":      config.Options, // 添加选项到模板数据
	}
}

// generateFromTemplate 从模板生成文件
func (cg *CurdGenerator) generateFromTemplate(ctx context.Context, templatePath, outputPath string, data g.Map) error {
	// 读取模板内容
	templateContent := gfile.GetContents(templatePath)
	if templateContent == "" {
		return fmt.Errorf("无法读取模板文件: %s", templatePath)
	}
	// g.Log().Info(ctx, templateContent, "模板内容")
	// g.Log().Info(ctx, data, "模板数据")
	// 解析模板
	result, err := cg.view.ParseContent(ctx, templateContent, data)
	if err != nil {
		return fmt.Errorf("模板解析失败: %v", err)
	}

	// 确保输出目录存在
	outputDir := filepath.Dir(outputPath)
	if !gfile.Exists(outputDir) {
		if err := gfile.Mkdir(outputDir); err != nil {
			return fmt.Errorf("创建输出目录失败: %v", err)
		}
	}

	// 写入文件 - 这里会直接覆盖已存在的文件
	if err := gfile.PutContents(outputPath, result); err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	g.Log().Infof(ctx, "文件生成成功: %s", outputPath)
	return nil
}

// GenerateLogic 生成Logic文件
func (cg *CurdGenerator) GenerateLogic(ctx context.Context, config GenerateConfig) error {
	// 动态获取模板文件路径
	templatePath := cg.getTemplatePath("logic.go.template")
	fmt.Println(templatePath, "logic模板路径")

	// 动态获取输出文件路径
	outputPath := cg.getLogicOutputPath(config.PackageName)
	fmt.Println(outputPath, "logic输出路径")

	// 准备模板数据
	templateData := cg.prepareTemplateData(config)

	// 生成代码
	return cg.generateFromTemplate(ctx, templatePath, outputPath, templateData)
}

// GenerateController 生成Controller文件
func (cg *CurdGenerator) GenerateController(ctx context.Context, config GenerateConfig) error {
	// 动态获取模板文件路径
	templatePath := cg.getTemplatePath("controller.go.template")
	fmt.Println(templatePath, "controller模板路径")

	// 动态获取输出文件路径
	outputPath := cg.getControllerOutputPath(config.PackageName)
	fmt.Println(outputPath, "controller输出路径")

	// 准备模板数据
	templateData := cg.prepareTemplateData(config)

	// 生成代码
	err := cg.generateFromTemplate(ctx, templatePath, outputPath, templateData)
	if err != nil {
		return err
	}

	// 自动更新控制器注册表
	return cg.updateControllerRegistry(config.PackageName)
}

// updateControllerRegistry 更新控制器注册表
func (cg *CurdGenerator) updateControllerRegistry(packageName string) error {
	registryPath := filepath.Join(cg.getProjectRoot(), "app", "admin", "internal", "router", "registry.go")

	content := gfile.GetContents(registryPath)
	if content == "" {
		return fmt.Errorf("无法读取 registry.go 文件")
	}

	// 检查是否已经存在该控制器的导入
	importLine := fmt.Sprintf(`"server/app/admin/internal/controller/%s"`, packageName)
	if !strings.Contains(content, importLine) {
		// 添加导入
		content = cg.addRegistryImport(content, importLine)
	}

	// 检查是否已经存在该控制器的注册项
	registryLine := fmt.Sprintf(`	"%s":`, packageName)
	if !strings.Contains(content, registryLine) {
		// 添加注册项
		newRegistryLine := fmt.Sprintf(`	"%s": func() interface{} { return %s.NewV1() },`, packageName, packageName)
		content = cg.addRegistryEntry(content, newRegistryLine)
	}

	return gfile.PutContents(registryPath, content)
}

// addRegistryImport 添加导入语句到注册表文件
func (cg *CurdGenerator) addRegistryImport(content, importLine string) string {
	// 找到最后一个controller导入的位置
	lastImportPattern := `"server/app/admin/internal/controller/`
	lastImportIndex := strings.LastIndex(content, lastImportPattern)

	if lastImportIndex != -1 {
		// 找到该行的结束位置
		lineEnd := strings.Index(content[lastImportIndex:], "\n")
		if lineEnd != -1 {
			insertPos := lastImportIndex + lineEnd
			// 在最后一个导入后添加新的导入
			content = content[:insertPos] + "\n\t" + importLine + content[insertPos:]
		}
	}

	return content
}

// addRegistryEntry 添加注册项到注册表
func (cg *CurdGenerator) addRegistryEntry(content, registryLine string) string {
	// 找到ControllerRegistry的结束位置（最后一个逗号后）
	registryStart := "var ControllerRegistry = map[string]ControllerFactory{"
	registryStartIndex := strings.Index(content, registryStart)

	if registryStartIndex != -1 {
		// 找到最后一个注册项的位置
		lastCommaIndex := strings.LastIndex(content, "},")
		if lastCommaIndex != -1 && lastCommaIndex > registryStartIndex {
			// 在最后一个注册项后添加新的注册项
			content = content[:lastCommaIndex+2] + "\n" + registryLine + content[lastCommaIndex+2:]
		}
	}

	return content
}

// getLogicOutputPath 获取Logic文件输出路径
func (cg *CurdGenerator) getLogicOutputPath(packageName string) string {
	projectRoot := cg.getProjectRoot()
	return filepath.Join(projectRoot, "app", "admin", "internal", "logic", packageName, packageName+".go")
}

// getControllerOutputPath 获取Controller文件输出路径
func (cg *CurdGenerator) getControllerOutputPath(packageName string) string {
	projectRoot := cg.getProjectRoot()
	// 修改文件名格式为 packageName_v1_packageName.go
	fileName := packageName + "_v1_" + packageName + ".go"
	return filepath.Join(projectRoot, "app", "admin", "internal", "controller", packageName, fileName)
}
