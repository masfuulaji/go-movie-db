package repositories

import (
	"math"

	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/response"
)


var (
    movies []models.Movie
    movie  models.Movie
)

func CreateMovie(movie models.Movie) error {
    return config.DB.Create(&movie).Error
}

func GetMovieById(movieID string) (response.APIMovie, error) {
    var result response.APIMovie
    return result, config.DB.Where("id = ?", movieID).First(&result).Error
}

func GetAllMovies(page, limit int) (response.PaginatedResponse, error) {
    var results []response.APIMovie
    var totalItems int64

    offset := (page - 1) * limit

    err := config.DB.Model(&movies).Offset(offset).Limit(limit).Find(&results).Error
    if err != nil {
        return response.PaginatedResponse{}, err
    }

    err = config.DB.Model(&movies).Count(&totalItems).Error
    if err != nil {
        return response.PaginatedResponse{}, err
    }

    totalPage := int(math.Ceil(float64(totalItems) / float64(limit)))

    pagination := response.PaginatedResponse{
        Page:       page,
        Result:     results,
        TotalPage:  totalPage,
        TotalItems: int(totalItems),
    }

    return pagination, nil
}

func UpdateMovie(movieID string, movie models.Movie) (models.Movie, error) {
    var updatedMovie models.Movie
    err := config.DB.Model(&movie).Where("id = ?", movieID).Updates(movie).Error
    return updatedMovie, err
}

func DeleteMovie(movieID string) error {
    return config.DB.Delete(&models.Movie{}, movieID).Error
}


