package repositories_test

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ntk221/refactor_notion_backend/models"
	"github.com/ntk221/refactor_notion_backend/repositories"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest",
			expected: models.Article{
				ID:        1,
				Title:     "Hello World",
				Contents:  "This is my first article.",
				UserName:  "John Doe",
				NiceNum:   0,
				CreatedAt: time.Now(),
			},
		},
		{
			testTitle: "subtest2",
			expected: models.Article{
				ID:        2,
				Title:     "Hello World 2",
				Contents:  "This is my second article.",
				UserName:  "John Doe",
				NiceNum:   0,
				CreatedAt: time.Now(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("got %v, expected %v", got, test.expected)
			}

			if got.Title != test.expected.Title {
				t.Errorf("got %v, expected %v", got, test.expected)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("got %v, expected %v", got, test.expected)
			}

			if got.UserName != test.expected.UserName {
				t.Errorf("got %v, expected %v", got, test.expected)
			}

			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("got %v, expected %v", got, test.expected)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("got %v, expected %v", num, expectedNum)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "Hello World",
		Contents: "This is my TEST article.",
		UserName: "John Doe",
	}

	expectedId := 3
	newArticle, err := repositories.InserArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != uint(expectedId) {
		t.Errorf("got %v, expected %v", newArticle.ID, expectedId)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?;
		`

		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateArticleNice(t *testing.T) {
	var articleID uint

	articleID = 1

	before, _ := repositories.SelectArticleDetail(testDB, articleID)

	err := repositories.UpdateArticleNice(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	got, _ := repositories.SelectArticleDetail(testDB, articleID)

	if got.NiceNum != before.NiceNum+1 {
		t.Errorf("got %v, expected %v", got.NiceNum, before.NiceNum)
	}
}
