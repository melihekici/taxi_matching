package services

import (
	"auth/client"
	"auth/models"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func GetUserObject(email string) (models.User, bool) {
	db := client.PostgresConnection()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM usertable WHERE email=$1", email)

	var user models.User
	var id int

	err := row.Scan(&id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		log.Println("Error executing sql", err.Error())
		return models.User{}, false
	}

	return user, true
}

func AddUserObject(email string, username string, passwordhash string) bool {
	db := client.PostgresConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO usertable (email, username, pass) VALUES($1, $2, $3)")
	if err != nil {
		log.Println("Error preparing statement", err)
		return false
	}

	_, err = stmt.Exec(email, username, passwordhash)
	if err != nil {
		log.Println("Error while inserting user to the database")
		return false
	}

	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ValidatePasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
