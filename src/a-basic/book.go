package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func NewBook(title, author string) Book {
	return Book{
		Title:  title,
		Author: author,
	}
}

func main() {
	var book Book
	var handelCode string
	var title string
	var author string

	for {
		fmt.Println("请根据一下提示操作")
		fmt.Println("1：添加书籍")
		fmt.Println("2：编辑书籍")
		fmt.Println("3：展示书籍")
		fmt.Println("4：退出")
		fmt.Scanln(&handelCode)
		if handelCode == "1" {
			fmt.Println("请输入书籍名称")
			fmt.Scanln(&title)
			fmt.Println("请输入书籍作者")
			fmt.Scanln(&author)
			book = NewBook(title, author)
		} else if handelCode == "2" {
			fmt.Println("请输入修改的书籍名称")
			fmt.Scanln(&title)
			fmt.Println("请输入修改的书籍作者")
			fmt.Scanln(&author)
			book.Title = title
			book.Author = author
		} else if handelCode == "3" {
			fmt.Println(book)
		} else {
			fmt.Println("退出了, 拜拜")
			return
		}
	}
}
