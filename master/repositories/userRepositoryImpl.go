package repositories

import (
	"blog-article-go/master/models"
	"blog-article-go/utils"
	"database/sql"
	"fmt"
	"log"
)

type userRepoImpl struct {
	db *sql.DB
}

func (u *userRepoImpl) GetAllUser() ([]*models.User, error) {
	rows, err := u.db.Query(utils.GET_ALL_USER)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.UserName, &user.Password, &user.Email)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	fmt.Println(users)
	return users, nil
}

func (u *userRepoImpl) GetUserByID(id string) (*models.User, error) {
	user := models.User{}
	err := u.db.QueryRow(utils.GET_USER_BY_ID, id).Scan(&user.UserName, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepoImpl) AddUser(username, password, email string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := u.db.Prepare(utils.ADD_USER)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password, email)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}

func InitUserRepositoryImpl(db *sql.DB) UserRepository {
	return &userRepoImpl{db}
}
