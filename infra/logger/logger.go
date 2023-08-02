package logger

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	lgr, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return lgr
}
