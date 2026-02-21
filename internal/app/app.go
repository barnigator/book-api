package app

import (
	"github.com/barnigator/book-api/config"
	"github.com/barnigator/book-api/internal/delivery"
	"github.com/barnigator/book-api/internal/repository"
	"github.com/barnigator/book-api/internal/usecase"
)

// Run создает объекты и запускает сервер
func Run() {
	repo := repository.NewStorage()
	uc := usecase.NewUseCase(repo)
	handler := delivery.NewHandler(uc)
	cfg := config.NewConfig()
	srv := NewServer(cfg, handler)

	StartServer(srv)
	StopServer(srv)
}
