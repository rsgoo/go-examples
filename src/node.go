package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/httplib"
	"io"
	"log"
	"net/http"
)


const HOST = "https://www.codecasts.com/"
const LOGIN_PAGE_URL = "user/login?redirect_url=https://www.codecasts.com"
const LOGIN_POST_URL = "user/login"
const USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"

var myCookie = []*http.Cookie{}
var myHeader = map[string]string{
	"User-Agent": USER_AGENT,
}

func main() {

	GetLoginInfo()


}

func PostLoginInfo() {
	//url := HOST + LOGIN_POST_URL
}

func GetLoginInfo() {
	url := HOST + LOGIN_PAGE_URL

	req := httplib.Get(url)

	resp, err := req.SetUserAgent(USER_AGENT).Response()
	if err != nil {
		log.Fatal(err)
	}

	myCookie = resp.Cookies()

	respBytes, err := req.Bytes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp,respBytes)

	document, err := goquery.NewDocumentFromReader()

	fmt.Println(document)
}

func CookieSliceToString(oriCookie []string) string {

	str := ""
	for _, v := range oriCookie {
		str += v + ";"
	}
	return str
}