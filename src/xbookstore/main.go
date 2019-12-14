package main

import (
	"html/template"
	"net/http"
	"xbookstore/controller"
)

//去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w, "")
}

func main() {

	//http.HandleFunc("/main", IndexHandler)
	http.HandleFunc("/main", controller.Index)

	//目录结构
	//--xbookstore
	//----views
	//------index.html
	//------pages
	//--------user
	//----------login.html
	//----------login_success.html
	//------static
	//--------css
	//----------style.css

	//设置处理静态资源，如css和js文件
	//将 "/static/" 的寻址路径映射为 "/views/static/"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))

	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	//登录
	http.HandleFunc("/login", controller.Login)

	//去注销
	http.HandleFunc("/logout", controller.Logout)

	//注册
	http.HandleFunc("/register", controller.Register)

	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	//获取待分页的所有图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)

	//
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)

	//添加图书-数据取出渲染html页面
	//http.HandleFunc("/addBook", controller.AddBook)

	//更新图书-数据取出渲染html页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)

	//增加 或者 修改图书
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)

	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)

	http.ListenAndServe(":8080", nil)
}
