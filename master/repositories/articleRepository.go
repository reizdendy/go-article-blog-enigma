package repositories

import "blog-article-go/master/models"

type ArticleRepository interface {
	GetAllArticle(username string) ([]*models.Article, error)
	GetArticleByID(username, articleID string) (*models.Article, error)
	AddArticle(articleName, articleContent, username string) error
	UpdateArticle(articleID, articleName, articleContent, username string) error
	DeleteArticle(username, articleID string) error
}
