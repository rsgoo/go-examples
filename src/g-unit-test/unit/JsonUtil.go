package unit

import (
	"os"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Rmb    float64
	Gender string
	Hobby  []string
}

func EncodePersonToJsonFile(p *Person, fileName string) bool {
	dstFile, e := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		panic("文件创建失败")
	}
	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(p)
	if err != nil {
		fmt.Println("encode failed，err is ", err)
		return false
	} else {
		fmt.Println("encode succeess")
		return true
	}
}

func DecodeJsonFileToPerson(fileName string) (*Person, error) {
	personPtr := &Person{}
	srcFile, e := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if e != nil {
		panic("文件不存在")
	}

	decoder := json.NewDecoder(srcFile)
	err := decoder.Decode(personPtr)
	if err != nil {
		fmt.Println("decode failed, err is ", err)
		return nil, err
	}

	return personPtr, nil

}
