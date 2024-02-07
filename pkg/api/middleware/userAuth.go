package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthMiddleware(c *gin.Context) {
	tokenString,err:= c.Cookie("Authorization")
	if err!=nil{
		return 
	}

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte("usersecret"), nil
	// })
	jwtToken,err:=ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	if err != nil || !jwtToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err,})
		c.Abort()
		return
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		c.JSON(http.StatusForbidden, gin.H{"error": err})
		c.Abort()
		return
	}
	role, ok := claims["role"].(string)
	if !ok || role != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized access"})
		c.Abort()
		return
	}

	id, ok := claims["id"].(float64)

	fmt.Println("user id",id)
	
	if !ok || id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": err})
		c.Abort()
		return
	}
	// userIdString := fmt.Sprintf("%v", id)

	c.Set("role", role)
	c.Set("id", int(id))

	fmt.Println("user auth is fine")

	c.Next()
}
