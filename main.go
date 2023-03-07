package main

import (
	"github.com/mariaeduardasouza/golang-api/database"
	"github.com/mariaeduardasouza/golang-api/models"
	"github.com/mariaeduardasouza/golang-api/routes"
)

func main() {
	models.Personalitie =[]models.Personalities {
		{Id:1, Name: "testenome", Story: "testeStory"},
		{Id:2, Name: "testenome2", Story: "testeStory2"},
	}

	database.BankConnection()
	routes.HandleRequest()
}