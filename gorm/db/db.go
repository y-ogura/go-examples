package db

import (
	"os"

	"github.com/jinzhu/gorm"
	// import postgres package
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB Connect Database instance
type DB *gorm.DB

// ConnectDB connection DB
func ConnectDB() *gorm.DB {
	url := os.Getenv("DATABASE_URL") + os.Getenv("SSL_MODE")
	DB, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	return DB
}

// SetMaxConnections setting max open connections.
func (d DB) SetMaxConnections(i int) {
	d.DB().SetMaxOpenConns(i)
	return
}
