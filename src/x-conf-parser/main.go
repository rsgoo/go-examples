package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"reflect"
	"errors"
	"strconv"
)

type Config struct {
	FilePath string `conf:"file_path"`
	FileName string `conf:"file_name"`
	FaxSize  string `conf:"max_size"`
}

//从conf文件中赋值给结构体指针
func parseConf(fileName string, result interface{}) error {
	refT := reflect.TypeOf(result)
	refV := reflect.ValueOf(result)

	if refT.Kind() != reflect.Ptr {
		err := errors.New("result必须是一个指针")
		return err
	}

	refTElem := refT.Elem()
	if refTElem.Kind() != reflect.Struct {
		err := errors.New("result的底层必须是结构体")
		return err
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("打开配置文件 %s 失败", fileName)
		return err
	}

	lineSlice := strings.Split(string(data), "\n")
	for index, value := range lineSlice {

		lineContent := strings.TrimSpace(value)

		if strings.HasPrefix(lineContent, "#") {
			continue
		}

		if len(lineContent) == 0 {
			continue
		}

		equalIndex := strings.Index(lineContent, "=")

		if equalIndex < 0 {
			err = fmt.Errorf("第%d行语法错误", index+1)
			return err
		}

		//根据 = 分割 字符串
		//explodeStr := strings.Split(lineContent, "=")
		confKey := strings.TrimSpace(lineContent[:equalIndex])
		confVal := strings.TrimSpace(lineContent[equalIndex+1:])
		if len(confKey) == 0 {
			err = fmt.Errorf("第%d行语法错误", index+1)
			return err
		}

		//遍历结构体的每一个字段和key比较，
		for i := 0; i < refTElem.NumField(); i++ {
			field := refTElem.Field(i)
			tag := field.Tag.Get("conf")
			if tag == confKey {
				//匹配上之后把value赋值给结构体
				fieldType := field.Type
				switch fieldType.Kind() {
				case reflect.String:
					fieldVal := refV.Elem().FieldByName(field.Name)
					fieldVal.SetString(confVal)
				case reflect.Int64:
					confValInt64, _ := strconv.ParseInt(confVal, 64, 10)
					refV.Elem().Field(i).SetInt(confValInt64)
				case reflect.Bool:
					confValBool, _ := strconv.ParseBool(confVal)
					refV.Elem().Field(i).SetBool(confValBool)
				}
			}
		}
	}
	return nil
}

func main() {

	var conf = &Config{}
	err := parseConf("parser.conf", conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", *conf)
}
