package main

import (
	//"apiServerBook/auth"
	//"apiServerBook/data"
	//"apiServerBook/rest"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/shuvo-14/api-server/api"
	"github.com/shuvo-14/api-server/auth"
	"github.com/shuvo-14/api-server/db"
	"log"
	"net/http"
)

func main() {
	db.Init()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Post("/login", auth.LogIn)
	r.Post("/logout", auth.LogOut)

	r.Group(func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Get("/", api.GetAllBooks)
			r.Get("/{id}", api.GetOneBook)
			r.Group(func(r chi.Router) {
				// need to add authentication
				r.Use(jwtauth.Verifier(db.TokenAuth))
				r.Use(jwtauth.Authenticator(db.TokenAuth))

				r.Post("/", api.NewBook)
				r.Put("/{id}", api.UpdateBook)
				r.Delete("/{id}", api.DeleteBook)
			})
		})
		r.Route("/authors", func(r chi.Router) {
			r.Get("/", api.GetAllAuthors)
			r.Get("/{id}", api.GetOneAuthor)
		})
		r.Get("/search/{sToken}", api.Search)
	})

	fmt.Println("Listening and Serving to 9090")
	err := http.ListenAndServe("localhost:9090", r)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
