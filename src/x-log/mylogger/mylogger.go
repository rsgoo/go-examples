package mylogger

import "strings"

//自定义日志
type Level uint8

const (
	DebugLevel   Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

func GetLevelStr(level Level) string {
	leverStr := ""
	switch level {
	case DebugLevel:
		leverStr = "Debug"
	case InfoLevel:
		leverStr = "Info"
	case WarningLevel:
		leverStr = "Warn"
	case ErrorLevel:
		leverStr = "Error"
	case FatalLevel:
		leverStr = "Fatal"
	default:
		leverStr = "Debug"
	}

	return leverStr
}

func ParseLogLevel(levelStr string) Level {
	var level Level
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		level = DebugLevel
	case "info":
		level = InfoLevel
	case "warn":
		level = WarningLevel
	case "error":
		level = ErrorLevel
	case "fatal":
		level = FatalLevel
	default:
		level = DebugLevel
	}

	return level
}
