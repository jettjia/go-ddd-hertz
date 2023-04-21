package router

import (
	"github.com/jettjia/go-ddd-hertz/boot"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jettjia/go-ddd-hertz/global"
	"github.com/jettjia/go-ddd-hertz/interfaces/http/middleware"
	internalRouter "github.com/jettjia/go-ddd-hertz/interfaces/http/router/internal"
	sysRouter "github.com/jettjia/go-ddd-hertz/interfaces/http/router/sys"
)

func Routers(app *boot.App) *gin.Engine {
	gin.SetMode(global.Gconfig.Server.Mode)
	engine := gin.Default()
	// 健康检查
	engine.GET("/health/ready", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "ready")
	})
	engine.GET("/health/alive", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "alive")
	})

	// 配置跨域
	engine.Use(middleware.Cors())
	// 全局recover
	engine.Use(middleware.Recover())
	// 全局错误
	engine.Use(middleware.ErrorHandler)
	// auth jwt
	engine.Use(middleware.TokenAuthorization())

	// 注册路由
	ApiGroup := engine.Group("/openapi/pc/v1")
	sysRouter.InitSysRouter(ApiGroup, app) // sys

	return engine
}

func RoutersInternal(app *boot.App) *gin.Engine {
	gin.SetMode(global.Gconfig.Server.Mode)
	engine := gin.Default()
	// 健康检查
	engine.GET("/health/ready", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "ready")
	})
	engine.GET("/health/alive", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "alive")
	})

	// 配置跨域
	engine.Use(middleware.Cors())
	// 全局recover
	engine.Use(middleware.Recover())
	// 全局错误
	engine.Use(middleware.ErrorHandler)

	// 注册路由
	ApiGroup := engine.Group("/api/pc/v1")
	internalRouter.InitInternalRouter(ApiGroup, app) // sys

	return engine
}
