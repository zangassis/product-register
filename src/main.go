package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/assiszang/product-register/pkg/database"
	"github.com/assiszang/product-register/pkg/database/api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./products.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Database does not exist or connection failed:", err)
		return
	}

	query := "SELECT 1 FROM products LIMIT 1"
	_, err = db.Exec(query)
	if err != nil {
		database.CreateTable(db)
	}

	api.Execute(db)
}
