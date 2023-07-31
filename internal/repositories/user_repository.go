package repositories

import (
	"math"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/request"
	"github.com/masfuulaji/go-movie-db/internal/response"
	"gorm.io/gorm"
)

var (
	user  models.User
	users []models.User
)

type UserRepository interface {
	CreateUser(user request.UserCreateRequest) error
	GetAllUser(page, limit int) (response.PaginatedResponse, error)
	GetUserById(userID string) (response.APIUser, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(userID string, user request.UserUpdateRequest) (models.User, error)
	DeleteUser(userID string) error
    GenerateToken(userID int) (string, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (c *UserRepositoryImpl) CreateUser(user request.UserCreateRequest) error {
	return c.db.Create(&user).Error
}

func (c *UserRepositoryImpl) GetUserById(userID string) (response.APIUser, error) {
	var result response.APIUser
	return result, c.db.Model(&user).First(&result, userID).Error
}

func (c *UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {
	return user, c.db.Where("email = ?", email).First(&user).Error
}

func (c *UserRepositoryImpl) UpdateUser(userID string, user request.UserUpdateRequest) (models.User, error) {
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

func (c *UserRepositoryImpl) GenerateToken(userID int) (string, error) {
    secret := os.Getenv("JWT_SECRET")

    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
