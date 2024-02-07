package middleware

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

// middleware for admin authentication

func AdminAuthMiddleware(c *gin.Context) {
	token,err:=c.Cookie("Authorization")
	fmt.Println("first token: ",token)
	if err!=nil{
		c.AbortWithStatus(401)
		return
	}
	

	// fmt.Println("token: ",token)
	token = strings.TrimPrefix(token, "Bearer ")

	// fmt.Println("in middleware: ", token)

	jwtToken, err := ValidateToken(token)
	fmt.Println("token after validation ",jwtToken)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	if err != nil || !jwtToken.Valid {
		c.JSON(401, gin.H{"error": err})
		c.Abort()
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		c.JSON(401, gin.H{"error": err})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(401, gin.H{"error": err})
		c.Abort()
		return
	}
	fmt.Println("admin auth is fine")
	c.Next()
}

func ValidateToken(token string) (*jwt.Token, error) {
	fmt.Println("token validating...")
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		fmt.Println("signature: ", t.Signature)
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", t.Header["alg"])

		}
		
		return []byte("adminsecret"), nil
	})
	return jwtToken, err
}
