package book

import "errors"

type Book struct {
	Name        string
	Total       int
	Author      string
	PublishDate string
}

func CreateBook(name string, author string, publishdate string, total int) (book *Book) {
	return &Book{
		Name:        name,
		Total:       total,
		Author:      author,
		PublishDate: publishdate,
	}
}

func (book *Book) canBorrow(c int) bool {
	return book.Total >= c
}

func (book *Book) Barrow(c int) (err error) {
	if book.canBorrow(c) == false {
		err = errors.New("Not enough books to brrow!")
		return
	}
	book.Total -= c
	return
}

func (book *Book) Back(c int) {
	book.Total += c
	return
}
