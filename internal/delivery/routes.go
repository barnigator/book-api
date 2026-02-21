package delivery

import "net/http"

func RegisterRoutes(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("/books", h.GetAllBooks)
	mux.HandleFunc("/books/add", h.AddNewBook)
	mux.HandleFunc("/books/search", h.GetBookById)
	mux.HandleFunc("/books/update", h.UpdateBook)
	mux.HandleFunc("/books/delete", h.DeleteBook)
}
