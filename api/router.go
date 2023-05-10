package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ntk221/refactor_notion_backend/controllers"
	"github.com/ntk221/refactor_notion_backend/services"
)

func NewRouter(db *sql.DB) *chi.Mux {
	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := chi.NewRouter()

	r.Get("/hello", con.HelloHandler)
	r.Post("/article", con.PostArticleHandler)
	r.Get("/article/list", con.ListArticleHandler)
	r.Get("/article/{id:[0-9]+}", con.GetArticleByIDHanlder)

	r.Get("/test/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Printf("id: %s\n", id)
	})

	return r
}
