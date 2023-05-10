package services

import "github.com/ntk221/refactor_notion_backend/models"

type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page uint) ([]models.Article, error)
	GetArticleService(articleID uint) (models.Article, error)
}

type CommentServicer interface {
	InsertCommentService(comment models.Comment) (models.Comment, error)
	SelectCommentList(articleID uint) ([]models.Comment, error)
}
