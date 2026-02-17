package app

import (
	"fmt"
	"net/http"

	"github.com/barnigator/book-api/internal/delivery"
	"github.com/barnigator/book-api/internal/repository"
	"github.com/barnigator/book-api/internal/usecase"
)

// Run создает объекты и запускает сервер
func Run() {
	repo := repository.NewStorage()
	uc := usecase.NewUseCase(repo)
	handler := delivery.NewHandler(*uc)

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.AddNewBookHandler(w, r)
		case http.MethodGet:
			handler.GetAllBooksHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			handler.UpdateBookHandler(w, r)
		case http.MethodGet:
			handler.GetBookByIdHandler(w, r)
		case http.MethodDelete:
			handler.DeleteBookHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	})

	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Stopped listening: %v\n", err)
	}
}
