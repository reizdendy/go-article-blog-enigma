package repositories

import "blog-article-go/master/models"

type UserRepository interface {
	GetAllUser() ([]*models.User, error)
	GetUserByID(username string) (*models.User, error)
	AddUser(username, password, email string) error
	// UpdateUser(UserID, UserName string) error
	// DeleteUser(UserID string) error
}
