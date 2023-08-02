package cntxt

import (
	"context"

	"go.uber.org/zap"
)

type IFactory interface {
	Create(traceId string) IContext
}

type IContext interface {
	context.Context
	Logger() *zap.Logger
}
