package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// initLogger returns a zap.Logger instance with the given log level.
//
// The following log levels are supported:
//
//   - debug
//   - info
//   - warn
//   - error
//   - fatal
//   - panic
//
// The logger is configured to write to stderr with the following
// format:
//
//	{
//	  "timestamp": "2006-01-02T15:04:05.000Z07:00",
//	  "level": "LEVEL",
//	  "app": "dyndns-route53",
//	  "pid": PID,
//	  "caller": "FILE:LINE",
//	  "msg": "MESSAGE",
//	  "stacktrace": "STACKTRACE"
//	}
//
// The logger is also configured to include the following initial
// fields:
//
//   - pid: the process ID of the process writing the log
func initLogger(level string) *zap.Logger {
	logLevels := map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
		"fatal": zap.FatalLevel,
		"panic": zap.PanicLevel,
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevels[level]),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return zap.Must(config.Build()).With(zap.String("app", "dyndns-route53"))
}
