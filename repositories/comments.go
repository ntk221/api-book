package repositories

import (
	"database/sql"

	"github.com/ntk221/refactor_notion_backend/models"
)

func InserArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlInsertArticle = `
		insert into articles (title, contents, user_name, nice, created_at)
		values (?, ?, ?, ?, ?);
	`
	result, err := db.Exec(sqlInsertArticle, article.Title, article.Contents, article.UserName, article.NiceNum, article.CreatedAt)
	if err != nil {
		return models.Article{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}

	article.ID = uint(id)
	return article, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
	select article_id, title, contents, user_name, nice, created_at
	from articles
	limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, 10, page*10)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
	select *
	from articles
	where article_id = ?;
	`

	row := db.QueryRow(sqlStr, articleID)

	var article models.Article
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func UpdateArticleNice(db *sql.DB, articleID int, nicenum int) error {
	const sqlStr = `
	update articles
	set nice = ?
	where article_id = ?;
	`

	_, err := db.Exec(sqlStr, nicenum, articleID)
	if err != nil {
		return err
	}

	return nil
}

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

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
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
