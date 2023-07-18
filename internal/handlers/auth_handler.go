package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	result, err := repositories.GetUserByEmail(email)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sessions := sessions.Default(c)
	sessions.Set("user_id", result.ID)
	sessions.Save()

	c.JSON(200, gin.H{"message": "login success"})
}

func LogoutHandler(c *gin.Context) {
    sessions := sessions.Default(c)
    sessions.Delete("user_id")
    sessions.Save()

    c.JSON(200, gin.H{"message": "logout success"})
}
