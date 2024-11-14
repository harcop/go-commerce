package main

import (
	"ecommerce-backend/database"
	"ecommerce-backend/models"
	"ecommerce-backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	r := mux.NewRouter()
	routes.AuthRoutes(r)
	routes.ProductRoutes(r)

	log.Println("Server started")
	http.ListenAndServe(":8999", r)
}

func init() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Product{})
}
