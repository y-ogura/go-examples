package handler

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var conn *sql.DB

// Setup setup handler
func Setup() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	conn = db
	return nil
}

// ConnectDB connect database
func ConnectDB() (*sql.DB, error) {
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	ssl := os.Getenv("SSL_MODE")

	url := "dbname=" + name + "user=" + user
	if pass != "" {
		url = url + "password=" + pass
	}
	url = url + "host=" + host + "sslmode=" + ssl

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
