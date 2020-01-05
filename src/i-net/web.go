package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("./web.html")

	if err != nil {
		fmt.Println("read file err!", err)
	}

	//fmt.Fprintln(w, data)
	w.Write(data)
}

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe：", err)
	}
}
