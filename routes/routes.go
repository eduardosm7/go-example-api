package routes

import (
	"github.com/eduardosm7/go-example-api/controllers/product"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {

	router := mux.NewRouter()

	// Products
	router.HandleFunc("/products", product.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", product.GetProduct).Methods("GET")
	router.HandleFunc("/products", product.CreateProduct).Methods("POST")
	router.HandleFunc("/products", product.UpdateProduct).Methods("PUt")
	router.HandleFunc("/products/{id}", product.DeleteProduct).Methods("DELETE")

	return router
}
