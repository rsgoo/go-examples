package dao

import (
	"xbookstore/model"
	"xbookstore/utils"
	"strconv"
)

//获取全部图书
func GetBooks() ([]*model.Book, error) {
	//sqlStr := "select id, title, author, price, sales, stock, img_path, ctime, mtime from books"
	sqlStr := "select * from books"
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book结构体的每一条记录赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &book.Ctime, &book.Mtime)
		books = append(books, book)
	}

	return books, nil
}

//增加图书
func AddBook(b *model.Book) error {
	sqlStr := "insert into books(title, author, price, sales, stock, img_path) values(?, ?, ?, ?, ?, ?)"
	result, err := utils.DB.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

//删除图书
func DeleteBook(bookId string) error {
	sqlStr := "delete from books where id = ?"
	_, err := utils.DB.Exec(sqlStr, bookId)
	if err != nil {
		return err
	}

	return nil
}

//获取图书详情
func GetBookById(bookId string) (*model.Book, error) {
	sqlStr := "select * from books where id = ?"

	row := utils.DB.QueryRow(sqlStr, bookId)

	book := &model.Book{}

	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &book.Ctime, &book.Mtime)

	return book, nil
}

//更新图书信息
func UpdateBookById(bookId string, book *model.Book) (error) {
	sqlStr := "update books set title = ?, author = ?, price = ?, sales = ?, stock = ?, img_path = ? where id = ?"

	_, err := utils.DB.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, "static/img/default.jpg", bookId)

	if err != nil {
		return err
	}
	return nil
}

//获取分页的图书
func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)

	//获取图书总数
	sqlStr := "select count(*) from books"
	row := utils.DB.QueryRow(sqlStr)

	//总记录数
	var totalRecord int64

	row.Scan(&totalRecord)

	//设置每一页的显示条数
	var pageSize int64 = 10

	//设置一个变量接收总页数
	var totalPageNo int64

	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页的图书
	sqlStr2 := "select * from books limit ?,?"

	rows, err := utils.DB.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var books []*model.Book

	for rows.Next() {
		var book = &model.Book{}

		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &book.Ctime, &book.Mtime)
		books = append(books, book)
	}

	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, nil
}

//获取带分页和价格范围的图书
func GetPageBooksByPrice(pageNo string, minPrice, maxPrice string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)

	//获取图书总数
	sqlStr := "select count(*) from books where price between ? and ? "
	row := utils.DB.QueryRow(sqlStr, minPrice, maxPrice)

	//总记录数
	var totalRecord int64

	row.Scan(&totalRecord)

	//设置每一页的显示条数
	var pageSize int64 = 10

	//设置一个变量接收总页数
	var totalPageNo int64

	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页的图书
	sqlStr2 := "select * from books where price between ? and ? limit ?,?"

	rows, err := utils.DB.Query(sqlStr2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var books []*model.Book

	for rows.Next() {
		var book = &model.Book{}

		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &book.Ctime, &book.Mtime)
		books = append(books, book)
	}

	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, nil
}


//GetBookByID 根据图书的id从数据库中查询出一本图书
func GetBookByID(bookID string) (*model.Book, error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	//执行
	row := utils.DB.QueryRow(sqlStr, bookID)
	//创建Book
	book := &model.Book{}
	//为book中的字段赋值
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}