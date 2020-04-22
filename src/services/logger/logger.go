// Package logger contains logging utils
package logger

import (
	"fmt"
	"os"
	"strings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/siesgstarena/epicentre/src/config"
)

// Log is a package variable, which is initialized once during NewLogger() and shared to whole application
var Log Logger

type (
	// LogConfig : Logger configuration
	LogConfig struct {
		FileName   string
		MaxSize    int
		MaxAge     int
		MaxBackUp  int
		Compress   bool
		Level      string
		OutputType string
	}

	// Logger ...
	Logger struct {
		logConfig *LogConfig
		log       *zap.SugaredLogger
	}
)

func (c LogConfig) getLogFileName() string {
	return fmt.Sprintf("./logs/%s", c.FileName)
}

// LoadLogger : Creates a new instance of logger
func LoadLogger(inputConfig config.MainConfig) error {

	// access outside logger package
	config := LogConfig {
		FileName:   config.Config.FileName,
		MaxSize:    config.Config.MaxSize,
		MaxAge:     config.Config.MaxAge,
		MaxBackUp:  config.Config.MaxBackUp,
		Compress:   config.Config.Compress,
		Level:      config.Config.Level,
		OutputType: config.Config.OutputType,
	}

	// initialize a new zap logger with lumberjack configuration
	fileLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.getLogFileName(),
		MaxAge:     config.MaxAge,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackUp,
		Compress:   config.Compress,
	})

	// define the log level
	logLevel := config.getLogLevel()

	// configuration for zap logger
	zapConfig := newZapConfig()

	// define the core based on output type
	core := config.getZapCore(fileLogger, logLevel, zapConfig)

	// define the instance of new zap logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

	Log = Logger{
		logConfig: &config,
		log:       logger,
	}

	return nil
}

func (c *LogConfig) getLogLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(c.Level)) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZapConfig() zapcore.EncoderConfig {
	zapConfig := zap.NewProductionEncoderConfig()
	zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapConfig
}

func (c *LogConfig) getZapCore(fileLogger zapcore.WriteSyncer, logLevel zapcore.Level, zapConfig zapcore.EncoderConfig) zapcore.Core {
	switch strings.ToLower(strings.TrimSpace(c.OutputType)) {
	case "json":
		return zapcore.NewCore(zapcore.NewJSONEncoder(zapConfig), zapcore.NewMultiWriteSyncer(fileLogger, zapcore.AddSync(os.Stdout)), logLevel)
	default:
		return zapcore.NewCore(zapcore.NewConsoleEncoder(zapConfig), zapcore.NewMultiWriteSyncer(fileLogger, zapcore.AddSync(os.Stdout)), logLevel)
	}
}

// Debug ...
func (l *Logger) Debug(message string, args ...interface{}) {
	l.log.Debugw(message, args...)
}

// Info ...
func (l *Logger) Info(message string, args ...interface{}) {
	l.log.Infow(message, args...)
}

// Warn ...
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log.Warnw(message, args...)
}

// Error ...
func (l *Logger) Error(message string, args ...interface{}) {
	l.log.Errorw(message, args...)
}
