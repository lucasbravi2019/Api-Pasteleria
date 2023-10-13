package core

import (
	"log"
	"os"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var databaseConnection *sql.DB

func GetDatabaseConnection() *sql.DB {
	if databaseConnection != nil {
		return databaseConnection
	}

	connString := os.Getenv("CONNECTION_STRING")

	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		log.Fatal(err)
	}

	databaseConnection = db

	if err := CheckDatabaseHealth(); err != nil {
		log.Fatal(err)
	}

	return databaseConnection
}

func CheckDatabaseHealth() error {
	log.Println("Checking database health")
	return GetDatabaseConnection().Ping()
}
