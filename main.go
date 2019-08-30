package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/eduardosm7/go-example-api/config"
	"github.com/eduardosm7/go-example-api/routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	config.ConnectDB()

	router := routes.GetRouter()

	port := os.Getenv("port")
	if port == "" {
		port = "5000"
	}

	fmt.Println("Server listening on port " + port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}
}
