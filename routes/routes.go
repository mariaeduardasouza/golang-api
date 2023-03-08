package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mariaeduardasouza/golang-api/controllers"
)

func HandleRequest() {
	r  := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.AllPersonalities).Methods(("Get"))
	r.HandleFunc("/personalidades/{id}", controllers.ReturnsPersonalities).Methods(("Get"))
	r.HandleFunc("/personalidades", controllers.CreatePersonality).Methods(("Post"))
	r.HandleFunc("/personalidades/{id}", controllers.DeletePrsonality).Methods(("Delete"))
	r.HandleFunc("/personalidades/{id}", controllers.EditPersonality).Methods(("Put"))
	log.Fatal(http.ListenAndServe(":8000", r))
}
