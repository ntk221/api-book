package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ntk221/refactor_notion_backend/handlers"
)

func main() {
	r := chi.NewRouter()

	r.Get("/hello", handlers.HelloHandler)
	r.Post("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.ListArticleHandler)
	r.Get("/article/{id: [0-9]+}", handlers.GetArticleByIDHanlder)

	log.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
