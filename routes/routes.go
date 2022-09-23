package routes

import (
	"log"
	"net/http"

	"github.com/bandrade/golang-simple-restapi/controllers"
	"github.com/bandrade/golang-simple-restapi/middleware"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities/", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.FindPersonalityById).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
	r.HandleFunc("/api/personalities/", controllers.CreatePersonality).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
