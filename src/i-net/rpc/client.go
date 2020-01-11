package main

import (
	"net/rpc"
	"fmt"
)

func SyncCall() {
//func main() {

	//1：连接服务器服务
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println("dialHttp err:", err)
		return
	}
	var req float32 = 3

	var rsp float32

	//同步调用
	err = client.Call("MathUtil.CalculateCircleArea", req, &rsp)
	if err != nil {
		fmt.Println("call err: ", err)
		return
	}
	fmt.Println(req)
	fmt.Println(rsp)
}

//使用异步调用的方式
//func AsyncCall() {
func main() {
	//1：连接服务器服务
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println("dialHttp err:", err)
		return
	}
	var req float32 = 10

	var rsp float32

	//异步调用
	syncCall := client.Go("MathUtil.CalculateCircleArea", req, &rsp, nil)

	//调用完成标识管道
	replayDone := <-syncCall.Done

	fmt.Println(syncCall.Args)

	fmt.Println(replayDone)

	fmt.Println(rsp)

}
