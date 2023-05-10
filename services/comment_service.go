package services

import (
	"github.com/ntk221/refactor_notion_backend/models"
	"github.com/ntk221/refactor_notion_backend/repositories"
)

func (s *MyAppService) InsertCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}

func (s *MyAppService) SelectCommentList(articleID uint) ([]models.Comment, error) {
	comments, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
