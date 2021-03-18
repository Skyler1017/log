package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func InitLogger() *zap.Logger {
	//logger, _ = zap.NewProduction()
	writeSyncer := zapcore.AddSync(os.Stdout)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core)
	return Logger
}

func Sync() error{
	err := Logger.Sync()
	return err
}

func GetLogger() *zap.Logger {
	if Logger != nil {
		return Logger
	}
	return InitLogger()
}