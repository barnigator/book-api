package delivery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/barnigator/book-api/internal/entity"
	"github.com/barnigator/book-api/internal/usecase"
)

type Handler struct {
	uc usecase.UseCase
}

func NewHandler(uc usecase.UseCase) *Handler {
	return &Handler{uc}
}

func (h *Handler) AddNewBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Content-type must be json")
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var book entity.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.uc.CreateBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *Handler) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := h.uc.GetAllBooks()

	jsonBooks, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBooks)
}

func (h *Handler) GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong URL format"))
		return
	}

	idBook := pathParts[2]

	book, ok := h.uc.GetBookById(idBook)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book doesn't exist"))
		return
	}
	jsonBook, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBook)
}

func (h *Handler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Content-type must be json")
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var book entity.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong URL format"))
		return
	}

	idBook := pathParts[2]

	err = h.uc.UpdateBook(idBook, book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *Handler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong URL format"))
		return
	}

	idBook := pathParts[2]

	err := h.uc.Delete(idBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
