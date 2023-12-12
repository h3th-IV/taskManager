package database

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	dB *sql.DB
)

func InitDB() error {
	dataSrcName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"))

	db, err := sql.Open("postgres", dataSrcName)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Connection to Database was succesfull")
	return nil
}

func CloseDB() {
	dB.Close()
}
