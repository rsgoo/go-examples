package model

import "errors"

var ErrNotFoundBook = errors.New("not found the book")

type Student struct {
	Name  string
	Grade string
	Id    string
	Sex   string
	Books []*Book
}

//借书记录
type BorrowItem struct {
	book *Book
	num  int
}

func CreateStudent(name, grade, id, sex string) *Student {
	stu := &Student{
		Name:  name,
		Grade: grade,
		Id:    id,
		Sex:   sex,
	}
	return stu
}

func (s *Student) AddBook(b *BorrowItem) {
	s.Books = append(s.Books, b.book)
}

func (s *Student) DeleteBook(b *BorrowItem) (err error) {
	for i := 0; i < len(s.Books); i++ {
		if s.Books[i].Name == b.book.Name {
			//删除图书
			if b.num == s.Books[i].Total {
				front := s.Books[0:i]
				end := s.Books[i+1:]
				front = append(front, end...)
				s.Books = front
				return
			}
			s.Books[i].Total -= b.num
		}
	}
	return
}
