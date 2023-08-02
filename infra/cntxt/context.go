package cntxt

import (
	"context"

	"go.uber.org/zap"
)

type Context struct {
	context.Context
	lgr *zap.Logger
}

func (c *Context) Logger() *zap.Logger {
	return c.lgr
}
