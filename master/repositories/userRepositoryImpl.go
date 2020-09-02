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

func (u *userRepoImpl) GetAllUser(keyword, page, limit, orderBy, sort string) ([]*models.User, error) {
	// tx, err := u.db.Begin()
	// if err != nil {
	// 	return nil, err
	// }
	queryInput := fmt.Sprintf("SELECT user_name, password, email from m_user where user_name like ? order by %s %s limit %s, %s", orderBy, sort, page, limit)
	// fmt.Println(queryInput)

	// stmt, err := u.db.Prepare(utils.GET_ALL_USER)
	// if err != nil {
	// 	tx.Rollback()
	// 	log.Printf("%v", err)
	// 	return nil, err
	// }
	// defer stmt.Close()

	rows, err := u.db.Query(queryInput, "%"+keyword+"%")
	fmt.Println(err)
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
