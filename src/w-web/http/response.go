package main

import (
	"net/http"
	"strconv"
	"encoding/json"
)

type User struct {
	Name     string
	Age      int
	Language string
}

func main() {
	//http.HandleFunc("/hello", jsonHandle)
	http.HandleFunc("/hello", redirectHandle)
	http.ListenAndServe(":8080", nil)
}

//http://127.0.0.1:8080/hello?name=linus&age=111&language=golang
//json解析
func jsonHandle(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	//json响应
	res.Header().Set("Content-Type", "application/json")
	//res.Write([]byte("这是一条响应信息"))

	age, _ := strconv.Atoi(req.Form.Get("age"))
	user := User{
		Name:     req.Form.Get("name"),
		Age:      age,
		Language: req.Form.Get("language"),
	}
	jsonStr, _ := json.Marshal(user)
	res.Write(jsonStr)
}

//重定向
func redirectHandle(res http.ResponseWriter, req *http.Request) {

	//重定向地址
	res.Header().Set("Location", "https://www.baidu.com")

	//重定向状态码
	res.WriteHeader(302)
}
