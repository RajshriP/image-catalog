package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type Store struct {
	db  *gorm.DB
	log *logrus.Logger
}

func New(db *gorm.DB, log *logrus.Logger) *Store {
	s := &Store{db: db, log: log}
	http.HandleFunc("/api/upload/", s.Upload)
	http.HandleFunc("/api/images/", s.Images)
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
