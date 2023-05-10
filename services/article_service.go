package services

import (
	"fmt"

	"github.com/ntk221/refactor_notion_backend/models"
	"github.com/ntk221/refactor_notion_backend/repositories"
)

func (s *MyAppService) GetArticleService(articleID uint) (models.Article, error) {
	article, err := repositories.GetArticleByID(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, comments...)

	return article, nil

}

func (s *MyAppService) GetArticleListService(page uint) ([]models.Article, error) {
	articles, err := repositories.SelectArticleList(s.db, page)

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	fmt.Println("articles: ", articles)

	return articles, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InserArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}
