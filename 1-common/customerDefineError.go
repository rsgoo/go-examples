package main

import "fmt"

type parseError struct {
	FileName string
	Line     int
}

func (e *parseError) Error() string {
	return fmt.Sprintf("error in %s:%d\n", e.FileName, e.Line)
}

//这里的返回值时 error interface
func NewParseError(fileName string, line int) *parseError {
	return &parseError{FileName: fileName, Line: line}
}

//自定义错误及简单使用
func main() {
	var e error
	e = NewParseError("node.go", 1)
	fmt.Println(e.Error())
	switch detail := e.(type) {
	case *parseError:
		fmt.Println(detail.FileName, detail.Line)
	default:
		fmt.Println("other error")
	}
}
