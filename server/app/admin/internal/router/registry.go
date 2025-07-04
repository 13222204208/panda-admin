package router

import (
	"server/app/admin/internal/controller/attachment"
	"server/app/admin/internal/controller/department"
	"server/app/admin/internal/controller/dict"
	"server/app/admin/internal/controller/generate"
	"server/app/admin/internal/controller/member"
	"server/app/admin/internal/controller/menu"
	"server/app/admin/internal/controller/role"
	"server/app/admin/internal/controller/user"
)

// ControllerFactory 控制器工厂函数类型
type ControllerFactory func() interface{}

// ControllerRegistry 控制器注册表
var ControllerRegistry = map[string]ControllerFactory{
	"role":       func() interface{} { return role.NewV1() },
	"department": func() interface{} { return department.NewV1() },
	"user":       func() interface{} { return user.NewV1() },
	"menu":       func() interface{} { return menu.NewV1() },
	"generate":   func() interface{} { return generate.NewV1() },
	"dict":       func() interface{} { return dict.NewV1() },
	"attachment": func() interface{} { return attachment.NewV1() },
	"member":     func() interface{} { return member.NewV1() },
}

// GetAllControllers 获取所有控制器实例
func GetAllControllers() []interface{} {
	var controllers []interface{}
	for _, factory := range ControllerRegistry {
		controllers = append(controllers, factory())
	}
	return controllers
}

// GetController 根据名称获取控制器实例
func GetController(name string) interface{} {
	if factory, exists := ControllerRegistry[name]; exists {
		return factory()
	}
	return nil
}

// RegisterController 注册新的控制器
func RegisterController(name string, factory ControllerFactory) {
	ControllerRegistry[name] = factory
}

// GetControllerNames 获取所有控制器名称
func GetControllerNames() []string {
	var names []string
	for name := range ControllerRegistry {
		names = append(names, name)
	}
	return names
}

// HasController 检查是否存在指定名称的控制器
func HasController(name string) bool {
	_, exists := ControllerRegistry[name]
	return exists
}
