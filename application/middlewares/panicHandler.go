package middlewares

import (
	"fmt"

	"github.com/Soreing/growler/domain/common"
	"github.com/Soreing/growler/domain/general/cntxt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PanicHandlerMiddleware struct {
	lgr *zap.Logger
}

func NewPanicHandlerMiddleware(
	lgr *zap.Logger,
) *PanicHandlerMiddleware {
	return &PanicHandlerMiddleware{
		lgr: lgr,
	}
}

func (m *PanicHandlerMiddleware) Handler() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				fmt.Println("asasdadsads", err)

				lgr := m.lgr
				var ctx cntxt.IContext
				if c, ok := gctx.Get(common.ContextKey); ok {
					if ic, ok := c.(cntxt.IContext); ok {
						ctx = ic
						lgr = ctx.Logger()
					}
				}

				lgr.Error("panic recovered", zap.Any("error", err))
				gctx.JSON(500, err)
			}
		}()

		gctx.Next()
	}
}
