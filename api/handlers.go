package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/shuvo-14/api-server/db"
	"net/http"
)

type BookHandler struct {
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(db.ListBooks())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book := getBook(id)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book db.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	isPresent := bookExists(book)

	if isPresent != nil {
		http.Error(w, "This ID already Exist", http.StatusBadRequest)
		return
	}
	db.Books = append(db.Books, &book)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book db.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBook := updateBook(id, book)

	if updatedBook == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book := deleteBook(id)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getBook(id string) *db.Book {
	for _, book := range db.Books {
		if book.ID == id {
			return book
		}
	}
	return nil
}

func bookExists(book db.Book) *db.Book {
	for _, bookI := range db.Books {
		if bookI.ID == book.ID {
			return &book
		}
	}

	return nil
}

func deleteBook(id string) *db.Book {
	for i, book := range db.Books {
		if book.ID == id {
			db.Books = append(db.Books[:i], db.Books[i+1:]...)
			return &db.Book{}
		}
	}
	return nil
}

func updateBook(id string, bookUpdate db.Book) *db.Book {
	for i, book := range db.Books {
		if book.ID == id {
			db.Books[i] = &bookUpdate
			return book
		}
	}
	return nil
}
