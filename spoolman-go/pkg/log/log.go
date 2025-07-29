package log

import (
	"fmt"
	"os"
	"path/filepath"
	"spoolman-go/pkg/config"
	"strings"
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	baseLogger *zap.Logger
	reloadMu   sync.Mutex
	curCfg     *config.LogConfig
)

// Setup initializes logger with config.LogConfig.
// If called again with different config, it reconfigures logger (safe for live reload).
func Setup(cfg *config.LogConfig) error {
	reloadMu.Lock()
	defer reloadMu.Unlock()
	if cfg == curCfg && baseLogger != nil {
		return nil
	}
	level := zapcore.InfoLevel
	if &level == nil {
		return fmt.Errorf("unreachable error %s", "zapcore.InfoLevel is nil")
	}
	_ = level.Set(strings.ToLower(cfg.Level))
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	_ = os.MkdirAll(filepath.Dir(cfg.LogFile), 0755)
	rotator := &lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    cfg.MaxSizeMB,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAgeDays,
		Compress:   false,
	}

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(rotator),
		level,
	)
	stdoutCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		level,
	)
	tee := zapcore.NewTee(fileCore, stdoutCore)
	logger := zap.New(tee, zap.AddCaller())

	if baseLogger != nil {
		_ = baseLogger.Sync()
	}
	baseLogger = logger
	curCfg = cfg
	return nil
}

// Clone returns a named logger.
func Clone(name string) *zap.Logger {
	reloadMu.Lock()
	defer reloadMu.Unlock()
	if baseLogger == nil {
		_ = Setup(&config.Get().Log)
	}
	return baseLogger.Named(name)
}

// With returns a logger with additional context fields.
func With(fields ...zap.Field) *zap.Logger {
	reloadMu.Lock()
	defer reloadMu.Unlock()
	if baseLogger == nil {
		_ = Setup(&config.Get().Log)
	}
	return baseLogger.With(fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Clone("default").Debug(msg, fields...)
}
func Info(msg string, fields ...zap.Field) {
	Clone("default").Info(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
	Clone("default").Warn(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	Clone("default").Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Clone("default").Fatal(msg, fields...)
}
func Sync() {
	reloadMu.Lock()
	defer reloadMu.Unlock()
	if baseLogger != nil {
		_ = baseLogger.Sync()
	}
}
