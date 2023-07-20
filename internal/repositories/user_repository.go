package repositories

import (
	"math"

	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/response"
)

var (
	user  models.User
	users []models.User
)

func CreateUser(user models.User) error {
	return config.DB.Create(&user).Error
}

func GetUserById(userID string) (response.APIUser, error) {
	var result response.APIUser
	return result, config.DB.Model(&user).First(&result, userID).Error
}

func GetUserByEmail(email string) (models.User, error) {
	return user, config.DB.Where("email = ?", email).First(&user).Error
}

func UpdateUser(userID string, user models.User) (models.User, error) {
	var updatedUser models.User
	err := config.DB.Model(&user).Where("id = ?", userID).Updates(user).Error
	return updatedUser, err
}

func DeleteUser(userID string) error {
	return config.DB.Delete(&models.User{}, userID).Error
}

func GetAllUser(page, limit int) (response.PaginatedResponse, error) {
	var results []response.APIUser
    var totalItems int64

    offset := (page - 1) * limit
    err := config.DB.Model(&users).Offset(offset).Limit(limit).Find(&results).Error
    if err != nil {
        return response.PaginatedResponse{}, err
    }

    err = config.DB.Model(&users).Count(&totalItems).Error
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
