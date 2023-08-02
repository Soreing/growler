package application

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Soreing/growler/application/controllers"
	"github.com/Soreing/growler/application/middlewares"
	"github.com/Soreing/growler/domain"
	"github.com/Soreing/growler/infra"
	"github.com/Soreing/growler/infra/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var dependencySet = wire.NewSet(
	domain.DependencySet,
	infra.DependencySet,
	controllers.NewAlertsController,
	middlewares.NewPanicHandlerMiddleware,
	middlewares.NewErrorHandlerMiddleware,
	middlewares.NewAddContextMiddleware,
	NewApp,
)

func Start() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	app.start()
}

type app struct {
	cfg    *config.AppConfigs
	alrts  *controllers.AlertsController
	mdwPnc *middlewares.PanicHandlerMiddleware
	mdwErr *middlewares.ErrorHandlerMiddleware
	mdwCtx *middlewares.AddContextMiddleware
	infra  *infra.Infrastructure
}

func NewApp(
	cfg *config.AppConfigs,
	alrts *controllers.AlertsController,
	mdwPnc *middlewares.PanicHandlerMiddleware,
	mdwErr *middlewares.ErrorHandlerMiddleware,
	mdwCtx *middlewares.AddContextMiddleware,
	infra *infra.Infrastructure,
) *app {
	return &app{
		cfg:    cfg,
		alrts:  alrts,
		mdwPnc: mdwPnc,
		mdwErr: mdwErr,
		mdwCtx: mdwCtx,
		infra:  infra,
	}
}

func (a *app) start() {

	// - Empty Context
	//ctx := context.TODO()
	lgr, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// - Setting up gin router
	var router *gin.Engine
	if a.cfg.Development {
		router = gin.Default()
	} else {
		router = gin.New()
		gin.SetMode(gin.ReleaseMode)
		router.SetTrustedProxies(nil)
	}

	// - Setting up middlewares
	router.Use(a.mdwPnc.Handler())
	router.Use(a.mdwErr.Handler())
	router.Use(a.mdwCtx.Handler())

	// - Setting up routes
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "alive",
		})
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, "Not Found")
	})
	a.alrts.RegisterRoutes(router.Group("alerts"))

	a.infra.Start()
	go func() {

		err := router.Run(":" + a.cfg.PortNumber)
		if err != nil && err != http.ErrServerClosed {
			lgr.Error("failed running router", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.infra.Stop()
	lgr.Info("Server exiting")
}
