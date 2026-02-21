package deps

import "github.com/barnigator/book-api/internal/entity"

type Repository interface {
	Create(book entity.Book) error
	GetAll() []entity.Book
	GetByID(id string) (entity.Book, bool)
	Update(id string, book entity.Book) error
	Delete(id string) error
}

type UseCase interface {
	CreateBook(book entity.Book) error
	GetAllBooks() []entity.Book
	GetBookById(id string) (entity.Book, bool)
	UpdateBook(id string, book entity.Book) error
	Delete(id string) error
}
