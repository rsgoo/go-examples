package pressure

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
	dstFile, e := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	//必须加上，否则不停的创建会导致内存溢出
	defer dstFile.Close()
	if e != nil {
		fmt.Println("err is ", e)
		panic("文件创建失败")
	}
	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(p)
	if err != nil {
		//fmt.Println("encode failed，err is ", err)
		return false
	} else {
		//fmt.Println("encode succeess")
		return true
	}
}

func DecodeJsonFileToPerson(fileName string) (*Person, error) {
	personPtr := &Person{}
	srcFile, e := os.OpenFile(fileName, os.O_RDONLY, 0666)
	defer srcFile.Close()
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
