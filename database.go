package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Data Source Name (DSN) untuk koneksi MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	var errDB error
	db, errDB = sql.Open("mysql", dsn)
	if errDB != nil {
		log.Fatal(errDB)
	}

	errDB = db.Ping()
	if errDB != nil {
		log.Fatal(errDB)
	}

	fmt.Println("Database connected successfully!")
}
