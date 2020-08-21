package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	//普通的日志文件大小为1G，存储1个
	//hubble日志的文件大小为2g,存储10个
	commonmaxSize    = 1024
	commonmaxBackups = 1
	hubbleSize       = 2048
	hubbleBackups    = 10
)

func NewCommonlogger(filename string) *zap.Logger {
	return NewLogger(filename, NewCommonEncoderConfig(), commonmaxSize, commonmaxBackups)
}
func NewLoggerWithonlyMes(filename string) *zap.Logger {
	return NewLogger(filename, NewHubbleEncoderConfig(), hubbleSize, hubbleBackups)
}

func NewLogger(fileName string, encodeConfig zapcore.EncoderConfig, maxsize int, maxbackup int) *zap.Logger {

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	infoLog := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath + fileName + ".info",
		MaxSize:    maxsize, // megabytes
		MaxBackups: maxbackup,
		LocalTime:  true,
	})
	errLog := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath + fileName + ".error",
		MaxSize:    maxsize, // megabytes
		MaxBackups: maxbackup,
		LocalTime:  true,
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encodeConfig),
		zapcore.NewMultiWriteSyncer(
			infoLog),
		lowPriority,
	)
	error := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encodeConfig),
		zapcore.NewMultiWriteSyncer(
			errLog),
		highPriority,
	)
	tee := zapcore.NewTee(core, error)
	defaultLogger := zap.New(tee)
	defaultLogger.Sugar()
	return defaultLogger

}

func NewCommonEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
func NewHubbleEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "",
		LevelKey:       "",
		NameKey:        "",
		CallerKey:      "",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
