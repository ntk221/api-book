package models

import "time"

type Comment struct {
	CommentID uint      `json:"comment_id"`
	ArticleID uint      `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
