package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"autoDeleteTime"`
}
