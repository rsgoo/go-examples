package dao

import (
	"testing"
	"fmt"
	"xbookstore/model"
	"strconv"
)

func TestMain(m *testing.M) {
	fmt.Println("testMain")
	m.Run()
}

func TestBook(t *testing.T) {
	//t.Run("测试获取所有的图书", testGetBooks)
	//t.Run("测试添加图书", testGAddBooks)
	//t.Run("测试删除图书", testDeleteBook)
	//t.Run("测试图书详情", testGetBookById)
	//t.Run("测试图书修改", testUpdateBookById)
	//t.Run("测试图书获取", testGetPageBooks)
	//t.Run("测试session新增", testAddSession)
	//t.Run("测试session删除", testDelSession)
	t.Run("测试session获取", testGetSession)
}

func testGetSession(t *testing.T) {
	session, err := GetSession("01635090-a1f2-4139-5dba-0936e4437a6d")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(session)
	}
}

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "102403",
		UserName:  "大表哥",
		UserID:    3,
	}
	err := AddSession(sess)
	fmt.Println(err)
}

func testDelSession(t *testing.T) {
	err := DelSession("102403")
	fmt.Println(err)
}

func testGetPageBooks(t *testing.T) {
	books, err := GetPageBooksByPrice("1", "20", "30")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("TotalRecord:", books.TotalRecord)
	fmt.Println("PageNo:", books.PageNo)
	fmt.Println("TotalPageNo:", books.TotalPageNo)
	fmt.Println("TotalPageNo:", books.PageSize)

	for _, value := range books.Books {
		fmt.Println(*value)
	}
}

func testUpdateBookById(t *testing.T) {
	book := &model.Book{
		Title:   "测试图书修改",
		Author:  "郝林",
		Price:   64.50,
		Sales:   11019,
		Stock:   293, //库存
		ImgPath: "static/img/default.jpg",
	}
	err := UpdateBookById("53", book)
	fmt.Println(err)
}

func testDeleteBook(t *testing.T) {
	err := DeleteBook(strconv.Itoa(34))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("删除成功")
}

func testGetBookById(t *testing.T) {
	book, _ := GetBookById("2")
	fmt.Println(*book)
}

func testGAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "Go并发编程实战",
		Author:  "郝林",
		Price:   64.50,
		Sales:   11019,
		Stock:   293, //库存
		ImgPath: "static/img/default.jpg",
	}
	err := AddBook(book)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("添加成功")
	}
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	//遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%v本图书的信息是：%v\n", k+1, v)
	}
}

//func TestSaveUser(t *testing.T) {
//	err := SaveUser("admin", "123456", "admin@go.com")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}

//func TestCheckUserNameAndPassword(t *testing.T) {
//	user, _ := CheckUserNameAndPassword("admin", "123456")
//	fmt.Println(user)
//}

//func TestCheckUserName(t *testing.T) {
//	user, _ := CheckUserName("admin")
//	fmt.Println(user)
//}
