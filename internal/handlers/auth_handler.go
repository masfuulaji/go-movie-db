package handlers

import (

	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
    userRepo repositories.UserRepository
}

func NewAuthHandler(repo repositories.UserRepository) *AuthHandler {
    return &AuthHandler{userRepo: repo}
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	result, err := h.userRepo.GetUserByEmail(email)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

    token, err := h.userRepo.GenerateToken(int(result.ID))
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"token": token})
}

func (h *AuthHandler) LogoutHandler(c *gin.Context) {

    c.JSON(200, gin.H{"message": "logout success"})
}
