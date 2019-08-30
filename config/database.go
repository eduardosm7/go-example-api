package config

import (
	"fmt"
	"os"

	"github.com/eduardosm7/go-example-api/entities/product"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

const connStringTemplate string = `host=%s port=%s dbname=%s user=%s password=%s sslmode=%s`

var db *gorm.DB

func ConnectDB() {

	connString := buildConnString()

	dbConn, err := gorm.Open(os.Getenv("db_type"), connString)

	if err != nil {
		panic(err)
	}

	db = dbConn

	migrate(db)
}

func GetDB() *gorm.DB {
	return db
}

func buildConnString() string {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(
		connStringTemplate,
		os.Getenv("db_host"),
		os.Getenv("db_port"),
		os.Getenv("db_name"),
		os.Getenv("db_user"),
		os.Getenv("db_pass"),
		os.Getenv("db_ssl_mode"),
	)
}

func migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&product.Product{})
}
