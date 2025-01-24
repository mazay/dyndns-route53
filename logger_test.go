package main

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestInitLogger_LogLevels(t *testing.T) {
	logLevels := []string{"debug", "info", "warn", "error", "fatal", "panic"}
	for _, level := range logLevels {
		logger := initLogger(level)
		if logger.Core().Enabled(zapcore.DebugLevel) != (level == "debug") {
			t.Errorf("expected log level %s to be enabled, but got %v", level, logger.Core().Enabled(zapcore.DebugLevel))
		}
	}
}
