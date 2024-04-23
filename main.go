package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shuvo-14/api-server/api"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK\n"))
		if err != nil {
			fmt.Println(err)
		}
	})
	r.Mount("/books", BookRoutes())
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println(err)
	}
}

func BookRoutes() chi.Router {
	r := chi.NewRouter()
	bookHandler := api.BookHandler{}
	r.Get("/", bookHandler.ListBooks)
	r.Post("/", bookHandler.CreateBook)
	r.Get("/{id}", bookHandler.GetBooks)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)

	return r
}
