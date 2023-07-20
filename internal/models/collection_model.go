package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name        string `gorm:"not null;"`
	Description string
	PosterPath  string
	Movies      []*Movie
}
