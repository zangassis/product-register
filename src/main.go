package main

import (
	"log"

	"github.com/assiszang/product-register/pkg/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
