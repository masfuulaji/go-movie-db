package repositories

import (
	"math"

	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/response"
)

var (
	genres []models.Genre
	genre  models.Genre
)

func CreateGenre(genre models.Genre) error {
	return config.DB.Create(&genre).Error
}

func GetGenreById(genreID string) (response.APIGenre, error) {
	var result response.APIGenre
	return result, config.DB.Where("id = ?", genreID).First(&result).Error
}

func GetAllGenres(page, limit int) (response.PaginatedResponse, error) {
	var results []response.APIGenre
    var totalItems int64

    offset := (page - 1) * limit
	err := config.DB.Model(&genres).Offset(offset).Limit(limit).Find(&results).Error
    if err != nil {
        return response.PaginatedResponse{}, err
    }

    err = config.DB.Model(&genres).Count(&totalItems).Error
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

func UpdateGenre(genreID string, genre models.Genre) (models.Genre, error) {
	var updatedGenre models.Genre
	err := config.DB.Model(&genre).Where("id = ?", genreID).Updates(genre).Error
	return updatedGenre, err
}

func DeleteGenre(genreID string) error {
	return config.DB.Delete(&models.Genre{}, genreID).Error
}
