package routes

import (
	"ecommerce-backend/controllers"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")
}
