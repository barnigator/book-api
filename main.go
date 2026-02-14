package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	st := NewStorage()

	mux.HandleFunc("/books", MainHandler(st))
	mux.HandleFunc("/books/", BookHandler(st))

	http.ListenAndServe(":8080", mux)
}
