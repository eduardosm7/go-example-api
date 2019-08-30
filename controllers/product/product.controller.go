package product

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	productEntity "github.com/eduardosm7/go-example-api/entities/product"
	productService "github.com/eduardosm7/go-example-api/services/product"
	"github.com/eduardosm7/go-example-api/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product productEntity.Product

	_ = json.NewDecoder(r.Body).Decode(&product)

	newProduct := productService.Create(
		product.Code,
		product.Price,
	)

	utils.Respond(w, newProduct)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	var products []productEntity.Product

	products = productService.FindAll()

	utils.Respond(w, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	product := productService.FindByCode(params["id"])

	utils.Respond(w, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var product productEntity.Product

	_ = json.NewDecoder(r.Body).Decode(&product)

	updatedProduct := productService.Update(
		product.Code,
		product.Price,
	)

	utils.Respond(w, updatedProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	product := productService.FindByCode(params["id"])

	deletedProduct := productService.Delete(product.Code)

	utils.Respond(w, deletedProduct)
}
