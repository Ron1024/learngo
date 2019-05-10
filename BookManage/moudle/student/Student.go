package student

import (
	"errors"
	"learngo/BookManage/moudle/book"
)

type Student struct {
	Name          string
	Class         string
	ID            string
	Sex           string
	BorrowedBooks []*book.Book
}

type BorrowItem struct {
	book *book.Book
	num  int
}

func CreateStudent(name, class, id, sex, borrowedBooks string) *Student {
	return &Student{
		Name:  name,
		ID:    id,
		Class: class,
		Sex:   sex,
	}
}

func (stu *Student) AddBook(b *BorrowItem) {
	stu.BorrowedBooks = append(stu.BorrowedBooks, b.book)
}

func (stu *Student) DelBook(b *BorrowItem) (err error) {
	for i := 0; i < len(stu.BorrowedBooks); i++ {
		if stu.BorrowedBooks[i].Name == b.book.Name {
			if stu.BorrowedBooks[i].Total == b.book.Total {
				front := stu.BorrowedBooks[0:i]
				left := stu.BorrowedBooks[i+1:]
				front = append(front, left...)
				stu.BorrowedBooks = front
				return
			}
		}
	}
	return errors.New("There is no book to return")
}
