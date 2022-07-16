package logger

import (
	"github.com/Nexters/myply/infrastructure/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(config *configs.Config) *zap.SugaredLogger {
	var (
		zapConfig zap.Config
		logger    *zap.Logger
		err       error
	)

	switch config.Phase {
	case configs.Production:
		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig = getProdEncoderConfig()
		logger, err = zapConfig.Build(zap.AddCallerSkip(1))
	default:
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = zapConfig.Build()
	}

	if err != nil {
		panic(err)
	}

	appLogger := logger.Sugar()
	defer appLogger.Sync()

	return appLogger
}

func getProdEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return encoderConfig
}
