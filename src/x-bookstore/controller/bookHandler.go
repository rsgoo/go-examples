package controller

import (
	"net/http"
	"x-bookstore/dao"
	"html/template"
	"x-bookstore/model"
	"strconv"
)

//去首页
func Index(w http.ResponseWriter, r *http.Request) {

	//获取页码
	pageNo := r.FormValue("pageNo")

	//设置 pageNo 默认值
	if pageNo == "" {
		pageNo = "1"
	}

	page, _ := dao.GetPageBooks(pageNo)

	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w, page)
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {

	//获取页码
	pageNo := r.FormValue("pageNo")

	//设置 pageNo 默认值
	if pageNo == "" {
		pageNo = "1"
	}

	//books, _ := dao.GetBooks()
	//获取带分页的图书数据
	pageBook, _ := dao.GetPageBooks(pageNo)
	//fmt.Println("pageBook.PageNo", pageBook.PageNo)
	//fmt.Println("pageBook.PageSize", pageBook.PageSize)
	//fmt.Println("pageBook.TotalPageNo", pageBook.TotalPageNo)
	//fmt.Println("pageBook.TotalRecord", pageBook.TotalRecord)
	//fmt.Println("IsHasPrev", pageBook.IsHasPrev())
	//fmt.Println("IsHasNext", pageBook.IsHasNext())
	//fmt.Println("GetNextPageNo", pageBook.GetNextPageNo())
	//fmt.Println("GetPrevPageNo", pageBook.GetPrevPageNo())
	//os.Exit(65)

	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, pageBook)
}

//获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {

	//获取页码
	pageNo := r.FormValue("pageNo")

	//设置 pageNo 默认值
	if pageNo == "" {
		pageNo = "1"
	}

	minPrice := r.FormValue("min_price")
	maxPrice := r.FormValue("max_price")

	//这里必须要声明pageBook类型
	var pageBook *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用bookdao中获取带分页的图书的函数
		pageBook, _ = dao.GetPageBooks(pageNo)
	} else {
		pageBook, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格范围设置到page中
		pageBook.MinPrice = minPrice
		pageBook.MaxPrice = maxPrice
	}

	//获取cookie
	//获取指定的cookie
	/*
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := dao.GetSession(cookieValue)
		if session.UserID > 0 {
			pageBook.IsLogin = true
			pageBook.Username = session.UserName
		}
	}
	*/

	flag, session := dao.IsLogin(r)
	if flag {
		pageBook.IsLogin = true
		pageBook.Username = session.UserName
	}

	//books, _ := dao.GetBooks()
	//获取带分页的图书数据
	//pageBook, _ := dao.GetPageBooksByPrice(pageNo,"1","1")
	//fmt.Println("pageBook.PageNo", pageBook.PageNo)
	//fmt.Println("pageBook.PageSize", pageBook.PageSize)
	//fmt.Println("pageBook.TotalPageNo", pageBook.TotalPageNo)
	//fmt.Println("pageBook.TotalRecord", pageBook.TotalRecord)
	//fmt.Println("IsHasPrev", pageBook.IsHasPrev())
	//fmt.Println("IsHasNext", pageBook.IsHasNext())
	//fmt.Println("GetNextPageNo", pageBook.GetNextPageNo())
	//fmt.Println("GetPrevPageNo", pageBook.GetPrevPageNo())
	//os.Exit(65)

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, pageBook)
}

//跳转到添加页面
func AddBook(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
	t.Execute(w, "")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("bookId")
	dao.DeleteBook(id)
	GetPageBooks(w, r)
}

//去更新或者添加图书
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//获取到图书id
	bookId := r.FormValue("bookId")

	//查询图书详情
	book, _ := dao.GetBookById(bookId)

	//存在book_id，更新图书信息
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		//不存在book_id，新增图书信息
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}

}

//更新或是添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PostFormValue("bookId")

	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	priceTrans, _ := strconv.ParseFloat(price, 10)
	salesTrans, _ := strconv.Atoi(sales)
	stockTrans, _ := strconv.Atoi(stock)
	bookIdTrans, _ := strconv.Atoi(bookId)

	var book = &model.Book{
		ID:     int(bookIdTrans),
		Title:  title,
		Author: author,
		Price:  priceTrans,
		Sales:  salesTrans,
		Stock:  stockTrans,
	}

	if book.ID > 0 {
		//更新图书
		dao.UpdateBookById(bookId, book)
	} else {
		//添加
		dao.AddBook(book)
	}
	GetPageBooks(w, r)
}
