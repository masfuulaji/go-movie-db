package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        sessions := sessions.Default(c)

        if sessions.Get("user_id") == nil {
            c.AbortWithStatusJSON(401, gin.H{
                "message": "Unauthorized",
            })
            return
        }

        c.Next()
    }
}
