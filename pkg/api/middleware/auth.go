package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware(c *gin.Context) {
	token, _ := c.Cookie("Authorization")
	fmt.Println("Token:", token)
	jwtToken, err := validateToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err != nil || !jwtToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
		c.Abort()
		return
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
		c.Abort()
		return
	}
	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized access"})
		c.Abort()
		return
	}
	c.Next()
}
