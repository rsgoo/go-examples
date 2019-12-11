package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func HandleGetError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	url := "https://www.zhihu.com/hot"
	resp, err := http.Get(url)
	HandleGetError(err, "http.Get")
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	HandleGetError(err, "ioutil.ReadAll")
	fmt.Println(string(bytes))

}
