package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	var err error

	// changing time from epoch to iso
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	logger, err = config.Build(zap.AddCallerSkip(1)) // applying the custom configs

	//logger, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func InfoLog(message string) {
	logger.Info(message)
}

func ErrorLog(message string) {
	logger.Error(message)
}

func DebugLog(message string) {
	logger.Debug(message)
}



