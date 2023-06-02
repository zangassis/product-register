package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/assiszang/product-register/pkg/database"

	"github.com/assiszang/product-register/internal/structs"
	"github.com/gorilla/mux"
)

var products []*structs.Product
var dbS *sql.DB

func Execute(db *sql.DB) {
	dbS = db
	router := mux.NewRouter()
	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	getProducts, err := database.GetAllProducts(db)
	products = getProducts

	log.Fatal(http.ListenAndServe(":8000", router))
	log.Fatal(err)
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.NotFound(w, r)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product *structs.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	products = append(products, product)
	database.CreateProduct(dbS, product)
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedProduct *structs.Product
	_ = json.NewDecoder(r.Body).Decode(&updatedProduct)

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			updatedProduct.ID = product.ID
			json.NewEncoder(w).Encode(updatedProduct)
			database.UpdateProduct(dbS, product)
			return
		}
	}

	http.NotFound(w, r)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			database.DeleteProduct(dbS, id)
			return
		}
	}

	http.NotFound(w, r)
}
