package main

import (
	"zinx-demo/protobufDemo"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main() {
	//定义一个Person对象
	Person := &pb.Person{
		Name:   "雨醉风尘",
		Age:    26,
		Emails: []string{"11@qq.com", "22@qq.com"},
		Phones: []*pb.PhoneNumber{
			{
				Numer: "13111011101",
				Type:  pb.PhoneType_MOBILE,
			},
			{
				Numer: "13111011102",
				Type:  pb.PhoneType_HOME,
			},
			{
				Numer: "13111011103",
				Type:  pb.PhoneType_WORK,
			},
		},
	}

	//将person对象 进行序列化
	data, err := proto.Marshal(Person)
	if err != nil {
		fmt.Println("err is ", err)
		return
	}

	//解码
	newData := &pb.Person{}
	err = proto.Unmarshal(data, newData)

	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Println(Person)
	fmt.Println(newData)
}
