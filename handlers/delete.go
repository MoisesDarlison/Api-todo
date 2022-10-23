package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgres/api/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error in get ID params: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error in delete request: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error in update more one data %d", err)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Delete Success to ID %v", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
