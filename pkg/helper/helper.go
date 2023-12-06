package helper

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
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

func GenerateAdminToken(admin domain.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": admin.Username,
		"role":  "admin",
	})
	tokenString, err := token.SignedString(viper.GetString("KEY"))
	if err == nil {
		fmt.Println("token created")
	}
	return tokenString, nil
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

// This function will setup the twilio
var client *twilio.RestClient

func TwilioSetup(username string, password string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})
}

func TwilioSendOTP(phone string, serviceId string) (string, error) {
	to := "+91" + phone
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(serviceId, params)

	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

func TwilioVerifyOTP(serviceId string, code string, phone string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(serviceId, params)
	if err != nil {
		return err
	}
	if *resp.Status == "approved" {
		return nil
	}
	return errors.New("otp verification failed")
}

func FindMostBroughtProduct(products []domain.ProductReport) []int {
	productMap := make(map[int]int)

	for _, item := range products {
		productMap[item.InventoryID] += item.Quantity
	}
	maxQty := 0
	for _, item := range productMap {
		if item > maxQty {
			maxQty = item
		}
	}
	var bestSeller []int
	for k, item := range productMap {
		if item == maxQty {
			bestSeller = append(bestSeller, k)
		}
	}
	return bestSeller
}
