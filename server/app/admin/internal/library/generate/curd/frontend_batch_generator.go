package curd

import (
	"context"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

// FrontendBatchGenerator 前端批量生成器
type FrontendBatchGenerator struct {
	frontendGenerator *FrontendGenerator
	mu                sync.Mutex
}

// NewFrontendBatchGenerator 创建新的前端批量生成器
func NewFrontendBatchGenerator() *FrontendBatchGenerator {
	return &FrontendBatchGenerator{
		frontendGenerator: NewFrontendGenerator(),
	}
}

// BatchGenerateConfig 批量生成配置
type BatchGenerateConfig struct {
	Configs   []FrontendConfig `json:"configs"`   // 多个模块配置
	Overwrite bool             `json:"overwrite"` // 是否覆盖已存在的文件
	Parallel  bool             `json:"parallel"`  // 是否并行生成
}

// GenerateMultipleModules 批量生成多个模块的前端文件
func (fbg *FrontendBatchGenerator) GenerateMultipleModules(ctx context.Context, batchConfig BatchGenerateConfig) error {
	if len(batchConfig.Configs) == 0 {
		return fmt.Errorf("没有提供生成配置")
	}

	if batchConfig.Parallel {
		return fbg.generateParallel(ctx, batchConfig)
	}
	return fbg.generateSequential(ctx, batchConfig)
}

// generateSequential 顺序生成
func (fbg *FrontendBatchGenerator) generateSequential(ctx context.Context, batchConfig BatchGenerateConfig) error {
	for i, config := range batchConfig.Configs {
		g.Log().Infof(ctx, "开始生成模块 %d/%d: %s", i+1, len(batchConfig.Configs), config.ModuleName)

		if err := fbg.validateAndGenerate(ctx, config, batchConfig.Overwrite); err != nil {
			return fmt.Errorf("生成模块 %s 失败: %v", config.ModuleName, err)
		}

		g.Log().Infof(ctx, "模块 %s 生成完成", config.ModuleName)
	}
	return nil
}

// generateParallel 并行生成
func (fbg *FrontendBatchGenerator) generateParallel(ctx context.Context, batchConfig BatchGenerateConfig) error {
	var wg sync.WaitGroup
	errorChan := make(chan error, len(batchConfig.Configs))

	for _, config := range batchConfig.Configs {
		wg.Add(1)
		go func(cfg FrontendConfig) {
			defer wg.Done()
			if err := fbg.validateAndGenerate(ctx, cfg, batchConfig.Overwrite); err != nil {
				errorChan <- fmt.Errorf("生成模块 %s 失败: %v", cfg.ModuleName, err)
			}
		}(config)
	}

	wg.Wait()
	close(errorChan)

	// 检查是否有错误
	for err := range errorChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// validateAndGenerate 验证配置并生成文件
func (fbg *FrontendBatchGenerator) validateAndGenerate(ctx context.Context, config FrontendConfig, overwrite bool) error {
	fbg.mu.Lock()
	defer fbg.mu.Unlock()

	// 验证配置
	if err := fbg.frontendGenerator.ValidateConfig(config); err != nil {
		return err
	}

	// 检查文件是否已存在
	if !overwrite {
		if err := fbg.checkExistingFiles(config); err != nil {
			return err
		}
	}

	// 生成文件
	return fbg.frontendGenerator.GenerateAll(ctx, config)
}

// checkExistingFiles 检查文件是否已存在
func (fbg *FrontendBatchGenerator) checkExistingFiles(config FrontendConfig) error {
	files := fbg.frontendGenerator.GetGeneratedFiles(config)
	for _, file := range files {
		if gfile.Exists(file) {
			return fmt.Errorf("文件已存在: %s，请设置 overwrite=true 来覆盖", file)
		}
	}
	return nil
}

// GenerateByFileTypes 根据文件类型批量生成
func (fbg *FrontendBatchGenerator) GenerateByFileTypes(ctx context.Context, config FrontendConfig, fileTypes []string) error {
	if err := fbg.frontendGenerator.ValidateConfig(config); err != nil {
		return err
	}

	for _, fileType := range fileTypes {
		if err := fbg.frontendGenerator.GenerateSingleFile(ctx, config, fileType); err != nil {
			return fmt.Errorf("生成 %s 文件失败: %v", fileType, err)
		}
		g.Log().Infof(ctx, "文件类型 %s 生成完成", fileType)
	}

	return nil
}

// GetBatchGenerationSummary 获取批量生成摘要
func (fbg *FrontendBatchGenerator) GetBatchGenerationSummary(configs []FrontendConfig) map[string]interface{} {
	summary := map[string]interface{}{
		"total_modules": len(configs),
		"modules":       make([]map[string]interface{}, 0),
		"total_files":   0,
	}

	totalFiles := 0
	modules := make([]map[string]interface{}, 0)

	for _, config := range configs {
		files := fbg.frontendGenerator.GetGeneratedFiles(config)
		totalFiles += len(files)

		moduleInfo := map[string]interface{}{
			"module_name":  config.ModuleName,
			"package_name": config.PackageName,
			"entity_name":  config.EntityName,
			"files_count":  len(files),
			"files":        files,
		}
		modules = append(modules, moduleInfo)
	}

	summary["modules"] = modules
	summary["total_files"] = totalFiles

	return summary
}
