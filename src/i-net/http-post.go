package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"net/url"
	"strings"
)

func HandlePostError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}

}

func main() {
	httpPost()
	httpPostForm()
	httpDo()
}

func httpPost() {
	contentType := "application/x-www-form-urlencoded";
	resp, err := http.Post(
		"http://httpbin.org/post?age=1&name=汤姆",
		contentType,
		//表单数据
		strings.NewReader("lang=golang&ver=1.13"),
	)
	HandlePostError(err, "http.Post")
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	HandlePostError(err, "ioutil.ReadAll")
	fmt.Println(string(bytes))
}

func httpPostForm() {
	resp, err := http.PostForm(
		"http://httpbin.org/post?age=1&name=汤姆",
		url.Values{
			"key": {"Value1", "Value2"},
			"id":  {"123"}},
	)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		"http://httpbin.org/post?age=1&name=汤姆",
		strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	//设置hander信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

//{
//  "args": {
//    "age": "1"
//  },
//  "data": "",
//  "files": {},
//  "form": {
//    "lang": "golang",
//    "ver": "1.13"
//  },
//  "headers": {
//    "Accept-Encoding": "gzip",
//    "Content-Length": "20",
//    "Content-Type": "application/x-www-form-urlencoded",
//    "Host": "httpbin.org",
//    "User-Agent": "Go-http-client/1.1"
//  },
//  "json": null,
//  "origin": "219.143.151.230, 219.143.151.230",
//  "url": "https://httpbin.org/post?age=1"
//}
