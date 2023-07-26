package repositories

import (
	"math"

	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/response"
	"gorm.io/gorm"
)

var (
	user  models.User
	users []models.User
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetAllUser(page, limit int) (response.PaginatedResponse, error)
	GetUserById(userID string) (response.APIUser, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(userID string, user models.User) (models.User, error)
	DeleteUser(userID string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (c *UserRepositoryImpl) CreateUser(user models.User) error {
	return c.db.Create(&user).Error
}

func (c *UserRepositoryImpl) GetUserById(userID string) (response.APIUser, error) {
	var result response.APIUser
	return result, c.db.Model(&user).First(&result, userID).Error
}

func (c *UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {
	return user, c.db.Where("email = ?", email).First(&user).Error
}

func (c *UserRepositoryImpl) UpdateUser(userID string, user models.User) (models.User, error) {
	var updatedUser models.User
	err := c.db.Model(&user).Where("id = ?", userID).Updates(user).Error
	return updatedUser, err
}

func (c *UserRepositoryImpl) DeleteUser(userID string) error {
	return c.db.Delete(&models.User{}, userID).Error
}

func (c *UserRepositoryImpl) GetAllUser(page, limit int) (response.PaginatedResponse, error) {
	var results []response.APIUser
	var totalItems int64

	offset := (page - 1) * limit
	err := c.db.Model(&users).Offset(offset).Limit(limit).Find(&results).Error
	if err != nil {
		return response.PaginatedResponse{}, err
	}

	err = c.db.Model(&users).Count(&totalItems).Error
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
