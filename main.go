package main

import (
	"fmt"

	"github.com/bandrade/golang-simple-restapi/database"
	"github.com/bandrade/golang-simple-restapi/routes"
)

func main() {
	database.ConnectWithDB()
	fmt.Println("Starting rest server with golang")
	routes.HandleResquest()
}
