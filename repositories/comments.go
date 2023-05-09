package repositories

import (
	"database/sql"

	"github.com/ntk221/refactor_notion_backend/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
	insert into comments (article_id, message, created_at)
	values (?, ?, now());
	`

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message, comment.CreatedAt)
	if err != nil {
		return models.Comment{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	comment.CommentID = uint(id)
	return comment, nil
}

func SelectCommentList(db *sql.DB, articleID uint) ([]models.Comment, error) {
	const sqlStr = `
	select *
	from comments
	where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
