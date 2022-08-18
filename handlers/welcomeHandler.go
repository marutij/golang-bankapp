package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marutijogdand17/golang-bankapp/models"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&models.Message{Data: "Welcome to maruti bank"})
}
