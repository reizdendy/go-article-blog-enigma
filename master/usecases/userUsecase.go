package usecases

import "blog-article-go/master/models"

type UserUsecase interface {
	GetAllUser() ([]*models.User, error)
	GetUserByID(username string) (*models.User, error)
	AddUser(username, password, email string) error
	// UpdateUser(UserID, UserName string) error
	// DeleteUser(UserID string) error
}
