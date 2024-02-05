package middleware

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

// middleware for admin authentication

func AdminAuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer")

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("accesssecret"), nil
	})
	if err != nil {
		fmt.Println("error in admin auth")
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}

func validateToken(token string) (*jwt.Token, error) {
	fmt.Println("token validating...")
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", t.Header["alg"])
		}
		secret := viper.GetString("KEY")
		return []byte(secret), nil
	})
	return jwtToken, err
}
