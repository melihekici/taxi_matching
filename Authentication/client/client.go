package client

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func PostgresConnection() *sql.DB {
	// postgresInfo := fmt.Sprintf(
	// 	"host =%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5438, "postgres", "postgres", "postgres")

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", "postgres", "postgres", "localhost", 5438, "postgres")

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Panicln(err)
	}

	// err = db.Ping()
	// if err != nil {
	// 	log.Panicln(err)
	// }

	return db
}
