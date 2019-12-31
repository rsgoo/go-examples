package main

import (
	"x-log/mylog"
	"time"
	"x-log/mylogger"
)

var logger mylogger.Logger

func main() {
	testMyLogger()
	testConsoleLogger()
}

func testConsoleLogger() {
	logger = mylogger.NewConsoleLogger("debug")
	logger.Debug("Debug user_id = %s", time.Now().Month())
}

func testMyLogger() {
	logFilePath := "C:\\gitNote\\Go-tour\\src\\x-use"
	logFileName := "myLogger.log"
	logger = mylogger.NewFileLogger("debug", logFilePath, logFileName)
	defer logger.Close()
	for i := 0; i < 100; i++ {
		logger.Debug("Debug user_id = %v", time.Now().Minute())
	}
}

func testMyLog() {
	logFilePath := "C:\\gitNote\\Go-tour\\src\\x-use"
	logFileName := "test.log"
	debugLogger := mylog.NewFileLogger(mylog.DEBUG, logFilePath, logFileName)
	debugLogger.Debug("这是一个测试debug使用 user_id = %s", time.Now())
}
