package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ntk221/refactor_notion_backend/models"
	"github.com/ntk221/refactor_notion_backend/repositories"
)

func SelectCommentList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("got %v, expected %v", num, expectedNum)
	}
}

func InsertCommentList(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "This is my TEST comment.",
	}

	expectedId := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != uint(expectedId) {
		t.Errorf("got %v, expected %v", newComment.CommentID, expectedId)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where contents = ? and username = ?;
		`

		testDB.Exec(sqlStr, comment.Message, comment.CommentID)
	})
}
