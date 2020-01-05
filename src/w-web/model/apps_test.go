package model

import (
	"testing"
	"fmt"
)

func TestApp_Create(t *testing.T) {
	fmt.Println("测试添加记录")
	app := &App{}
	app.Create()
	app.Add()
}
