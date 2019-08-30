package product

import (
	"github.com/eduardosm7/go-example-api/config"
	"github.com/eduardosm7/go-example-api/entities/product"
)

func Create(code string, price uint) product.Product {

	db := config.GetDB()

	newProduct := product.Product{
		Code:  code,
		Price: price,
	}

	db.Create(&newProduct)

	return newProduct
}

func FindAll() []product.Product {

	db := config.GetDB()

	var products []product.Product

	db.Find(&products)

	return products
}

func FindByCode(code string) product.Product {

	db := config.GetDB()

	var product product.Product

	db.First(&product, "code = ?", code)

	return product
}

func Update(code string, price uint) product.Product {

	db := config.GetDB()

	product := FindByCode(code)

	db.Model(&product).Update("Price", price)

	return product
}

func Delete(code string) product.Product {

	db := config.GetDB()

	product := FindByCode(code)

	db.Delete(&product)

	return product
}
