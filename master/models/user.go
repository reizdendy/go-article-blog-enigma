package models

// User is a struct of User Data
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
