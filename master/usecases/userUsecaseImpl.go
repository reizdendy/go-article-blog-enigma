package usecases

import (
	"blog-article-go/master/models"
	"blog-article-go/master/repositories"
	"blog-article-go/utils"
	"errors"
)

type userUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func (u userUsecaseImpl) GetAllUser() ([]*models.User, error) {
	users, err := u.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u userUsecaseImpl) GetUserByID(username string) (*models.User, error) {
	user, err := u.userRepo.GetUserByID(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userUsecaseImpl) AddUser(username, password, email string) error {
	checkValidation := utils.UserValidation(username, password)
	pass := utils.HashPassword(password)
	if checkValidation == true {
		err := u.userRepo.AddUser(username, pass, email)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Check again Username or Password")
	}
	return nil
}

func InitUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepo}
}
