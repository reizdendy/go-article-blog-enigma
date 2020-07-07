package controllers

import (
	"blog-article-go/master/models"
	"blog-article-go/master/usecases"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type userHandler struct {
	userUsecase usecases.UserUsecase
}

var (
	user         models.User
	userResponse models.Response
)

func getMuxUser(value string, r *http.Request) string {
	parameter := mux.Vars(r)
	return parameter[value]
}

func UserController(r *mux.Router, service usecases.UserUsecase) {
	userHandler := userHandler{service}
	r.HandleFunc("/users", userHandler.ListUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHandler.ListUserById).Methods(http.MethodGet)
	r.HandleFunc("/user", userHandler.AddUserPage).Methods(http.MethodPost)
}

func (u *userHandler) ListUser(w http.ResponseWriter, r *http.Request) {
	users, err := u.userUsecase.GetAllUser()

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		userResponse.Messages = "List of user Data"
		userResponse.Data = users
		userResponse.Status = http.StatusOK
		byteOfuser, err := json.Marshal(userResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfuser)
	}
}

func (u *userHandler) ListUserById(w http.ResponseWriter, r *http.Request) {
	iduser := getMuxUser("id", r)
	user, err := u.userUsecase.GetUserByID(iduser)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		userResponse.Messages = "data user berhasil diambil"
		userResponse.Data = user
		userResponse.Status = http.StatusOK
		byteOfuser, err := json.Marshal(userResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfuser)
	}
}

//AdduserPage is a controller of AddAuser service (Insert / POST). It got the json data from Request Body RAW
func (u *userHandler) AddUserPage(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := u.userUsecase.AddUser(user.UserName, user.Password, user.Email)
	if err != nil {
		w.Write([]byte("Check again Username or Password. Username is max 20 characters and Password is max 12 characters"))
	} else {
		userResponse.Messages = "Insert user Data Success!"
		userResponse.Status = http.StatusOK
		userResponse.Data = ""
		byteOfuser, err := json.Marshal(userResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfuser)
	}
}
