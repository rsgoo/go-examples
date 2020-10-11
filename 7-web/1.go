package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request info:", r)
	//fmt.Println("header:", r.Header)
	if r.URL.Path != "/" && r.URL.Path != "/index" {
		http.NotFound(w, r)
		return
	}

	fmt.Println("method:", r.Method)
	switch strings.ToLower(r.Method) {
	case "get":
		for key, val := range r.URL.Query() {
			fmt.Printf("%s: %s\n", key, val)
			w.Write([]byte(key))
			w.Write([]byte(val[0]))
		}

	case "post":
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", requestBody)
		w.Write([]byte("post"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

	//header 的设置需要放在数据相应之前

}
