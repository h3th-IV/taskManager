package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Database struct {
	DB *sql.DB
}

func NewDB() *Database {
	dataSrcName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"))

	db, err := sql.Open("postgres", dataSrcName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to Database was succesfull")

	return &Database{
		DB: db,
	}
}

func InitDB() {
	dataBase := NewDB()
	defer dataBase.DB.Close()

}
