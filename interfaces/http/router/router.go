package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/jettjia/go-ddd-demo/boot"
	"github.com/jettjia/go-ddd-demo/interfaces/http/middleware"
	sysRouter "github.com/jettjia/go-ddd-demo/interfaces/http/router/sys"
)

func Routers(appBoot *boot.App, engine *server.Hertz) {
	// 健康检查
	engine.GET("/health/ready", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, "ready")
	})
	engine.GET("/health/alive", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, "alive")
	})

	// 中间件
	engine.Use(middleware.Cors) // 配置跨域
	engine.Use(middleware.Recover)
	engine.Use(middleware.ErrorHandler)

	// 注册路由
	ApiGroup := engine.Group("/openapi/pc/v1")
	sysRouter.InitSysRouter(ApiGroup, appBoot) // sys
}
