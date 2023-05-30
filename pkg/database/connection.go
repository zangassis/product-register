package database

import (
	"database/sql"

	"github.com/assiszang/product-register/internal/structs"
)

func InitializeDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			price REAL
		);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetProduct(db *sql.DB, id int) (*structs.Product, error) {
	selectSQL := `
		SELECT id, name, price
		FROM products
		WHERE id = ?
	`

	row := db.QueryRow(selectSQL, id)

	product := &structs.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func GetAllProducts(db *sql.DB) ([]*structs.Product, error) {
	selectSQL := `
		SELECT id, name, price
		FROM products
	`

	rows, err := db.Query(selectSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*structs.Product{}
	for rows.Next() {
		product := &structs.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func CreateProduct(db *sql.DB, product *structs.Product) error {
	insertSQL := `
		INSERT INTO products (name, price)
		VALUES (?, ?)
	`

	_, err := db.Exec(insertSQL, product.Name, product.Price)
	return err
}

func UpdateProduct(db *sql.DB, product *structs.Product) error {
	updateSQL := `
		UPDATE products
		SET name = ?, price = ?
		WHERE id = ?
	`

	_, err := db.Exec(updateSQL, product.Name, product.Price, product.ID)
	return err
}

func DeleteProduct(db *sql.DB, id int) error {
	deleteSQL := `
		DELETE FROM products
		WHERE id = ?
	`

	_, err := db.Exec(deleteSQL, id)
	return err
}
