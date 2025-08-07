package logger

import (
	"my-pet-simple-messenger/api/internal/config"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(newCfg *config.Config) *Logger {

	lvl := newCfg.Loglevel

	var cfg zap.Config
	if lvl == "debug" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	switch lvl {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	logg, err := cfg.Build()
	if err != nil {
		panic("не удалось инициализировать логгер: " + err.Error())
	}

	return &Logger{
		SugaredLogger: logg.Sugar(),
	}
}
