package log

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	OSStderrOutputPath = "os:stderr"
	OSStdoutOutputPath = "os:stdout"

	wrapperCallerSkip = 1
)

var (
	loggerValue        atomic.Value
	sugaredLoggerValue atomic.Value

	configValue  atomic.Value
	optionsValue atomic.Value
)

func utcRFC3339NanoTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	zapcore.RFC3339NanoTimeEncoder(t.UTC(), enc)
}

func init() {
	registerOSSink()
	if err := SetGlobalConfig(DefaultConfig()); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to initialize global logging: %v\n", err)
	}
}

// LastGlobalOptions gets the last options used in a call to SetGlobalConfig
func LastGlobalOptions() []zap.Option {
	ops, _ := optionsValue.Load().([]zap.Option)
	return ops
}

func SetGlobalConfig(c zap.Config) error {
	return SetGlobalConfigWithOptions(c, LastGlobalOptions()...)
}

// SetGlobal sets the Global() logger and overrides zap globals
func SetGlobal(l *zap.Logger) {
	sl := l.Sugar()
	if ol, ok := loggerValue.Swap(l).(*zap.Logger); ok {
		if err := ol.Sync(); err != nil {
			sl.Errorf("failed to sync previous logger before global swap: %v", err)
		}
	}
	sugaredLoggerValue.Store(sl)

	// override zap globals
	_ = zap.ReplaceGlobals(l)
}

func SetGlobalConfigWithOptions(c zap.Config, ops ...zap.Option) error {
	configValue.Store(c)
	optionsValue.Store(ops)
	l, err := c.Build(ops...)
	if err != nil {
		return err
	}
	SetGlobal(l)
	return nil
}

// GlobalConfig gets a copy of the current global config
func GlobalConfig() zap.Config {
	// note: using 2 values here to prevent panic if the underlying atomic interface{} is nil
	c, _ := configValue.Load().(zap.Config)
	return c
}

func DefaultConfig() zap.Config {
	c := zap.NewProductionConfig()
	c.EncoderConfig.EncodeTime = utcRFC3339NanoTimeEncoder
	c.EncoderConfig.TimeKey = "time"
	c.Level = zap.NewAtomicLevel()
	c.OutputPaths = []string{OSStderrOutputPath}
	c.Sampling = nil
	return c
}

// Global gets the global logger (thread-safe)
func Global() *zap.Logger {
	l, ok := loggerValue.Load().(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}
	return l
}

func Debug(args ...interface{}) {
	Global().WithOptions(zap.AddCallerSkip(wrapperCallerSkip)).Sugar().Debug(args...)
}

// Debugf writes a Debug level message with fmt.Sprintf and the Global() logger, see zap.SugaredLogger.Debugf
func Debugf(format string, args ...interface{}) {
	Global().WithOptions(zap.AddCallerSkip(wrapperCallerSkip)).Sugar().Debugf(format, args...)
}

// Error writes an Error level message with fmt.Sprint and the Global() logger, see zap.SugaredLogger.Error
func Error(args ...interface{}) {
	Global().WithOptions(zap.AddCallerSkip(wrapperCallerSkip)).Sugar().Error(args...)
}

// Errorf writes an Error level message with fmt.Sprintf and the Global() logger, see zap.SugaredLogger.Errorf
func Errorf(format string, args ...interface{}) {
	Global().WithOptions(zap.AddCallerSkip(wrapperCallerSkip)).Sugar().Errorf(format, args...)
}

// Info writes an Info level message with fmt.Sprint and the Global() logger, see zap.SugaredLogger.Info
func Info(args ...interface{}) {
	Global().WithOptions(zap.AddCallerSkip(wrapperCallerSkip)).Sugar().Info(args...)
}

// Infof writes an Info level message with fmt.Sprintf and the Global() logger, see zap.SugaredLogger.Infof
func Infof(format string, args ...interface{}) {
	Global().WithOptions(zap.AddCallerSkip(wrapperCallerSkip)).Sugar().Infof(format, args...)
}
