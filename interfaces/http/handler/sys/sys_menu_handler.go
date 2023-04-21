package sys

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gogf/gf/v2/errors/gerror"

	dto "github.com/jettjia/go-ddd-demo/application/dto/sys"
	service "github.com/jettjia/go-ddd-demo/application/service/sys"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/validate"
)

type SysMenuHandler struct {
	SysMenuSrv *service.SysMenuService
}

func (h *SysMenuHandler) CreateSysMenu(ctx context.Context, c *app.RequestContext) {
	// 参数解析
	dtoReq := dto.CreateSysMenuReq{}
	err := c.Bind(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	res, err := h.SysMenuSrv.CreateSysMenu(ctx, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusCreated, res.Id)
}

func (h *SysMenuHandler) DeleteSysMenu(ctx context.Context, c *app.RequestContext) {
	// 参数解析
	dtoReq := dto.DelSysMenusReq{}
	err := c.Bind(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	err = h.SysMenuSrv.DeleteSysMenu(ctx, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusNoContent, nil)
}

func (h *SysMenuHandler) UpdateSysMenu(ctx context.Context, c *app.RequestContext) {
	// 参数解析
	dtoReq := dto.UpdateSysMenuReq{}
	err := c.Bind(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}
	err = c.Bind(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}
	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	err = h.SysMenuSrv.UpdateSysMenu(ctx, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusNoContent, nil)
}

func (h *SysMenuHandler) FindSysMenuById(ctx context.Context, c *app.RequestContext) {
	// 参数解析
	dtoReq := dto.FindSysMenuByIdReq{}
	err := c.Bind(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuById(ctx, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rsp)
}

func (h *SysMenuHandler) FindSysMenuByQuery(ctx context.Context, c *app.RequestContext) {
	// 参数解析
	dtoReq := dto.FindSysMenuByQueryReq{}
	err := c.Bind(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuByQuery(ctx, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusOK, rsp)
}

func (h *SysMenuHandler) FindSysMenuAll(ctx context.Context, c *app.RequestContext) {
	// 参数解析
	dtoReq := dto.FindSysMenuAllReq{}
	err := c.Bind(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuAll(ctx, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}
	responseutil.RspOk(c, http.StatusOK, rsp)
}

func (h *SysMenuHandler) FindSysMenuPage(ctx context.Context, c *app.RequestContext) {
	// 参数解析
	dtoReq := dto.FindSysMenuPageReq{}
	err := c.Bind(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 参数过滤
	err = validate.Validate(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, err.Error())
		_ = c.Error(err)
		return
	}

	// 业务处理
	rsp, err := h.SysMenuSrv.FindSysMenuPage(ctx, &dtoReq)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responseutil.RspOk(c, http.StatusOK, rsp)
}
