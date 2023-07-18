package repositories

import (
	"github.com/masfuulaji/go-movie-db/internal/config"
	"github.com/masfuulaji/go-movie-db/internal/models"
)

func CreateUser(user models.User) error {
	return config.DB.Create(&user).Error
}

func GetUserById(userID string) (models.User, error) {
	var user models.User
	return user, config.DB.First(&user, userID).Error
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	return user, config.DB.Where("email = ?", email).First(&user).Error
}

func UpdateUser(userID string, user models.User) (models.User, error) {
	var updatedUser models.User
	err := config.DB.Model(&updatedUser).Where("id = ?", userID).Updates(user).Error
	return updatedUser, err
}

func DeleteUser(userID string) error {
	return config.DB.Delete(&models.User{}, userID).Error
}

func GetAllUser() ([]models.User, error) {
	var users []models.User
	return users, config.DB.Find(&users).Error

}
