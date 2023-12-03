package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func UserAuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		c.Abort()
		return
	}
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		secret := viper.GetString("KEY")
		return []byte(secret), nil
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
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized access"})
		c.Abort()
		return
	}
	userIdString := fmt.Sprintf("%v", id)

	var key models.UserKey = "userId"
	var val models.UserKey = models.UserKey(userIdString)

	ctx := context.WithValue(c, key, val)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
