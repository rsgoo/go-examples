package main

import (
	"os"
	"errors"
	"fmt"
)

type fileWriter struct {
	file *os.File
}

func (f *fileWriter) SetFile(fileName string) (err error) {
	if f.file != nil {
		f.file.Close()
	}

	f.file, err = os.Create(fileName)
	return err
}

func (f *fileWriter) Write(data interface{}) (n int, err error) {
	if f.file == nil {
		return 0, errors.New("file not created")
	}
	defer f.file.Close()
	str := fmt.Sprintf("%v\n", data)

	n, err = f.file.Write([]byte(str))
	return n, err
}

func NewFileWriter() *fileWriter {
	return &fileWriter{}
}

func main() {
	file := NewFileWriter()
	err := file.SetFile("inscode.txt")
	if err != nil {
		fmt.Println("setFile err:", err)
	}

	n, err := file.Write("雨醉风尘")
	if err != nil {
		fmt.Println("fileWrite err:", err)
	}
	fmt.Println(n)
}
