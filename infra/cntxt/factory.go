package cntxt

import (
	icntxt "github.com/Soreing/growler/domain/general/cntxt"
	"go.uber.org/zap"
)

type Factory struct {
	lgr *zap.Logger
}

func NewFactory(
	lgr *zap.Logger,
) *Factory {
	return &Factory{
		lgr: lgr,
	}
}

func (f *Factory) Create(
	traceId string,
) icntxt.IContext {
	return &Context{
		lgr: f.lgr.With(
			zap.String("traceId", traceId),
		),
	}
}
