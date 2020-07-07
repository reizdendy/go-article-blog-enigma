package usecases

import (
	"blog-article-go/master/models"
	"blog-article-go/master/repositories"
	"blog-article-go/utils"
	"errors"
)

type articleUsecaseImpl struct {
	articleRepo repositories.ArticleRepository
}

func (a articleUsecaseImpl) GetAllArticle(username string) ([]*models.Article, error) {
	articles, err := a.articleRepo.GetAllArticle(username)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a articleUsecaseImpl) GetArticleByID(username, articleID string) (*models.Article, error) {
	article, err := a.articleRepo.GetArticleByID(username, articleID)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a articleUsecaseImpl) AddArticle(articleName, articleContent, articleUser string) error {
	checkValidation := utils.ArticleValidation(articleName)
	if checkValidation == true {
		err := a.articleRepo.AddArticle(articleName, articleContent, articleUser)
		if err != nil {
			return err
		}
	} else {
		return errors.New("data is empty")
	}
	return nil
}

func (a articleUsecaseImpl) UpdateArticle(articleID, articleName, articleContent, articleUser string) error {
	err := a.articleRepo.UpdateArticle(articleID, articleName, articleContent, articleUser)
	if err != nil {
		return err
	}
	return nil
}

func (a articleUsecaseImpl) DeleteArticle(username, articleID string) error {
	err := a.articleRepo.DeleteArticle(username, articleID)
	if err != nil {
		return err
	}
	return nil
}

func InitArticleUsecase(articleRepo repositories.ArticleRepository) ArticleUsecase {
	return &articleUsecaseImpl{articleRepo}
}
