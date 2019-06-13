package model

import "errors"

var ErrStockNotEnough = errors.New("stock is not enough")

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

func CreateBook(name string, total int, author, createTime string) (b *Book) {
	b = &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createTime,
	}
	return
}

func (b *Book) CanBorrow(c int) bool {
	return b.Total > c
}

//借书
func (b *Book) Borrow(c int) (err error) {
	if b.CanBorrow(c) == false {
		err = ErrStockNotEnough
		return
	}

	b.Total -= c
	return
}

//还书
func (b *Book) GiveBack(c int){
	b.Total += c
	return
}
