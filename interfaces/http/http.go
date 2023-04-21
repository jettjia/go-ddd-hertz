package http

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/requestid"

	"github.com/jettjia/go-ddd-demo/boot"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/interfaces/http/router"
)

func InitHttp(app *boot.App) {
	// open api
	go func() {
		hertz := server.New(
			server.WithStreamBody(true),
			server.WithHostPorts(fmt.Sprintf("0.0.0.0:%d", global.Gconfig.Server.PublicPort)),
		)

		// 注册路由
		router.Routers(app, hertz)

		hertz.Use(recovery.Recovery())
		hertz.Use(requestid.New())

		hertz.Spin()
	}()
}
