package services

import (
	"auth/client"
	"auth/models"
	"log"

	_ "github.com/lib/pq"
)

func GetUserObject(email string) (models.User, bool) {
	db := client.PostgresConnection()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM usertable WHERE email=$1", email)

	var user models.User
	var id int

	err := row.Scan(&id, &user.Email, &user.Username, &user.Passwordhash)
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
