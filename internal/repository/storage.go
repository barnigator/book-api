package repository

import (
	"errors"
	"sync"

	"github.com/barnigator/book-api/internal/entity"
)

type Storage struct {
	Books map[string]entity.Book
	mu    *sync.RWMutex
}

func (st *Storage) Create(book entity.Book) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	if _, ok := st.Books[book.ID]; ok {
		return errors.New("Book already exists")
	} else {
		st.Books[book.ID] = book
	}

	return nil
}

func (st *Storage) GetAll() []entity.Book {
	st.mu.RLock()
	defer st.mu.RUnlock()

	books := make([]entity.Book, 0, len(st.Books))

	for _, book := range st.Books {
		books = append(books, book)
	}

	return books
}

func (st *Storage) GetByID(id string) (entity.Book, bool) {
	st.mu.RLock()
	defer st.mu.RUnlock()

	if book, isExist := st.Books[id]; isExist {
		return book, isExist
	} else {
		return entity.Book{}, isExist
	}
}

func (st *Storage) Update(id string, book entity.Book) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	if _, isExist := st.Books[id]; isExist {
		st.Books[id] = book
		return nil
	} else {
		return errors.New("This book doesn't exist")
	}
}

func (st *Storage) Delete(id string) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	if _, isExist := st.Books[id]; isExist {
		delete(st.Books, id)
		return nil
	} else {
		return errors.New("This book doesn't exist")
	}
}

func NewStorage() *Storage {
	return &Storage{
		Books: make(map[string]entity.Book),
		mu:    &sync.RWMutex{},
	}
}
