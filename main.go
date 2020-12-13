package main

import (
	"image_catalog/api"
	"image_catalog/db"
	"log"
	"net/http"
)

func main() {
	database, err := db.Init()
	if err != nil {
		log.Fatalf("Error initializing database: %s", err.Error())
	}
	s := api.Store{DB: database}
	http.HandleFunc("/api/upload/", s.Upload)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
