package routes

import (
	"log"
	"net/http"

	"github.com/bandrade/golang-simple-restapi/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
