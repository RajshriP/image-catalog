package db

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Path string
}

func (img Image) MarshalJSON() ([]byte, error) {
	type image struct {
		ID   uint   `json:"id"`
		Path string `json:"path"`
	}
	return json.Marshal(image{
		ID:   img.ID,
		Path: img.Path,
	})
}
