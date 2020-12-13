package db

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Path string
}
