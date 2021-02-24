package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shawnzxx/go/inventoryservice/database"
	"github.com/shawnzxx/go/inventoryservice/product"
	"github.com/shawnzxx/go/inventoryservice/receipt"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	product.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
