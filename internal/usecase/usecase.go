package usecase

import (
	"github.com/barnigator/book-api/internal/deps"
	"github.com/barnigator/book-api/internal/entity"
)

type UseCase struct {
	repo deps.Repository
}

func NewUseCase(r deps.Repository) *UseCase { return &UseCase{r} }

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
