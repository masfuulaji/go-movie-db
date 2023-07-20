package repositories

import (
	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/models"
)

type APICollection struct {
	Name string
}

var (
	collections []models.Collection
	collection  models.Collection
)

func CreateCollection(collection models.Collection) error {
	return config.DB.Create(&collection).Error
}

func GetCollectionById(collectionID string) (APICollection, error) {
	var result APICollection
	return result, config.DB.Where("id = ?", collectionID).First(&result).Error
}

func GetAllCollections() ([]APICollection, error) {
    var results []APICollection
    return results, config.DB.Model(&collections).Find(&results).Error
}

func UpdateCollection(collectionID string, collection models.Collection) (models.Collection, error) {
    var updatedCollection models.Collection
    err := config.DB.Model(&collection).Where("id = ?", collectionID).Updates(collection).Error
    return updatedCollection, err
}

func DeleteCollection(collectionID string) error {
    return config.DB.Delete(&models.Collection{}, collectionID).Error
}
