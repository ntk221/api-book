package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var HelloHandler = func(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

var PostArticleHandler = func(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
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

	resString := fmt.Sprintf("Listing Article on Page: %d\n", page)
	io.WriteString(w, resString)
}

func GetArticleByIDHanlder(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(w, "Invalid Article ID", http.StatusBadRequest)
		return
	}
	res := fmt.Sprintf("Getting Article with ID: %d\n", articleID)
	io.WriteString(w, res)
}
