package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ntk221/refactor_notion_backend/handlers"
)

func main() {

	r := chi.NewRouter()

	r.Get("/hello", handlers.HelloHandler)
	r.Post("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.ListArticleHandler)
	r.Get("/article/{id:[0-9]+}", handlers.GetArticleByIDHanlder)

	r.Get("/test/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Printf("id: %s\n", id)
	})

	log.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
