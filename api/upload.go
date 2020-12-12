package api

import (
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(r.Method + " method is not allowed"))
		return
	}
	err := r.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	fileHeaders, ok := r.MultipartForm.File["image"]
	if !ok || len(fileHeaders) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Data with key:image is not available"))
		return
	}
	fileHeader := fileHeaders[0]
	img, err := fileHeader.Open()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	defer img.Close()
	fileName := uuid.New().String() + strings.ToLower(path.Ext(fileHeader.Filename))
	file, err := os.Create(path.Join("images", fileName))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(img)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("File uploaded successfully."))
}
