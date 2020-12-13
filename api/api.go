package api

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type Store struct {
	DB *gorm.DB
}

func errMessage(w http.ResponseWriter, msg string) error {
	type errorMessage struct {
		Error string `json:"error"`
	}
	return json.NewEncoder(w).Encode(errorMessage{Error: msg})
}
