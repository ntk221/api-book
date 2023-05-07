package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/ntk221/refactor_notion_backend/models"
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

	article := reqArticle
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

	fmt.Println("page: ", page)

	articles := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articles)

}

func GetArticleByIDHanlder(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(w, "Invalid Article ID", http.StatusBadRequest)
		return
	}
	res := fmt.Sprintf("Getting Article with ID: %d\n", articleID)
	fmt.Println(res)
	artile := models.Article1
	json.NewEncoder(w).Encode(artile)
}
