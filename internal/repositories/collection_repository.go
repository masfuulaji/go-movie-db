package repositories

import (
	"math"

	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/response"
	"gorm.io/gorm"
)


var (
	collections []models.Collection
	collection  models.Collection
)

type CollectionRepository interface {
    CreateCollection(collection models.Collection) error
    GetCollectionById(collectionID string) (response.APICollection, error)
    GetAllCollections(page, limit int) (response.PaginatedResponse, error)
    UpdateCollection(collectionID string, collection models.Collection) (models.Collection, error)
    DeleteCollection(collectionID string) error
}

type CollectionRepositoryImpl struct {
    db *gorm.DB
}

func NewCollectionRepository(db *gorm.DB) CollectionRepository {
    return &CollectionRepositoryImpl{
    	db: db,
    }
}

func (c *CollectionRepositoryImpl) CreateCollection(collection models.Collection) error {
	return c.db.Create(&collection).Error
}

func (c *CollectionRepositoryImpl) GetCollectionById(collectionID string) (response.APICollection, error) {
	var result response.APICollection
	return result, c.db.Where("id = ?", collectionID).First(&result).Error
}

func (c *CollectionRepositoryImpl) GetAllCollections(page, limit int) (response.PaginatedResponse, error) {
    var results []response.APICollection
    var totalItems int64

    offset := (page - 1) * limit

    err := c.db.Model(&collections).Offset(offset).Limit(limit).Find(&results).Error
    if err != nil {
        return response.PaginatedResponse{}, err
    }

    err = c.db.Model(&collections).Count(&totalItems).Error
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

func (c *CollectionRepositoryImpl) UpdateCollection(collectionID string, collection models.Collection) (models.Collection, error) {
    var updatedCollection models.Collection
    err := c.db.Model(&collection).Where("id = ?", collectionID).Updates(collection).Error
    return updatedCollection, err
}

func (c *CollectionRepositoryImpl) DeleteCollection(collectionID string) error {
    return c.db.Delete(&models.Collection{}, collectionID).Error
}
