package main

import "github.com/go-examples/gorm/db"

func main() {
	DB := db.ConnectDB()
	db.SetMaxConnections(3)
	defer DB.Close()
}
