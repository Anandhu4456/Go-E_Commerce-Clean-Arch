package helper

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func GetUserId(c *gin.Context) (int, error) {
	var key models.UserKey = "user_id"
	val := c.Request.Context().Value(key)

	// Check if the value is not nil
	if val == nil {
		return 0, errors.New("user id not found in context")
	}
	// using type assertion to chech the type of val is models.UserKey

	userkey, ok := val.(models.UserKey)
	if !ok {
		return 0, errors.New("user id type is not expected type")
	}
	id := userkey.String()
	userId, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("failed to convert user id to int")
	}
	return userId, nil
}

func GenerateUserToken(user models.UserResponse) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   user.Username,
		"role":   "user",
		"userId": user.Id,
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("KEY")))
	if err == nil {
		fmt.Println("token created")
	}
	return tokenString, nil
}

func PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashedPassword)
	return hash, nil
}
