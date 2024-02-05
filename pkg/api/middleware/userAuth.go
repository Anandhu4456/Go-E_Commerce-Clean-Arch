package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("userpass"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
		c.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid authorization token"})
		c.Abort()
		return
	}
	role, ok := claims["role"].(string)
	if !ok || role != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized access"})
		c.Abort()
		return
	}
	id, ok := claims["user_id"].(float64)
	if !ok || id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "error in retrieving id"})
		c.Abort()
		return
	}
	// userIdString := fmt.Sprintf("%v", id)

	c.Set("role", role)
	c.Set("id", int(id))

	c.Next()
}
