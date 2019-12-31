package mylogger

import (
	"fmt"
	"time"
	"x-log/mylog"
	"os"
)

type consoleLogger struct {
	consoleLevel Level
	logTime      string
}

//文件日志结构体实例
func NewConsoleLogger(levelStr string) *consoleLogger {
	level := ParseLogLevel(levelStr)
	logObj := &consoleLogger{
		consoleLevel: level,
		logTime:      time.Now().Format("2006-01-02 15:04:05"),
	}
	return logObj
}

//将公用的记录日志方法功能封装为一个独立的方法
func (c *consoleLogger) log(level Level, format string, args ...interface{}) {
	if c.consoleLevel > level {
		fmt.Println("日志定义级别和调用方法不一致")
		return
	}
	//日志格式：[时间][文件:行号][函数名][日志级别]
	ConsoleName, funcName, line, ok := mylog.GetCallerInfo(3)
	if !ok {
		return
	}

	logFormat := fmt.Sprintf("[%s] [%v:%d] [%v()] [%v] ", c.logTime, ConsoleName, line, funcName, GetLevelStr(level))

	//日志基础信息
	baseMsg := fmt.Sprintf(format, args...)

	logMsg := logFormat + baseMsg

	fmt.Println(logMsg)
	os.Exit(1)

	fmt.Fprintln(os.Stdout, logMsg)

	//如果日志级别是err 或者是fatal
	if level >= ErrorLevel {
		fmt.Fprintln(os.Stdout, logMsg)
	}

}

func (c *consoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

func (c *consoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

func (c *consoleLogger) Warn(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

func (c *consoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

func (c *consoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel, format, args...)
}

func (f *consoleLogger) Close()  {
}