package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()
	Logger = logger.Sugar()
	Logger.Info("Logger initialised")
}

// logger, _ := zap.NewProduction()
// defer logger.Sync() // flushes buffer, if any
// sugar := logger.Sugar()