package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ntk221/refactor_notion_backend/models"
	"github.com/ntk221/refactor_notion_backend/services"
)

var HelloHandler = func(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

var PostArticleHandler = func(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "cannot decode json\n", http.StatusBadRequest)
		return
	}

	article, err := services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "cannot post article\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func ListArticleHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid Page Number", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articles, err := services.GetArticleListService(uint(page))
	if err != nil {
		http.Error(w, "Cannot Get Article List", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articles)

}

func GetArticleByIDHanlder(w http.ResponseWriter, req *http.Request) {
	idStr := chi.URLParam(req, "id")
	articleID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Article ID", http.StatusBadRequest)
		return
	}

	article, err := services.GetArticleService(uint(articleID))
	if err != nil {
		http.Error(w, "Cannot Get Article", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}
