package controllers

import (
	"encoding/json"
	"net/http"

	"ecommerce-backend/database"
	"ecommerce-backend/models"
)

// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	var product models.Product
// 	var input int

// 	json.NewEncoder(r.Query).Decode(input)

// }

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}
