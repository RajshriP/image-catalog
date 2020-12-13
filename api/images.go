package api

import (
	"encoding/json"
	"image_catalog/db"
	"net/http"
)

func (s *Store) Images(w http.ResponseWriter, r *http.Request) {
	type image struct {
		ID  uint   `json:"id"`
		Url string `json:"url"`
	}

	log := s.log.WithField("api", "images")
	w.Header().Add("Content-Type", "application/json")
	imagesData := make([]db.Image, 0)
	result := s.db.Find(&imagesData)
	if result.Error != nil {
		log.WithField("error", result.Error).Error("fetch images error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while fetching images from the database")
		return
	}
	images := make([]image, 0)
	for _, i := range imagesData {
		images = append(images, image{i.ID, i.Path})
	}
	_ = json.NewEncoder(w).Encode(images)
}
