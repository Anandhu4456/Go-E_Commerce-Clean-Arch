package interfaces

import "github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"

type OtpRepository interface {
	FindUserByMobileNumber(phone string)bool
	UserDetailsUsingPhone(phone string)(models.UserResponse,error)
}