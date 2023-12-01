package helper

import (
	"errors"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/gin-gonic/gin"
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
