package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"image_catalog/db"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func (s *Store) Upload(w http.ResponseWriter, r *http.Request) {
	log := s.log.WithField("api", "upload")

	w.Header().Add("Content-Type", "application/json")

	if r.Method != "POST" {
		log.Debug(r.Method + " request received")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = errMessage(w, r.Method+" method not allowed")
		return
	}
	err := r.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		log.WithField("error", err).Info("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		_ = errMessage(w, "Invalid request body: "+err.Error())
		return
	}
	fileHeaders, ok := r.MultipartForm.File["image"]
	if !ok || len(fileHeaders) == 0 {
		log.WithField("error", err).Info("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		_ = errMessage(w, "Invalid request body: image is required")
		return
	}
	fileHeader := fileHeaders[0]
	img, err := fileHeader.Open()
	if err != nil {
		log.WithField("error", err).Error("read image error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while reading image")
		return
	}
	defer img.Close()

	ext := strings.ToLower(path.Ext(fileHeader.Filename))
	image := db.Image{Path: "/images/" + uuid.New().String() + ext}
	root := "."
	fp := root + image.Path
	file, err := os.Create(fp)
	if err != nil {
		log.WithField("error", err).Error("create file error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while storing image")
		return
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(img)
	if err != nil {
		log.WithField("error", err).Error("read file error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while reading image")
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		log.WithField("error", err).Error("write to file error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while storing image")
		defer os.Remove(fp) // Should be called after file.Close()
		return
	}

	result := s.db.Create(&image)
	if result.Error != nil {
		log.WithField("error", err).Error("db insert error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while updating database")
		defer os.Remove(fp) // Should be called after file.Close()
		return
	}
	log.WithFields(logrus.Fields{"path": image.Path, "id": image.ID}).Info("file uploaded")
	_ = json.NewEncoder(w).Encode(image)
}
