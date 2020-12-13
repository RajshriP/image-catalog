package main

import (
	"image_catalog/api"
	"image_catalog/db"
	"log"
)

func main() {
	database, err := db.Init()
	if err != nil {
		log.Fatalf("Error initializing database: %s", err.Error())
	}
	s := api.New(database)
	log.Fatal(s.ListenAndServe(":8000"))
}
