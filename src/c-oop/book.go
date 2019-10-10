package main

import "fmt"

type Book struct {
	Name      string
	Price     float32
	Author    string
	CanBorrow bool
}

type Reader struct {
	ReaderID int
	Balance  float32 //押金余额
}

type Student struct {
	Reader
	Name  string
	Major string
}

type Teacher struct {
	Reader
	Name   string
	Course string
}

func (r *Reader) BorrowBook(b *Book) {
	if b.CanBorrow {
		fmt.Printf("%d 借阅了 %s\n", r.ReaderID, b.Name)
		b.CanBorrow = false
	} else {
		fmt.Printf("借阅失败,%s 已经被借阅完了\n", b.Name)
	}
}

func (r *Reader) ReturnBook(b *Book) {
	b.CanBorrow = true
	fmt.Printf("%v 归还了 %s\n", r.ReaderID, b.Name)
}

func (r *Reader) Penalty(amount float32) {
	r.Balance -= amount
	fmt.Printf("读者%v缴纳了罚金%.2f, 它的余额是%.2f\n", r.ReaderID, amount, r.Balance)
}

func (s *Student) Study() {
	fmt.Printf("%s正在学习\n", s.Name)
}

func (t *Teacher) Teaching() {
	fmt.Printf("%s正在教学\n", t.Name)
}

func (t *Teacher) Penalty(amount float32) {
	t.Balance -= 0
	fmt.Printf("读者%s缴纳了罚金%.2f, 它的余额是%.2f\n", t.ReaderID, amount, t.Balance)

}

func main() {
	book1 := Book{}
	book1.Name = "《西游记》"
	book1.Author = "吴承恩"
	book1.Price = 39.99
	book1.CanBorrow = true

	book2 := Book{}
	book2.Name = "《活着》"
	book2.Author = "余华"
	book2.Price = 29.99
	book2.CanBorrow = true

	book3 := Book{
		Name:      "《黑客与画家》",
		Author:    "保罗格雷厄姆",
		Price:     49.99,
		CanBorrow: true,
	}

	book4ptr := new(Book)
	book4ptr.Name = "《新华字典》"
	book4ptr.Author = "匿名"
	book4ptr.Price = 9.99
	book4ptr.CanBorrow = true

	fmt.Println(book1, book2, book3, book4ptr)
	fmt.Println()

	//创建学生
	student1 := Student{
		Reader: Reader{001, 100},
		Name:   "张全蛋",
		Major:  "质检",
	}
	fmt.Printf("%+v\n", student1)

	//创建老师
	teacher1 := Teacher{
		Reader: Reader{002, 0},
		Name:   "张老师",
		Course: "Go编程",
	}
	fmt.Printf("%+v\n", teacher1)

	//学习
	student1.Study()

	//教学
	teacher1.Teaching()

	//借书
	student1.BorrowBook(&book2)
	teacher1.BorrowBook(&book2)
	student1.ReturnBook(&book2)
	student1.Penalty(10)
	fmt.Println("剩下的余额：", student1.Balance)
	teacher1.BorrowBook(&book2)

}
