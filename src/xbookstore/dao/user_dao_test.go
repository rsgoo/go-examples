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
	//t.Run("测试session获取", testGetSession)
	//t.Run("测试添加购物车", testAddCart)
	//t.Run("测试获取购物项详情", testGetCartItemByBookIDAndCartID)
	//t.Run("测试获取购物车详情", testGetCartByUserID)
	t.Run("测试获取购物车详情", testUpdateBookCount)
}

func testUpdateBookCount(t *testing.T) {

}

func testGetCartByUserID(t *testing.T) {
	cart, err := GetCartByUserID(1)
	if err != nil {
		fmt.Println("err is:", err)
	} else {
		fmt.Println(cart.CartID)
		fmt.Println(cart.TotalAmount)
		fmt.Println(cart.TotalCount)
		for _, value := range cart.CartItems {
			fmt.Println(*value.Book)
		}
	}
}

func testGetCartItemByBookIDAndCartID(t *testing.T) {
	//cartItem, err := GetCartItemByBookIDAndCartID(string(1), string(66668888))
	cartItem, err := GetCartItemByBookIDAndCartID("1", "66668888")
	if err != nil {
		fmt.Println("err is：", err)
	} else {
		fmt.Println(cartItem)
	}

	fmt.Println(cartItem.CartItemId)
	fmt.Println(cartItem.Count)
	fmt.Println(cartItem.Amount)
	fmt.Println(cartItem.CartID)
	fmt.Println(*cartItem.Book)
	fmt.Println(cartItem.Ctime)
	fmt.Println(cartItem.Mtime)
}

func testGetCartItemsByCartID(t *testing.T) {

}

func testAddCart(t *testing.T) {
	//设置要买的第一本书
	book := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	//设置要买的第二本书
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	//创建一个购物项切片
	var cartItems []*model.CartItem
	//创建两个购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem2)

	//创建购物车
	cart := &model.Cart{
		CartID:    "66668888",
		CartItems: cartItems,
		UserID:    1,
	}
	//将购物车插入到数据库中
	err := AddCart(cart)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("添加成功")
	}

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
