package services

import (
	"github.com/ntk221/refactor_notion_backend/models"
	"github.com/ntk221/refactor_notion_backend/repositories"
)

func InsertCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}

func SelectCommentList(articleID uint) ([]models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	comments, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
