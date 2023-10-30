package bookstore

import "errors"

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
func GetAllBooks(b []Book) []Book {
	return b
}
func GetBook(b []Book, id int) Book {
	for _, c := range b {
		if c.ID == id {
			return c
		}
	}
	return Book{}
}
