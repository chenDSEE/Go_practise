package store

import (
	"bookstore/store"
	"bookstore/store/factory"
	"errors"
	"sync"
)

func init() {
	factory.Register("mem", &memStore{
		db: make(map[string]*store.Book),
	})
}

type memStore struct {
	sync.RWMutex
	db map[string]*store.Book
}

func (mem *memStore)CreateBook(book *store.Book) error {
	if book == nil {
		return errors.New("memDB: createBook without book-info")
	}

	if book.Name == "" || book.Authors == nil || book.ISBN == "" || book.Num == 0 {
		return store.ErrBookInfo
	}

	mem.Lock()
	defer mem.Unlock()

	if _, exist := mem.db[book.ISBN]; exist {
		return store.ErrExisted
	}

	newBook := *book
	mem.db[book.ISBN] = &newBook

	return nil
}


func (mem *memStore)GetBook(name string) (store.Book, error) {
	mem.RLock()
	defer mem.RUnlock()

	book, exist := mem.db[name]
	if !exist {
		return store.Book{}, store.ErrNotFound
	}

	return *book, nil
}


