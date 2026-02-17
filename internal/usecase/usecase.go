package usecase

import "github.com/barnigator/book-api/internal/entity"

type Repository interface {
	Create(book entity.Book) error
	GetAll() []entity.Book
	GetByID(id string) (entity.Book, bool)
	Update(id string, book entity.Book) error
	Delete(id string) error
}

type UseCase struct {
	repo Repository
}

func NewUseCase(r Repository) *UseCase { return &UseCase{r} }

func (uc *UseCase) CreateBook(book entity.Book) error {
	return uc.repo.Create(book)
}

func (uc *UseCase) GetAllBooks() []entity.Book {
	return uc.repo.GetAll()
}

func (uc *UseCase) GetBookById(id string) (entity.Book, bool) {
	return uc.repo.GetByID(id)
}

func (uc *UseCase) UpdateBook(id string, book entity.Book) error {
	return uc.repo.Update(id, book)
}

func (uc *UseCase) Delete(id string) error {
	return uc.repo.Delete(id)
}
