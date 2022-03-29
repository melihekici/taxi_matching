package client

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func PostgresConnection() *sql.DB {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", "postgres", "postgres", os.Getenv("POSTGRES_HOST"), 5432, "postgres")

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Panicln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicln("Unable to connect to postgres database " + err.Error())
	}

	log.Println("Connected to Postgres Database")
	return db
}
