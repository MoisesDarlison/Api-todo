package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgres/api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro in decode message: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error on insert %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Operation Success to ID %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
