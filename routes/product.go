package routes

import (
	"ecommerce-backend/controllers"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/proudcts", controllers.CreateProduct).Methods("POSt")
}
