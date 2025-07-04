package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/consts"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/library/generate"
	"server/app/admin/internal/library/generate/curd"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

// 代码生成
// 代码生成 - 同时生成后端和前端
func (s *sGenerate) CodeGenRecord(ctx context.Context, req v1.CodeGenRecordReq) (res *v1.CodeGenRecordRes, err error) {
	res = &v1.CodeGenRecordRes{}

	// 检查记录是否存在
	count, err := dao.CodeGenRecord.Ctx(ctx).Where(dao.CodeGenRecord.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询代码生成记录失败")
	}
	if count == 0 {
		return nil, gerror.New("代码生成记录不存在")
	}

	// 在代码生成前执行 gf gen dao
	g.Log().Info(ctx, "开始执行 gf gen dao...")
	if err := s.executeGfGenDao(ctx); err != nil {
		return nil, gerror.Wrap(err, "DAO生成失败")
	}

	// 解析生成选项
	generateOptions, err := generate.ParseGenerateOptionsFromJSON(req.Options)
	if err != nil {
		return nil, gerror.Wrap(err, "解析生成选项失败")
	}
	g.Log().Info(ctx, "生成选项:", generateOptions)

	// 准备统一的生成配置
	config := curd.GenerateConfig{
		TableName:    req.TableName,
		TableComment: req.TableComment,
		EntityName:   convertToGoField(req.PackageName),
		PackageName:  req.PackageName,
		ModuleName:   req.ModuleName,
		Columns:      convertToColumns(req.Columns),
		Options:      generateOptions,
	}

	// ========== 开始同步生成后端和前端代码 ==========
	g.Log().Info(ctx, "开始同时生成后端和前端代码...")

	// 1. 生成后端代码
	if err := s.generateBackendCode(ctx, config); err != nil {
		return nil, gerror.Wrap(err, "后端代码生成失败")
	}

	// 2. 生成前端代码
	if err := s.generateFrontendCode(ctx, config); err != nil {
		// 前端生成失败不影响整体流程，但要记录错误
		g.Log().Error(ctx, "前端代码生成失败:", err)
		g.Log().Warning(ctx, "前端代码生成失败，但后端代码已成功生成")
	}

	// 3. 创建菜单（如果配置了菜单信息）
	if err := s.createMenuIfConfigured(ctx, generateOptions, config); err != nil {
		g.Log().Error(ctx, "菜单创建失败:", err)
		// 菜单创建失败不影响整体流程，但要记录错误
	}

	// 更新记录状态
	_, err = dao.CodeGenRecord.Ctx(ctx).
		Where(dao.CodeGenRecord.Columns().Id, req.Id).
		Data(g.Map{
			dao.CodeGenRecord.Columns().TableName:    req.TableName,
			dao.CodeGenRecord.Columns().TableComment: req.TableComment,
			dao.CodeGenRecord.Columns().PackageName:  req.PackageName,
			dao.CodeGenRecord.Columns().ModuleName:   req.ModuleName,
			dao.CodeGenRecord.Columns().Options:      req.Options,
			dao.CodeGenRecord.Columns().Columns:      req.Columns,
			dao.CodeGenRecord.Columns().Status:       1,
		}).
		Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新代码生成记录失败")
	}

	g.Log().Info(ctx, "代码生成完成 - 后端和前端文件已同时生成")
	return
}

// generateBackendCode 生成后端代码
func (s *sGenerate) generateBackendCode(ctx context.Context, config curd.GenerateConfig) error {
	g.Log().Info(ctx, "开始生成后端代码...")

	generator := curd.NewCurdGenerator()

	// 生成API文件
	if err := generator.GenerateAPI(ctx, config); err != nil {
		return gerror.Wrap(err, "API代码生成失败")
	}

	// 生成Logic文件
	if err := generator.GenerateLogic(ctx, config); err != nil {
		return gerror.Wrap(err, "Logic代码生成失败")
	}

	// 执行 gf gen ctrl
	if err := s.executeGfGenCtrl(ctx); err != nil {
		return gerror.Wrap(err, "Controller生成失败")
	}

	// 生成Controller文件
	if err := generator.GenerateController(ctx, config); err != nil {
		return gerror.Wrap(err, "Controller代码生成失败")
	}

	// 执行 gf gen service
	if err := s.executeGfGenService(ctx); err != nil {
		return gerror.Wrap(err, "Service生成失败")
	}

	g.Log().Info(ctx, "后端代码生成完成")
	return nil
}

// generateFrontendCode 生成前端代码
func (s *sGenerate) generateFrontendCode(ctx context.Context, config curd.GenerateConfig) error {
	g.Log().Info(ctx, "开始生成前端代码...")

	// 创建前端生成器
	frontendGenerator := curd.NewFrontendGenerator()

	// 准备前端配置
	frontendConfig := curd.FrontendConfig{
		GenerateConfig: config,
		WebRoot:        "", // 使用默认前端项目路径
	}

	// 验证前端配置
	if err := frontendGenerator.ValidateConfig(frontendConfig); err != nil {
		return gerror.Wrap(err, "前端配置验证失败")
	}

	// 生成所有前端文件
	if err := frontendGenerator.GenerateAll(ctx, frontendConfig); err != nil {
		return gerror.Wrap(err, "前端文件生成失败")
	}

	g.Log().Info(ctx, "前端代码生成完成")
	return nil
}

// 批量生成多个模块的后端和前端代码
func (s *sGenerate) BatchGenerateFullStack(ctx context.Context, requests []v1.CodeGenRecordReq) error {
	g.Log().Infof(ctx, "开始批量生成 %d 个模块的完整代码...", len(requests))

	for i, req := range requests {
		g.Log().Infof(ctx, "正在生成模块 %d/%d: %s", i+1, len(requests), req.ModuleName)

		// 解析生成选项
		generateOptions, err := generate.ParseGenerateOptionsFromJSON(req.Options)
		if err != nil {
			return gerror.Wrapf(err, "模块 %s 解析生成选项失败", req.ModuleName)
		}

		// 准备配置
		config := curd.GenerateConfig{
			TableName:    req.TableName,
			TableComment: req.TableComment,
			EntityName:   convertToGoField(req.PackageName),
			PackageName:  req.PackageName,
			ModuleName:   req.ModuleName,
			Columns:      convertToColumns(req.Columns),
			Options:      generateOptions,
		}

		// 生成后端代码
		if err := s.generateBackendCode(ctx, config); err != nil {
			return gerror.Wrapf(err, "模块 %s 后端代码生成失败", req.ModuleName)
		}

		// 生成前端代码
		if err := s.generateFrontendCode(ctx, config); err != nil {
			g.Log().Errorf(ctx, "模块 %s 前端代码生成失败: %v", req.ModuleName, err)
			// 继续处理下一个模块
		}

		g.Log().Infof(ctx, "模块 %s 代码生成完成", req.ModuleName)
	}

	g.Log().Info(ctx, "批量代码生成完成")
	return nil
}

// convertToColumns 将API字段转换为生成器字段
func convertToColumns(apiColumns interface{}) []curd.Column {
	var columns []curd.Column

	// 处理字符串类型的JSON数据
	if jsonStr, ok := apiColumns.(string); ok {
		var columnInfos []consts.ColumnConfig
		if err := json.Unmarshal([]byte(jsonStr), &columnInfos); err != nil {
			return columns
		}
		// 转换ColumnInfo到curd.Column
		for _, colInfo := range columnInfos {
			column := curd.Column{
				ColumnName:     colInfo.FieldName,
				ColumnComment:  colInfo.FieldComment,
				GoField:        convertToGoField(colInfo.FieldName),   // 转换为Go字段名
				GoType:         convertToGoType(colInfo.HtmlType),     // 转换为Go类型
				JsonField:      convertToJsonField(colInfo.FieldName), // 转换为JSON字段名
				IsRequired:     colInfo.IsRequired,                    // 根据是否可空判断必填
				IsQuery:        colInfo.IsQuery,                       // 判断是否用于查询
				IsList:         colInfo.IsList,                        // 判断是否用于列表显示
				IsPointer:      true,                                  // 可空字段使用指针
				IsUnique:       colInfo.IsUnique,                      // 是否唯一
				ValidationRule: colInfo.ValidationRules,               // 生成验证规则
			}
			columns = append(columns, column)
		}
		return columns
	}

	// 保持原有的interface{}处理逻辑作为兼容
	if cols, ok := apiColumns.([]interface{}); ok {
		for _, col := range cols {
			if colMap, ok := col.(map[string]interface{}); ok {
				column := curd.Column{
					ColumnName:     getString(colMap, "columnName"),
					ColumnComment:  getString(colMap, "columnComment"),
					GoField:        getString(colMap, "goField"),
					GoType:         getString(colMap, "goType"),
					JsonField:      getString(colMap, "jsonField"),
					IsRequired:     getBool(colMap, "isRequired"),
					IsQuery:        getBool(colMap, "isQuery"),
					IsPointer:      getBool(colMap, "isPointer"),
					IsUnique:       getBool(colMap, "isUnique"),
					ValidationRule: getString(colMap, "validationRule"),
				}
				columns = append(columns, column)
			}
		}
	}

	return columns
}

// 辅助转换函数
func convertToGoField(columnName string) string {
	// 将数据库字段名转换为Go字段名（驼峰命名）
	return gstr.CaseCamel(columnName)
}

func convertToJsonField(columnName string) string {
	// 将数据库字段名转换为JSON字段名（蛇形命名）
	return gstr.CaseSnake(columnName)
}

func convertToGoType(dataType string) string {
	// 根据数据库类型转换为Go类型
	switch strings.ToLower(dataType) {
	case "int", "tinyint", "smallint", "mediumint":
		return "int"
	case "bigint":
		return "int64"
	case "varchar", "char", "text", "longtext", "mediumtext":
		return "string"
	case "decimal", "float", "double":
		return "float64"
	case "json":
		return "string"
	default:
		return "string"
	}
}

// 辅助函数
func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

// executeGfGenDao 执行 gf gen dao 命令
func (s *sGenerate) executeGfGenDao(ctx context.Context) error {
	// 获取 admin 模块的根目录
	adminDir := filepath.Join(".") // 当前已经在 admin 目录下运行

	// 创建命令
	cmd := exec.Command("gf", "gen", "dao")
	cmd.Dir = adminDir // 设置为 admin 模块目录

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		g.Log().Error(ctx, "gf gen dao 执行失败", string(output), err)
		return gerror.Wrapf(err, "gf gen dao 执行失败: %s", string(output))
	}

	g.Log().Info(ctx, "gf gen dao 执行成功", string(output))
	return nil
}

// executeGfGenService 执行 gf gen service 命令
func (s *sGenerate) executeGfGenService(ctx context.Context) error {
	// 获取 admin 模块的根目录
	adminDir := filepath.Join(".") // 当前已经在 admin 目录下运行

	// 创建命令
	cmd := exec.Command("gf", "gen", "service")
	cmd.Dir = adminDir // 设置为 admin 模块目录

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		g.Log().Error(ctx, "gf gen service 执行失败", string(output), err)
		return gerror.Wrapf(err, "gf gen service 执行失败: %s", string(output))
	}

	g.Log().Info(ctx, "gf gen service 执行成功", string(output))
	return nil
}

// executeGfGenCtrl 执行 gf gen ctrl -m yes 命令
func (s *sGenerate) executeGfGenCtrl(ctx context.Context) error {
	// 获取 admin 模块的根目录
	adminDir := filepath.Join(".") // 当前已经在 admin 目录下运行

	// 创建命令
	cmd := exec.Command("gf", "gen", "ctrl", "--m=yes")
	cmd.Dir = adminDir // 设置为 admin 模块目录

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		g.Log().Error(ctx, "gf gen ctrl 执行失败", string(output), err)
		return gerror.Wrapf(err, "gf gen ctrl 执行失败: %s", string(output))
	}

	g.Log().Info(ctx, cmd.String()+" 执行成功", string(output))
	return nil
}

// createMenuIfConfigured 如果配置了菜单信息则创建菜单
func (s *sGenerate) createMenuIfConfigured(ctx context.Context, options *consts.GenerateOptions, config curd.GenerateConfig) error {
	// 检查是否配置了菜单信息
	if options.MenuName == "" {
		g.Log().Info(ctx, "未配置菜单名称，跳过菜单创建")
		return nil
	}

	// 构建菜单路径和组件路径
	menuPath := fmt.Sprintf("/%s", config.ModuleName)
	componentPath := fmt.Sprintf("%s/index", config.ModuleName)
	routeName := gstr.CaseCamel(config.ModuleName)

	// 准备菜单数据
	menuData := g.Map{
		"menuType":  0, // 菜单类型：0菜单
		"parentId":  options.ParentMenuId,
		"title":     options.MenuName,
		"name":      routeName,
		"path":      menuPath,
		"component": componentPath,
		"rank":      99, // 默认排序
		"icon":      options.MenuIcon,
		"showLink":  1, // 显示在菜单中
		"keepAlive": 0, // 缓存页面
	}

	// 检查菜单名称是否已存在
	count, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Name, routeName).Count()
	if err != nil {
		return gerror.Wrap(err, "检查菜单名称失败")
	}
	if count > 0 {
		g.Log().Warningf(ctx, "菜单路由名称 %s 已存在，跳过菜单创建", routeName)
		return nil
	}

	// 检查父级菜单是否存在（如果指定了父级菜单）
	if options.ParentMenuId > 0 {
		parentCount, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, options.ParentMenuId).Count()
		if err != nil {
			return gerror.Wrap(err, "检查父级菜单失败")
		}
		if parentCount == 0 {
			g.Log().Warningf(ctx, "父级菜单ID %d 不存在，将创建为顶级菜单", options.ParentMenuId)
			menuData["parentId"] = 0
		}
	}

	// 插入菜单数据
	id, err := dao.Menu.Ctx(ctx).Data(menuData).InsertAndGetId()
	if err != nil {
		return gerror.Wrap(err, "创建菜单失败")
	}

	g.Log().Infof(ctx, "菜单创建成功，ID: %d, 名称: %s, 路径: %s", id, options.MenuName, menuPath)
	return nil
}
