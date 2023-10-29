package bookstore_test

import (
	"happy-fun-book/bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Hang Tuai",
		Author: "M Nusir",
		Copies: 10}
}
