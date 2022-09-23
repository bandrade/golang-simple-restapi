package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bandrade/golang-simple-restapi/database"
	"github.com/bandrade/golang-simple-restapi/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func FindPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	var p models.Personality
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var p models.Personality
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality
	database.DB.First(&p, id)
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)

}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var p models.Personality
	id := vars["id"]
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
