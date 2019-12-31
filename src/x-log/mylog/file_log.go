package mylog

import (
	"fmt"
	"os"
	"time"
)

//文件日志结构体
type fileLogger struct {
	level       int //日志级别
	logFilePath string
	logFileName string
	logFile     *os.File
	logTime     string
}

//文件日志结构体实例
func NewFileLogger(level int, logFilePath, logFileName string) *fileLogger {
	fileObj := &fileLogger{
		level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
		logTime:     time.Now().Format("2006-01-02 15:04:05"),
	}
	fileObj.initFilerLogger()

	return fileObj
}

//文件打开
func (f *fileLogger) initFilerLogger() {
	filePath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic("os.OpenFile err: " + err.Error())
	}
	f.logFile = file
}

//记录debug级别的日志
func (f *fileLogger) Debug(format string, args ...interface{}) {

	if f.level > DEBUG {
		return
	}

	callFileName, callFuncName, callLine, ok := GetCallerInfo(2)
	if !ok {
		return
	}

	callInfo := fmt.Sprintf("[%s : %s : %d ]", callFileName, callFuncName, callLine)

	now := time.Now().Format("[2006-01-02 15:04:05.000]")
	logMsg := fmt.Sprintf("%s %s %s", GetLevelStr(DEBUG), now, callInfo)

	fmt.Fprintf(f.logFile, logMsg+format, args...)
	fmt.Fprintln(f.logFile)
}

//记录debug级别的日志
func (f *fileLogger) Info(msg string) {
	n, err := f.logFile.WriteString(f.logTime + " info:" + msg + "\n")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
