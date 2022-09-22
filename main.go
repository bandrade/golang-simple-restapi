package main

import (
	"fmt"

	"github.com/bandrade/golang-simple-restapi/models"
	"github.com/bandrade/golang-simple-restapi/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Name 1", History: "History 1"},
		{Name: "Name 2", History: "History 2"},
	}

	fmt.Println("Starting rest server with golang")
	routes.HandleResquest()
}
