package usecases

import (
	"blog-article-go/master/models"
	"blog-article-go/master/repositories"
	"blog-article-go/utils"
	"errors"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type userUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func (u userUsecaseImpl) GetAllUser(keyword, page, limit, orderBy, sort string) ([]*models.User, error) {
	users, err := u.userRepo.GetAllUser(keyword, page, limit, orderBy, sort)
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

func (u userUsecaseImpl) AddUser(username, password, email string, uploadedFile multipart.File, handler *multipart.FileHeader) error {

	// ini Syntax untuk dapetin lokasi project kita
	////////////////////////////////////////////////////////////////////////////////////////
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}
	////////////////////////////////////////////////////////////////////////////////////////

	// ini syntax untuk generate file name nya biar random, filepath.Ext(handler.Filename) itu untuk dapetin extension file nya (.jpg)
	////////////////////////////////////////////////////////////////////////////////////////
	rand.Seed(time.Now().UnixNano())
	min := 11111111111
	max := 99999999999
	fileLocation := filepath.Join(dir, "files", "-"+strconv.Itoa(rand.Intn(max-min+1)+min)+filepath.Ext(handler.Filename))
	////////////////////////////////////////////////////////////////////////////////////////

	log.Println(`FileLocation ->`, fileLocation)

	// Ini syntax untuk bikin temporary file, file name nya pake fileLocation yang dibuat diatas
	////////////////////////////////////////////////////////////////////////////////////////
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer targetFile.Close()
	////////////////////////////////////////////////////////////////////////////////////////

	// INI SYNTAX UNTUK COPY FILE YANG KITA DAPET DARI CONTROLLER KE FILE TEMPORARY YANG BARUSAN KITA BUAT
	////////////////////////////////////////////////////////////////////////////////////////
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		log.Println(`Error While Coping File to Local Storage`, err)
		return err
	}
	////////////////////////////////////////////////////////////////////////////////////////

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
