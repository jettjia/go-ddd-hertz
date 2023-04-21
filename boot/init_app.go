package boot

import (
	sysSvc "github.com/jettjia/go-ddd-hertz/application/service/sys"
)

type App struct {
	// sys 后台系统
	SysMenuSvc *sysSvc.SysMenuService
}

func InitApp() *App {
	return &App{
		// sys 后台系统
		SysMenuSvc: sysSvc.NewSysMenuService(),
	}
}
