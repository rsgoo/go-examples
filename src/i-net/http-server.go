package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func HandleHttpError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc(
		"/hello",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Hello,World\n"))
			//writer.Write([]byte("请求长度" + string(request.ContentLength)))
			writer.Write([]byte("请求Host: " + string(request.Host)))
			writer.Write([]byte("请求Method: " + string(request.Method)))
			writer.Write([]byte("请求Proto: " + request.Proto))
		},
	)

	http.HandleFunc(
		"/163",
		func(writer http.ResponseWriter, request *http.Request) {
			musicBytes, _ := ioutil.ReadFile("music163.html")
			writer.Write(musicBytes)

		},
	)

	http.ListenAndServe("0.0.0.0:8080", nil)
}
