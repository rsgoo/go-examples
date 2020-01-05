package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(rw http.ResponseWriter, rq *http.Request) {
	//fmt.Fprintln(rw, "请求method:", rq.Method)
	//fmt.Fprintln(rw, "请求path:", rq.URL.Path)
	//fmt.Fprintln(rw, "请求query:", rq.URL.RawQuery)
	//fmt.Fprintln(rw, "请求header:", rq.Header)
	//fmt.Fprintln(rw, "请求header:", rq.Header.Get("Accept-Encoding"))
	//fmt.Fprintln(rw, "请求header:", rq.Header["Accept-Encoding"])
	//fmt.Fprintln(rw, "请求Host:", rq.Host)

	rq.ParseForm()
	fmt.Fprintln(rw, "表单请求Form:", rq.Form)
	fmt.Fprintln(rw, "表单请求PostForm:", rq.PostForm)
	fmt.Fprintln(rw, "表单请求PostForm.Get:", rq.PostForm.Get("user_name"))
	fmt.Fprintln(rw, "表单请求FormValue:", rq.FormValue("user_pass"))

	//获取请求体中的长度
	//reqLength := rq.ContentLength

	//body := make([]byte, reqLength)

	//将请求体中的内容读取到body中
	//rq.Body.Read(body)

	//fmt.Fprintln(rw, "请求体中的内容：", string(body))
}
