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
func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Sparkling Wota",
		Author: "Charles Dickens",
		Copies: 2}
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("Want %d copies from %d stock, got %d", want, b.Copies, got)
	}
}
func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Puteri Gunung Larut Matang",
		Author: "Joyah",
		Copies: 0}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Errorf("Want %v, got nil", err)
	}
}
