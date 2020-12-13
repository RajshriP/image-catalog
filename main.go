package main

import (
	"github.com/sirupsen/logrus"
	"image_catalog/api"
	"image_catalog/db"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	database, err := db.Init()
	if err != nil {
		log.Fatalf("Error initializing database: %s", err.Error())
	}
	s := api.New(database, log)
	log.Fatal(s.ListenAndServe(":8000"))
}
