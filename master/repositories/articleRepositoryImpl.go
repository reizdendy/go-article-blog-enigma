package repositories

import (
	"blog-article-go/master/models"
	"blog-article-go/utils"
	"database/sql"
	"fmt"
	"log"
	"time"

	guuid "github.com/google/uuid"
)

type articleRepoImpl struct {
	db *sql.DB
}

func (a *articleRepoImpl) GetAllArticle(username string) ([]*models.Article, error) {
	rows, err := a.db.Query(utils.GET_ALL_ARTICLE, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*models.Article
	for rows.Next() {
		article := models.Article{}
		err := rows.Scan(&article.ArticleID, &article.ArticleName, &article.ArticleContent, &article.UserCreator, &article.CreatedAt)

		if err != nil {
			return nil, err
		}

		articles = append(articles, &article)
	}
	fmt.Println(articles)
	return articles, nil
}

func (a *articleRepoImpl) GetArticleByID(username, articleID string) (*models.Article, error) {
	article := models.Article{}
	err := a.db.QueryRow(utils.GET_ARTICLE_BY_ID, username, articleID).Scan(&article.ArticleID, &article.ArticleName, &article.ArticleContent, &article.UserCreator, &article.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *articleRepoImpl) AddArticle(articleName, articleContent, articleUser string) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := a.db.Prepare(utils.ADD_ARTICLE)
	if err != nil {
		tx.Rollback()
		log.Printf("%v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(guuid.New(), articleName, articleContent, articleUser, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		tx.Rollback()
		log.Printf("%v", err)
		return err
	}
	return tx.Commit()
}

// UpdateArticle is a function to update Article data
func (a *articleRepoImpl) UpdateArticle(articleID, articleName, articleContent, articleUser string) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := a.db.Prepare(utils.UPDATE_ARTICLE)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(articleName, articleContent, articleUser, articleID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}

// DeleteArticle is a function to delete Article data
func (a *articleRepoImpl) DeleteArticle(username, articleID string) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := a.db.Prepare(utils.DELETE_ARTICLE)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, articleID)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}

	return tx.Commit()
}

func InitArticleRepositoryImpl(db *sql.DB) ArticleRepository {
	return &articleRepoImpl{db}
}
