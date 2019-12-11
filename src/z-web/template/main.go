package main

import (
	"net/http"
	"html/template"
)

func main() {

	http.HandleFunc("/temp", templateHandle)
	http.ListenAndServe(":8080", nil)
}

func templateHandle(writer http.ResponseWriter, request *http.Request) {
	//fmt.Fprintln(writer, "嘿嘿")
	//解析模板
	//temp, err := template.ParseFiles("index.html")

	temp := template.Must(template.ParseFiles("index.html"))
	temp.Execute(writer, "这是来自Main.go 的响应数据")
}
