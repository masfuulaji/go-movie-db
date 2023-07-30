package repositories

import (
	"math"

	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/response"
	"gorm.io/gorm"
)

var (
	genres []models.Genre
	genre  models.Genre
)

type GenreRepository interface {
    CreateGenre(genre models.Genre) error
    GetGenreById(genreID string) (response.APIGenre, error)
    GetAllGenres(page, limit int) (response.PaginatedResponse, error)
    UpdateGenre(genreID string, genre models.Genre) (models.Genre, error)
    DeleteGenre(genreID string) error
}

type GenreRepositoryImpl struct {
    db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
    return &GenreRepositoryImpl{
        db: db,
    }
}

func (c *GenreRepositoryImpl) CreateGenre(genre models.Genre) error {
	return c.db.Create(&genre).Error
}

func (c *GenreRepositoryImpl) GetGenreById(genreID string) (response.APIGenre, error) {
	var result response.APIGenre
	return result, c.db.Where("id = ?", genreID).First(&result).Error
}

func (c *GenreRepositoryImpl) GetAllGenres(page, limit int) (response.PaginatedResponse, error) {
	var results []response.APIGenre
    var totalItems int64

    offset := (page - 1) * limit
	err := c.db.Model(&genres).Offset(offset).Limit(limit).Find(&results).Error
    if err != nil {
        return response.PaginatedResponse{}, err
    }

    err = c.db.Model(&genres).Count(&totalItems).Error
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

func (c *GenreRepositoryImpl) UpdateGenre(genreID string, genre models.Genre) (models.Genre, error) {
	var updatedGenre models.Genre
	err := c.db.Model(&genre).Where("id = ?", genreID).Updates(genre).Error
	return updatedGenre, err
}

func (c *GenreRepositoryImpl) DeleteGenre(genreID string) error {
	return c.db.Delete(&models.Genre{}, genreID).Error
}
