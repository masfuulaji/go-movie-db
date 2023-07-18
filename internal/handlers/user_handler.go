package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(c *gin.Context) {
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

	err = repositories.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}

func GetUserHandler(c *gin.Context) {
	userID := c.Param("userID")

	user, err := repositories.GetUserById(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func UpdateUserHandler(c *gin.Context) {
	var user models.User
	userID := c.Param("userID")

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

	user, err = repositories.UpdateUser(userID, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("userID")

	err := repositories.DeleteUser(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted"})
}

func GetAllUserHandler(c *gin.Context) {
	users, err := repositories.GetAllUser()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
