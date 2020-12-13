package api

import (
	"encoding/json"
	"gorm.io/gorm"
	"image_catalog/db"
	"math"
	"net/http"
	"strconv"
)

type paginatedImages struct {
	PageNo  int        `json:"page_no"`
	PerPage int        `json:"per_page"`
	Data    []db.Image `json:"data"`
	Pages   int        `json:"pages"`
}

func (s *Store) Images(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithField("api", "images")
	w.Header().Add("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		log.WithField("error", err).Error("parse form error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error processing the request")
		return
	}
	pageNo, err := strconv.Atoi(r.FormValue("page_no"))
	if err != nil || pageNo < 1 {
		pageNo = 1
	}
	perPage, err := strconv.Atoi(r.FormValue("per_page"))
	if err != nil || perPage < 1 {
		perPage = 10
	}
	images, err := getPaginatedImages(s.db, pageNo, perPage)
	if err != nil {
		log.WithField("error", err).Error("fetch images error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while fetching images from the database")
		return
	}
	_ = json.NewEncoder(w).Encode(images)
}

func getPaginatedImages(gormDB *gorm.DB, pageNo, perPage int) (paginatedImages, error) {
	images := paginatedImages{
		Data:    make([]db.Image, 0, perPage),
		PerPage: perPage,
		PageNo:  pageNo,
	}
	result := gormDB.Order("created_at desc").Limit(perPage).Offset(perPage * (pageNo - 1)).Find(&images.Data)
	if result.Error != nil {
		return images, result.Error
	}
	var count int64
	result = gormDB.Model(&db.Image{}).Count(&count)
	images.Pages = int(math.Ceil(float64(count) / float64(perPage)))
	return images, result.Error
}
