package main

import (
	"image_catalog/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/Upload/", api.Upload)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
