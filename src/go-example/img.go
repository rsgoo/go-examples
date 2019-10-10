package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
)

func main() {
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"

	// Get the data
	resp, err := http.Get(imgUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create("img1.jpg")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	copyBytes, err := io.Copy(out, resp.Body)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(copyBytes)
}
