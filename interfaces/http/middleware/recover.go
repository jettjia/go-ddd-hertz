package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
)

func Recover(ctx context.Context, c *app.RequestContext) {
	if err := recover(); err != nil {
		errorInfo := responseutil.Panic(err)
		responseutil.RspErr(c, gerror.NewCodef(responseutil.CommInternalServer, "%+v", errorInfo.Internal)) // 前端返回
		return
	}

	c.Next(ctx)
}
