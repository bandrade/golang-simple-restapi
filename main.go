package main

import (
	"fmt"

	"github.com/bandrade/golang-simple-restapi/database"
	"github.com/bandrade/golang-simple-restapi/models"
	"github.com/bandrade/golang-simple-restapi/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}
	database.ConnectWithDB()
	fmt.Println("Starting rest server with golang")
	routes.HandleResquest()
}
