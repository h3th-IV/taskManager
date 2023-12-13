package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	dataSrcName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"))

	log.Printf("Data Source Name: %s", dataSrcName) //dBugging

	var err error
	DB, err = sql.Open("postgres", dataSrcName)
	if err != nil {
		log.Fatal(err)
	}
	return DB.Ping()
}

func CloseDB() {
	DB.Close()
}
