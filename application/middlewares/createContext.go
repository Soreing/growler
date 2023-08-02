package middlewares

import (
	"github.com/Soreing/growler/domain/common"
	"github.com/Soreing/growler/domain/usecases"
	"github.com/gin-gonic/gin"
)

type AddContextMiddleware struct {
	uscs *usecases.UseCases
}

func NewAddContextMiddleware(
	uscs *usecases.UseCases,
) *AddContextMiddleware {
	return &AddContextMiddleware{
		uscs: uscs,
	}
}

func (m *AddContextMiddleware) Handler() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		gctx.Set(common.ContextKey, m.uscs.CreateContext())
		gctx.Next()
	}
}
