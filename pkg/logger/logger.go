package logger

import (
	"os"

	"drako/pkg/build"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Start()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type Logger struct {
	sugarLogger *zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{}
}

var LoggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *Logger) Start() {
	logFile, err := os.OpenFile("./logs/log.log", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	writer := zapcore.AddSync(logFile)

	fileEncoderCfg := zap.NewProductionEncoderConfig()
	fileEncoderCfg.LevelKey = "LEVEL"
	fileEncoderCfg.CallerKey = "CALLER"
	fileEncoderCfg.TimeKey = "TIME"
	fileEncoderCfg.NameKey = "NAME"
	fileEncoderCfg.MessageKey = "MESSAGE"
	fileEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderCfg)

	consoleEncoderCfg := zap.NewDevelopmentEncoderConfig()
	consoleEncoderCfg.LevelKey = "LEVEL"
	consoleEncoderCfg.CallerKey = "CALLER"
	consoleEncoderCfg.TimeKey = "TIME"
	consoleEncoderCfg.NameKey = "NAME"
	consoleEncoderCfg.MessageKey = "MESSAGE"
	consoleEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderCfg)

	fileCore := zapcore.NewCore(fileEncoder, writer, zap.NewAtomicLevelAt(LoggerLevelMap["debug"]))
	var cores []zapcore.Core
	cores = append(cores, fileCore)

	if build.DRAKO_DEBUG {
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(LoggerLevelMap["debug"]))
		cores = append(cores, consoleCore)
	}

	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error("Error syncing logger:", err)
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *Logger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
