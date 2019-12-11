package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/set_cookie", SetCookieOne)
	//http.HandleFunc("/set_cookie", SetCookie)
	http.HandleFunc("/get_cookie", GetCookies)
	http.ListenAndServe(":8088", nil)
}

//设置cookie方式1
func SetCookieOne(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "name",
		Value:    "linus",
		HttpOnly: true,
		MaxAge:   120,	//设置有效时长
	}
	//将cookie发送给浏览器
	w.Header().Set("Set-Cookie", cookie1.String())

	//再次添加一个cookie
	cookie2 := http.Cookie{
		Name:     "language",
		Value:    "golang",
		HttpOnly: true,
	}
	w.Header().Add("set-cookie", cookie2.String())
	w.Header().Add("customer", "this is test page")
}

//设置cookie方式2
func SetCookie(w http.ResponseWriter, r *http.Request) {

	//再次添加一个cookie
	cookie := http.Cookie{
		Name:     "database",
		Value:    "mysql",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

//获取请求头中所有的cookie
func GetCookies(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	//fmt.Println(cookies)
	headers := r.Header
	for key, value := range headers {
		fmt.Printf("%v : %v\n", key, value)
	}
	fmt.Println("cookies:", cookies)

	//获取指定的cookie
	customer, _ := r.Cookie("database")
	fmt.Println(*customer)
}
