package services

import (
	"fmt"

	"github.com/ntk221/refactor_notion_backend/models"
	"github.com/ntk221/refactor_notion_backend/repositories"
)

func GetArticleService(articleID uint) (models.Article, error) {
	// TODO: sql.DB 型を受け取るようにする
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.GetArticleByID(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, comments...)

	return article, nil

}

func GetArticleListService(page uint) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fmt.Println("page: ", page)

	articles, err := repositories.SelectArticleList(db, page)

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	fmt.Println("articles: ", articles)

	return articles, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InserArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}
