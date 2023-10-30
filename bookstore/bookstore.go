package bookstore

import (
	"errors"
	"fmt"
)

type Customer struct {
	name  string
	email string
}

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(b map[int]Book) []Book {
	all := []Book{}
	for _, v := range b {
		all = append(all, v)
	}
	return all
}

func GetBook(b map[int]Book, id int) (Book, error) {
	c, ok := b[id]
	if !ok {
		return Book{}, fmt.Errorf("ID %d did not exist", id)
	}
	return c, nil
}
