package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MainHandler(st *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
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

			var book Book
			err = json.Unmarshal(body, &book)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			err = st.Create(book)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusOK)
			return

		case http.MethodGet:
			books := st.GetAll()

			jsonBooks, err := json.Marshal(books)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBooks)
			return

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

func BookHandler(st *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		IdBook := r.URL.Path[len(r.URL.Path)-2:]

		switch r.Method {
		case http.MethodGet:
			book, ok := st.GetByID(IdBook)
			if ok {
				jsonBook, err := json.Marshal(book)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(jsonBook)
				return
			} else {
				w.WriteHeader(http.StatusNotFound)
				return
			}

		case http.MethodPut:
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

			var book Book
			err = json.Unmarshal(body, &book)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			err = st.Update(IdBook, book)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusOK)
			return

		case http.MethodDelete:
			err := st.Delete(IdBook)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusOK)
			return

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}
