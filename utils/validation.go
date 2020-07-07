package utils

var valid bool

func UserValidation(username, password string) bool {
	if username == "" || len(username) > 20 || len(password) > 12 {
		valid = false
	} else {
		valid = true
	}
	return valid
}

func ArticleValidation(articleName string) bool {
	if articleName == "" {
		valid = false
	} else {
		valid = true
	}
	return valid
}
