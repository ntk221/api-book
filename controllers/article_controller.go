package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/ntk221/refactor_notion_backend/controllers/services"
	"github.com/ntk221/refactor_notion_backend/models"
)

type MyAppController struct {
	service services.ArticleServicer
}

func NewMyAppController(s services.ArticleServicer) *MyAppController {
	return &MyAppController{service: s}
}

func (c *MyAppController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "cannot decode json\n", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "cannot post article\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *MyAppController) ListArticleHandler(w http.ResponseWriter, req *http.Request) {
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

	articles, err := c.service.GetArticleListService(uint(page))
	if err != nil {
		http.Error(w, "Cannot Get Article List", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articles)

}

func (c *MyAppController) GetArticleByIDHanlder(w http.ResponseWriter, req *http.Request) {
	idStr := chi.URLParam(req, "id")
	articleID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Article ID", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(uint(articleID))
	if err != nil {
		http.Error(w, "Cannot Get Article", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}
