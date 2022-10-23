package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"postgres/api/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error in gel all todos: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
