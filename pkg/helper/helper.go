package helper

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateAdminToken(admin models.AdminDetailsResponse) (string, string, error) {
	tokenClaimes := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	refreshTokenClaims := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaimes)
	tokenString, err := token.SignedString([]byte("adminsecret"))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("adminrefresh"))
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}

func GenerateUserToken(user models.UserResponse) (string, error) {
	claims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("usersecret"))
	if err != nil {
		return "", err
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

func AddImageToS3(file *multipart.FileHeader) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		fmt.Println("configuration error: ", err)
		return "", err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	f, openErr := file.Open()
	if openErr != nil {
		fmt.Println("file open error: ", openErr)
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("samplebucket"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})
	if uploadErr != nil {
		fmt.Println("upload error : ", uploadErr)
		return "", uploadErr
	}
	return result.Location, nil
}
