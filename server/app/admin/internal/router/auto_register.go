package router

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

// AutoRegister 自动路由注册器
type AutoRegister struct {
	controllers map[string]interface{}
}

// NewAutoRegister 创建自动路由注册器
func NewAutoRegister() *AutoRegister {
	return &AutoRegister{
		controllers: make(map[string]interface{}),
	}
}

// RegisterController 注册控制器
func (ar *AutoRegister) RegisterController(name string, controller interface{}) {
	ar.controllers[name] = controller
	g.Log().Infof(context.Background(), "注册控制器: %s", name)
}

// AutoBindRoutes 自动绑定所有注册的控制器
func (ar *AutoRegister) AutoBindRoutes(group *ghttp.RouterGroup) {
	var controllers []interface{}
	for name, controller := range ar.controllers {
		g.Log().Infof(context.Background(), "绑定控制器: %s", name)
		controllers = append(controllers, controller)
	}
	if len(controllers) > 0 {
		group.Bind(controllers...)
	}
}

// GetControllerCount 获取已注册的控制器数量
func (ar *AutoRegister) GetControllerCount() int {
	return len(ar.controllers)
}

// ListControllers 列出所有已注册的控制器名称
func (ar *AutoRegister) ListControllers() []string {
	var names []string
	for name := range ar.controllers {
		names = append(names, name)
	}
	return names
}

// ScanAndRegister 扫描控制器目录并自动注册
func (ar *AutoRegister) ScanAndRegister(controllerPath string) error {
	// 扫描控制器目录
	dirs, err := gfile.ScanDir(controllerPath, "*", false)
	if err != nil {
		return fmt.Errorf("扫描控制器目录失败: %v", err)
	}

	for _, dir := range dirs {
		if gfile.IsDir(dir) {
			// 获取目录名作为包名
			packageName := filepath.Base(dir)

			// 检查是否存在 *_new.go 文件
			newFile := filepath.Join(dir, packageName+"_new.go")
			if gfile.Exists(newFile) {
				g.Log().Infof(context.Background(), "发现控制器包: %s", packageName)
				// 这里可以通过反射动态加载控制器
				// 或者通过注册表来管理
			}
		}
	}

	return nil
}
