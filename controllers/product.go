package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce-backend/database"
	"ecommerce-backend/models"

	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	var productID string

	vars := mux.Vars(r)
	if id, exists := vars["id"]; exists {
		productID = id
	} else {
		http.Error(w, "product id is requiresd", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(productID)

	if err := database.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

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
