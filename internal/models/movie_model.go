package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title        string `gorm:"not null;"`
	Overview     string
	ReleaseDate  time.Time
	Runtime      int     `gorm:"type:int;"`
	Budget       int     `gorm:"type:int;"`
	Revenue      int     `gorm:"type:int;"`
	Popularity   float64 `gorm:"type:float;"`
	PosterPath   string
	CollectionID int      `gorm:"type:int;"`
    Collection   *Collection
	LanguageID   int      `gorm:"type:int;"`
	Status       int      `gorm:"type:int;"`
	Genres       []*Genre `gorm:"many2many:movie_genres;"`
}
