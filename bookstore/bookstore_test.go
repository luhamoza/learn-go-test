package bookstore_test

import (
	"github.com/google/go-cmp/cmp"
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
func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"}, {Title: "The Power of Go"}}
	want := []bookstore.Book{
		{Title: "For the Love of Go"}, {Title: "The Power of Go"},
	}
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{{Title: "For the Love of Go", ID: 1}, {Title: "The Power of Go", ID: 2}}
	want := bookstore.Book{Title: "The Power of Go", ID: 2}
	got := bookstore.GetBook(catalog, 2)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
