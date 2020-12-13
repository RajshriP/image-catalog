package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dbName := "images.db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Image{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
