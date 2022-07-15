package logger

import (
	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Set = wire.NewSet(NewLogger)

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
	defer logger.Sync()

	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}

func getProdEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return encoderConfig
}
