package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Recover from any panics to avoid crashing the server
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		// Continue processing other middlewares and handlers
		c.Next()

		// If no routes match the request, handle the 404 Not Found error
		if c.Writer.Status() == http.StatusNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
			})
		}
	}
}
