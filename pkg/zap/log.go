package zap

import "go.uber.org/zap"

func ErrorLog(err error) {
	Logger.Error("SYSTEM ERROR", zap.Error(err))
}

func InfoLog(msg string) {
	Logger.Info("SYSTEM INFO", zap.String("msg: ", msg))
}
