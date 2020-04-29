package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"

	// Get the data
	resp, err := http.Get(imgUrl)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("1:", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("2：", err)
	}

	ioutil.WriteFile("node.png", data, 0777)
}
