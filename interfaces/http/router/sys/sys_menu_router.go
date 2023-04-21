package sys

import (
	"github.com/cloudwego/hertz/pkg/route"

	"github.com/jettjia/go-ddd-demo/boot"
	handler "github.com/jettjia/go-ddd-demo/interfaces/http/handler/sys"
)

func InitSysRouter(Router *route.RouterGroup, app *boot.App) {
	hand := handler.SysMenuHandler{
		SysMenuSrv: app.SysMenuSvc,
	}
	UserRouter := Router.Group("/sys")
	{
		UserRouter.POST("menu", hand.CreateSysMenu)       // 创建
		UserRouter.DELETE("menu/:id", hand.DeleteSysMenu) // 删除
		UserRouter.PUT("menu/:id", hand.UpdateSysMenu)    // 修改
		UserRouter.GET("menu/:id", hand.FindSysMenuById)  // 查询ByID
		UserRouter.POST("menuPage", hand.FindSysMenuPage) // 查询分页
	}
}
