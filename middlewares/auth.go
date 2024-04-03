package middlewares

import (
	"net/http"

	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not verified"})
		return
	}

	c.Set("userId", userId)

	c.Next()
}
