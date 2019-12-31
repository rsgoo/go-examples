package main

import (
	"html/template"
	"net/http"
	"x-bookstore/controller"
)

//去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w, "")
}

func main() {

	//http.HandleFunc("/main", IndexHandler)
	http.HandleFunc("/main", controller.GetPageBooksByPrice)

	//目录结构
	//--x-bookstore
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

	//添加图书到购物车
	http.HandleFunc("/addBook2Cart",controller.AddBook2Cart)

	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)

	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)

	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)

	//更新购物车
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	//结账
	http.HandleFunc("/checkout", controller.Checkout)

	//获取所有的订单
	http.HandleFunc("/getOrders", controller.GetOrders)

	//获取订单详情
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)

	//获取登录用户的订单
	http.HandleFunc("/getMyOrder", controller.GetMyOrders)

	//发货
	http.HandleFunc("/sendOrder", controller.SendOrder)

	//确认收货
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	http.ListenAndServe(":8080", nil)
}
