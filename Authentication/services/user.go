package services

import "auth/models"

var userList = []models.User{
	{"melih.ekici4@gmail.com", "mekici", "asdqwe", "melih ekici", "", 1},
	{"melih.ekici5@gmail.com", "mekici5", "asdqwe5", "melih ekici5", "", 0},
	{"melih.ekici6@gmail.com", "mekici6", "asdqwe6", "melih ekici6", "", 0},
}

func GetUserObject(email string) (models.User, bool) {
	// needs to be replaced using Database
	for _, user := range userList {
		if user.Email == email {
			return user, true
		}
	}
	return models.User{}, false
}

func AddUserObject(email string, username string, passwordhash string, fullname string, role int) bool {
	newUser := models.User{
		Email:        email,
		Passwordhash: passwordhash,
		Username:     username,
		Fullname:     fullname,
		Role:         role,
	}

	// check if user allready exists
	for _, u := range userList {
		if u.Equals(newUser) {
			return false
		}
	}
	return true
}
