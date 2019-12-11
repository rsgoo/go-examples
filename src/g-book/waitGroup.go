package main

import (
	"sync"
	"net/http"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	var urlList = []string{
		"https://www.baidu.com/",
		"https://www.ithome.com/",
		"https://www.jd.com/",
	}

	for _, url := range urlList {
		//等待组的计数器+1
		wg.Add(1)
		go func(urlLink string) {
			//等待组的计数器-1
			//defer wg.Add(-1)
			defer wg.Done()
			_, err := http.Get(urlLink)
			fmt.Println(urlLink, "---", err)

		}(url)
	}
	//Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
	fmt.Println("main over")
}
