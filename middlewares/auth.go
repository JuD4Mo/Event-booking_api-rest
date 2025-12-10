package middlewares

import (
	"net/http"

	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	//Proteger el route con las funcionalidades del token
	token := c.Request.Header.Get("Authorization")

	//El token es vacío
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	//El token está pero es inválido
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
