package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name        string
	Description string
	Movies      []*Movie `gorm:"many2many:movie_genres;"`
}
