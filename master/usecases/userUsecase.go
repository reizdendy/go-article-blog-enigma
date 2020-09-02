package usecases

import (
	"blog-article-go/master/models"
	"mime/multipart"
)

type UserUsecase interface {
	GetAllUser(keyword, page, limit, orderBy, sort string) ([]*models.User, error)
	GetUserByID(username string) (*models.User, error)
	AddUser(username, password, email string, uploadedFile multipart.File, handler *multipart.FileHeader) error
	// UpdateUser(UserID, UserName string) error
	// DeleteUser(UserID string) error
}
