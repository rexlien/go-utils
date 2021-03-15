package logger


import (
	"context"
	"go.uber.org/zap"
)
import "go.uber.org/zap/zapcore"

type LogContext struct {
	logger *zap.Logger
	sugar *zap.SugaredLogger
	config *zap.Config
}

func (lc *LogContext) GetSugarLogger() *zap.SugaredLogger {
	return lc.sugar
}
func (lc *LogContext) GetLogger() *zap.Logger {
	return lc.logger
}
func (lc *LogContext) GetConfig() *zap.Config {
	return lc.config
}

func CreateLogContext(fields... zapcore.Field) *LogContext {
	logger, sugar, c := WithFields(fields...)
	context.WithValue(context.Background(), "_logger", sugar)
	return &LogContext{logger: logger, sugar: sugar, config: c}

}

func WithLogContext(logContext *LogContext, fields... zapcore.Field) *LogContext {
	logger := logContext.logger//(*zap.SugaredLogger)
	newLogger := logger.With(fields...)
	return &LogContext{logger: newLogger, sugar: newLogger.Sugar(), config: logContext.config}

}

func WithFields(fields... zapcore.Field) (*zap.Logger, *zap.SugaredLogger, *zap.Config) {

	l, c := createLogger()
	return l.With(fields...), l.Sugar(), c
}


func createLogger() (*zap.Logger, *zap.Config){
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := zapConfig.Build(
		zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))


	if err != nil {
		panic("create logger failed")
	}

	return l, &zapConfig
}


