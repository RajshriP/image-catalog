package api

import (
	"encoding/json"
	"image_catalog/db"
	"net/http"
)

func (s *Store) Images(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithField("api", "images")
	w.Header().Add("Content-Type", "application/json")
	images := make([]db.Image, 0)
	result := s.db.Find(&images)
	if result.Error != nil {
		log.WithField("error", result.Error).Error("fetch images error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while fetching images from the database")
		return
	}
	_ = json.NewEncoder(w).Encode(images)
}
