package repositories

import (
	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/models"
)

type APIMovie struct {
    Name  string
}

var (
    movies []models.Movie
    movie  models.Movie
)

func CreateMovie(movie models.Movie) error {
    return config.DB.Create(&movie).Error
}

func GetMovieById(movieID string) (APIMovie, error) {
    var result APIMovie
    return result, config.DB.Where("id = ?", movieID).First(&result).Error
}

func GetAllMovies() ([]APIMovie, error) {
    var results []APIMovie
    return results, config.DB.Model(&movies).Find(&results).Error
}

func UpdateMovie(movieID string, movie models.Movie) (models.Movie, error) {
    var updatedMovie models.Movie
    err := config.DB.Model(&movie).Where("id = ?", movieID).Updates(movie).Error
    return updatedMovie, err
}

func DeleteMovie(movieID string) error {
    return config.DB.Delete(&models.Movie{}, movieID).Error
}


