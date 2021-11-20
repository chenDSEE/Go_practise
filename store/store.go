package store

import "errors"

// error type for book store error
var (
	ErrNotFound = errors.New("Book Can Not Found")
	ErrExisted  = errors.New("Book Already Exist")
	ErrBookInfo  = errors.New("Book Information Error")
)

type Book struct {
	Name      string `json:"Name"`
	ISBN      string `json:"ISBN"`
	Authors []string `json:"Authors"`
	Num       uint64 `json:"Num"`
}

type Store interface {
	CreateBook(*Book) error
	GetBook(string) (Book, error)	// get book information by book-name
}