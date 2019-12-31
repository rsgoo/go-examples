package mylog

import (
	"runtime"
	"path"
	"fmt"
)

func GetCallerInfo(skip int) (fileName string, funcName string, line int, ok bool) {

	//skip 以当前为起始路径(0), 上一层为1，再上一层为2, 以此类推，直至找到最外层调用处
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", "", 0, false
	}

	fileName = path.Base(file)

	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)

	return fileName, funcName, line, true
}

func HandlerErr() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
