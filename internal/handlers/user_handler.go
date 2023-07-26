package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
    userRepo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
    return &UserHandler{userRepo: repo}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hashedPassword)

	err = h.userRepo.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}

func (h *UserHandler) GetUserHandler(c *gin.Context) {
	userID := c.Param("user_id")

	user, err := h.userRepo.GetUserById(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	var user models.User
	userID := c.Param("user_id")

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user.Password = string(hashedPassword)
	}

	result, err := h.userRepo.UpdateUser(userID, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	userID := c.Param("user_id")

	err := h.userRepo.DeleteUser(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted"})
}

func (h *UserHandler) GetAllUserHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	users, err := h.userRepo.GetAllUser(page, limit)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
