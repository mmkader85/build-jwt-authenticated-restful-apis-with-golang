package driver

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	var err error
	pgUrl := os.Getenv("PGSQL_URL")

	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
