package logging

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger(logLevel string) *zap.Logger {
	var err error
	var logger *zap.Logger
	if logLevel == "dev" {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Println("Logger initialization failed...")
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Println("Logger initialization failed...")
		}
	}

	defer logger.Sync()
	logger.Info("honeypot-ingestion is initializing...", zap.String("logLevel", logLevel))

	return logger
}
