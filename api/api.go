package api

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type Store struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Store {
	s := &Store{DB: db}
	http.HandleFunc("/api/upload/", s.Upload)
	http.Handle("/images/", http.FileServer(http.Dir(".")))
	return s
}

func (s *Store) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}

func errMessage(w http.ResponseWriter, msg string) error {
	type errorMessage struct {
		Error string `json:"error"`
	}
	return json.NewEncoder(w).Encode(errorMessage{Error: msg})
}
