package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"server/app/admin/internal/middleware"
	"server/app/admin/internal/router"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			//设置静态资源访问目录
			s.AddStaticPath("/uploads", "./uploads")

			s.Group("/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.Auth) // 添加认证中间件

				// 自动绑定所有注册的控制器
				controllers := router.GetAllControllers()
				g.Log().Infof(ctx, "正在绑定 %d 个控制器", len(controllers))
				group.Bind(controllers...)

				// 打印已注册的控制器列表
				controllerNames := router.GetControllerNames()
				g.Log().Infof(ctx, "已注册的控制器: %v", controllerNames)
			})
			s.Run()
			return nil
		},
	}
)
