package repositories

import (
	"math"

	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/response"
	"gorm.io/gorm"
)


var (
    movies []models.Movie
    movie  models.Movie
)

type MovieRepository interface {
    CreateMovie(movie models.Movie) error
    GetAllMovies(page, limit int) (response.PaginatedResponse, error)
    GetMovieById(movieID string) (response.APIMovie, error)
    UpdateMovie(movieID string, movie models.Movie) (models.Movie, error)
    DeleteMovie(movieID string) error
}

type MovieRepositoryImpl struct {
    db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
    return &MovieRepositoryImpl{
        db: db,
    }
}

func (c *MovieRepositoryImpl) CreateMovie(movie models.Movie) error {
    return c.db.Create(&movie).Error
}

func (c *MovieRepositoryImpl) GetMovieById(movieID string) (response.APIMovie, error) {
    var result response.APIMovie
    return result, c.db.Where("id = ?", movieID).First(&result).Error
}

func (c *MovieRepositoryImpl) GetAllMovies(page, limit int) (response.PaginatedResponse, error) {
    var results []response.APIMovie
    var totalItems int64

    offset := (page - 1) * limit

    err := c.db.Model(&movies).Offset(offset).Limit(limit).Find(&results).Error
    if err != nil {
        return response.PaginatedResponse{}, err
    }

    err = c.db.Model(&movies).Count(&totalItems).Error
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

func (c *MovieRepositoryImpl) UpdateMovie(movieID string, movie models.Movie) (models.Movie, error) {
    var updatedMovie models.Movie
    err := c.db.Model(&movie).Where("id = ?", movieID).Updates(movie).Error
    return updatedMovie, err
}

func (c *MovieRepositoryImpl) DeleteMovie(movieID string) error {
    return c.db.Delete(&models.Movie{}, movieID).Error
}


