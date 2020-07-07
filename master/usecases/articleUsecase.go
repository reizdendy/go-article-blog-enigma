package usecases

import "blog-article-go/master/models"

type ArticleUsecase interface {
	GetAllArticle(username string) ([]*models.Article, error)
	GetArticleByID(username, articleID string) (*models.Article, error)
	AddArticle(articleName, articleContent, articleUser string) error
	UpdateArticle(articleID, articleName, articleContent, articleUser string) error
	DeleteArticle(username, articleID string) error
}
