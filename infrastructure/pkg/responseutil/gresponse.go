package responseutil

import (
	"context"
	"github.com/pkg/errors"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"gorm.io/gorm"

	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/types"
)

type rspCreateData struct {
	Id interface{} `json:"id"`
}

type rspErrorData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Cause   interface{} `json:"cause"`
}

type rspListNull struct {
	Entries  []int          `json:"entries"`
	PageData types.PageData `json:"page_data"`
}

// RspOk 返回操作成功
func RspOk(c *app.RequestContext, code int, any interface{}) {
	switch code {
	case 200:
		c.JSON(
			http.StatusOK,
			any,
		)
	case 201:
		c.JSON(
			http.StatusCreated,
			rspCreateData{
				Id: any,
			},
		)
	case 204:
		c.AbortWithStatus(
			http.StatusNoContent,
		)

	default:
		c.JSON(
			http.StatusOK,
			any,
		)
	}
}

// RspErr 返回操作失败
func RspErr(c *app.RequestContext, err error) {
	code := gerror.Code(err)
	if code == gcode.CodeNil && err != nil {
		code = gcode.CodeInternalError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = gerror.NewCode(CommNotFound, err.Error())
		} else {
			err = gerror.NewCode(CommInternalServer, err.Error())
		}
		code = gerror.Code(err)
	}
	subCode := code.(BizCode).BizDetail().SubCode
	msg := msg(code.Code())
	rspError(c, code.Code(), subCode, msg, err)
}

func rspError(c *app.RequestContext, code int, subCode int, message string, err error) {
	c.JSON(
		code,
		rspErrorData{
			Cause:   err.Error(),
			Code:    subCode,
			Message: message,
		},
	)
}

func msg(code int) (msg string) {
	lang := global.Gconfig.Server.Lang
	i18n := gi18n.New()
	ctx := gi18n.WithLanguage(context.TODO(), lang)

	switch code {
	case http.StatusBadRequest:
		msg = i18n.Translate(ctx, "{#BadRequest}")
	case http.StatusUnauthorized:
		msg = i18n.Translate(ctx, "{#Unauthorized}")
	case http.StatusForbidden:
		msg = i18n.Translate(ctx, "{#Forbidden}")
	case http.StatusNotFound:
		msg = i18n.Translate(ctx, "{#NotFound}")
	case http.StatusConflict:
		msg = i18n.Translate(ctx, "{#Conflict}")
	case http.StatusInternalServerError:
		msg = i18n.Translate(ctx, "{#InternalServerError}")
	}

	return
}
