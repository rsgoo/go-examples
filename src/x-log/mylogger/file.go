package mylogger

import (
	"time"
	"os"
	"path"
	"fmt"
	"x-log/mylog"
	"strings"
)

//文件日志结构体
type fileLogger struct {
	fileLevel   Level
	filePath    string
	fileName    string
	logTime     string
	file        *os.File
	errFile     *os.File
	fileMaxSize int64 //日志文件最大容量
}

//文件打开
func (f *fileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic("文件打开失败")
	}
	f.file = fileObj

	errLogName := path.Join(f.filePath, "error_"+f.fileName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic("文件打开失败")
	}
	f.errFile = errFileObj
}

//文件日志结构体实例
func NewFileLogger(levelStr string, filePath, fileName string) *fileLogger {
	level := ParseLogLevel(levelStr)
	logObj := &fileLogger{
		fileLevel: level,
		filePath:  filePath,
		fileName:  fileName,
		logTime:   time.Now().Format("2006-01-02 15:04:05"),
		//fileMaxSize: 10 * 1024 * 1024, //10 - kb - mb
		fileMaxSize: 5 * 1024, //1 - kb - mb
	}
	logObj.initFile()
	return logObj
}

//将公用的记录日志方法功能封装为一个独立的方法
func (f *fileLogger) log(level Level, format string, args ...interface{}) {
	if f.fileLevel > level {
		fmt.Println("日志定义级别和调用方法不一致")
		return
	}
	//日志格式：[时间][文件:行号][函数名][日志级别]
	fileName, funcName, line, ok := mylog.GetCallerInfo(3)
	if !ok {
		return
	}

	logFormat := fmt.Sprintf("[%s] [%v:%d] [%v()] [%v] ", f.logTime, fileName, line, funcName, GetLevelStr(level))

	//日志基础信息
	baseMsg := fmt.Sprintf(format, args...)

	logMsg := logFormat + baseMsg

	//往文件中写数据前检查，检查当前文件的大小是否超过了 fileMaxSize
	fileInfo, err := f.file.Stat()
	if err != nil {
		return
	}

	fileSize := fileInfo.Size()
	if fileSize > f.fileMaxSize {
		//切分文件
		fileFullName := f.file.Name() //完整的文件路径
		fileFullName = strings.Replace(fileFullName, "/", "\\", -1)

		backupName := fmt.Sprintf("%s_%v.back", fileFullName, time.Now().Unix())

		//1: 把原来的文件关闭
		f.file.Close()

		//2：备份原来的文件
		os.Rename(fileFullName, backupName)

		//3: 新建一个文件
		logName := path.Join(f.filePath, f.fileName)
		fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			return
		}
		f.file = fileObj
	}

	fmt.Fprintln(f.file, logMsg)

	//如果日志级别是err 或者是fatal
	if level >= ErrorLevel {
		fmt.Fprintln(f.errFile, logMsg)
	}

}

//文件切割
func (f *fileLogger) splitLogFile(fileObj *os.File) {

}

func (f *fileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

func (f *fileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

func (f *fileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

func (f *fileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

func (f *fileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)
}

func (f *fileLogger) Close()  {
	fmt.Println("文件关闭")
	f.file.Close()
	f.errFile.Close()
}
