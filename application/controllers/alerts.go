package controllers

import (
	"github.com/Soreing/growler/application/models"
	"github.com/Soreing/growler/domain/common"
	"github.com/Soreing/growler/domain/general/cntxt"
	"github.com/Soreing/growler/domain/usecases"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AlertsController struct {
	uscs *usecases.UseCases
}

func NewAlertsController(
	uscs *usecases.UseCases,
) *AlertsController {
	return &AlertsController{
		uscs: uscs,
	}
}

func (ctrl *AlertsController) RegisterRoutes(grp *gin.RouterGroup) {
	grp.POST("/", ctrl.publishAlerts)
}

func (h *AlertsController) publishAlerts(gctx *gin.Context) {
	var ctx cntxt.IContext
	if c, ok := gctx.Get(common.ContextKey); ok {
		if ic, ok := c.(cntxt.IContext); ok {
			ctx = ic
		}
	}
	if ctx == nil {
		// TODO: Do something clever here
		return
	}
	lgr := ctx.Logger()
	lgr.Info("processing publishAlerts request")

	var grp models.AlertGroup
	err := gctx.BindJSON(&grp)
	if err != nil {
		lgr.Error("failed to bind json", zap.Error(err))
		gctx.Error(err)
		// TODO: Create custom 400 error here
		return
	}

	for _, a := range grp.Alerts {
		err = h.uscs.PublishAlert(ctx, a.Status, a.Labels, a.Annotations)
		if err != nil {
			lgr.Error(
				"failed to publish alert",
				zap.Any("alert", a),
				zap.Error(err),
			)
			gctx.Error(err)
		}
	}
}
