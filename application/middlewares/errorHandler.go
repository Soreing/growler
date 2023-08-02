package middlewares

import (
	"github.com/Soreing/growler/domain/common"
	"github.com/Soreing/growler/domain/general/cntxt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ErrorHandlerMiddleware struct {
	lgr *zap.Logger
}

func NewErrorHandlerMiddleware(
	lgr *zap.Logger,
) *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{
		lgr: lgr,
	}
}

func (m *ErrorHandlerMiddleware) Handler() gin.HandlerFunc {
	return func(gctx *gin.Context) {

		lgr := m.lgr
		gctx.Next()

		var ctx cntxt.IContext
		if c, ok := gctx.Get(common.ContextKey); ok {
			if ic, ok := c.(cntxt.IContext); ok {
				ctx = ic
				lgr = ctx.Logger()
			}
		}

		if len(gctx.Errors) > 0 {
			errs := make([]error, len(gctx.Errors))
			for i, e := range gctx.Errors {
				errs[i] = e
			}

			lgr.Error(
				"errors processing request",
				zap.Errors("error", errs),
			)

			gctx.Status(500)
			gctx.JSON(500, errs[0])
		}
	}
}
