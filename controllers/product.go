package controllers

import (
	"encoding/json"
	"net/http"

	"ecommerce-backend/database"
	"ecommerce-backend/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	var response struct {
		Data  []models.Product `json:"data"`
		Message string 			`json:"message"`
	}
	database.DB.Find(&products)
	response.Data = products
	response.Message = "OK Done"
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}
