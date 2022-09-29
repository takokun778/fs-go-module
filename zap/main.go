package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	logger.Debug("debug")

	logger.Info("info")

	logger.Warn("warn")

	logger.Error("error")
}
