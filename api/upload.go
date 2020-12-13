package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"image_catalog/db"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func (s Store) Upload(w http.ResponseWriter, r *http.Request) {
	jw := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = errMessage(w, r.Method+" method not allowed")
		return
	}
	err := r.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = errMessage(w, "Invalid request body: "+err.Error())
		return
	}
	fileHeaders, ok := r.MultipartForm.File["image"]
	if !ok || len(fileHeaders) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = errMessage(w, "Invalid request body: image is required")
		return
	}
	fileHeader := fileHeaders[0]
	img, err := fileHeader.Open()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while reading image")
		return
	}
	defer img.Close()
	fileName := uuid.New().String() + strings.ToLower(path.Ext(fileHeader.Filename))
	fp := path.Join("images", fileName)
	file, err := os.Create(fp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while storing image")
		return
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(img)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while reading image")
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while storing image")
		defer os.Remove(fp) // Should be called after file.Close()
		return
	}
	image := db.Image{Path: fp}
	result := s.DB.Create(&image)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = errMessage(w, "Error while updating database")
		defer os.Remove(fp) // Should be called after file.Close()
		return
	}
	type response struct {
		ID  uint   `json:"id"`
		Url string `json:"url"`
	}
	w.Header().Add("Content-Type", "application/json")
	out := response{ID: image.ID, Url: "/" + image.Path}
	_ = jw.Encode(out)
}
