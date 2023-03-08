package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mariaeduardasouza/golang-api/database"
	"github.com/mariaeduardasouza/golang-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request)  {
	var list []models.Personalities
	database.DB.Find(&list)
	json.NewEncoder(w).Encode(list)
}

func ReturnsPersonalities(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	var listPersonalities models.Personalities

	database.DB.First(&listPersonalities, id)
	json.NewEncoder(w).Encode(listPersonalities)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request)  {
	var newPersonalitie models.Personalities
	json.NewDecoder(r.Body).Decode(&newPersonalitie)
	database.DB.Create(&newPersonalitie)
	json.NewEncoder(w).Encode(newPersonalitie)
}

func DeletePrsonality(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	var listPersonality models.Personalities

	database.DB.Delete(&listPersonality, id)
	json.NewEncoder(w).Encode(listPersonality)
}