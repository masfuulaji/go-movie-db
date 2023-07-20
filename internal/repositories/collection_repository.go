package repositories

import (
	"math"

	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/response"
)


var (
	collections []models.Collection
	collection  models.Collection
)

func CreateCollection(collection models.Collection) error {
	return config.DB.Create(&collection).Error
}

func GetCollectionById(collectionID string) (response.APICollection, error) {
	var result response.APICollection
	return result, config.DB.Where("id = ?", collectionID).First(&result).Error
}

func GetAllCollections(page, limit int) (response.PaginatedResponse, error) {
    var results []response.APICollection
    var totalItems int64

    offset := (page - 1) * limit

    err := config.DB.Model(&collections).Offset(offset).Limit(limit).Find(&results).Error
    if err != nil {
        return response.PaginatedResponse{}, err
    }

    err = config.DB.Model(&collections).Count(&totalItems).Error
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

func UpdateCollection(collectionID string, collection models.Collection) (models.Collection, error) {
    var updatedCollection models.Collection
    err := config.DB.Model(&collection).Where("id = ?", collectionID).Updates(collection).Error
    return updatedCollection, err
}

func DeleteCollection(collectionID string) error {
    return config.DB.Delete(&models.Collection{}, collectionID).Error
}
