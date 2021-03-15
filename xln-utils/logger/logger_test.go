package logger_test

import (
	"github.com/rexlien/go-utils/xln-utils/logger"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLogLevel(t *testing.T) {

	ctx := logger.CreateLogContext()
	ctx.GetConfig().Level.SetLevel(zapcore.ErrorLevel)
	l := ctx.GetSugarLogger()

	l.Debugf("Test Debug Level")

	ctx.GetConfig().Level.SetLevel(zapcore.DebugLevel)
	l.Debugf("Test Debug Level")
}


