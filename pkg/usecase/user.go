package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo   interfaces.UserRepository
	offerRepo  interfaces.OfferRepository
	walletRepo interfaces.WalletRepository
}

// Constructor function
func NewUserUsecase(userRepo interfaces.UserRepository, offerRepo interfaces.OfferRepository, walletRepo interfaces.WalletRepository) services.UserUsecase {
	return &userUsecase{
		userRepo:   userRepo,
		offerRepo:  offerRepo,
		walletRepo: walletRepo,
	}
}

func (usrU *userUsecase) Login(user models.UserLogin) (models.UserToken, error) {
	// check the user already exist or not

	ok := usrU.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.UserToken{}, errors.New("user not exist")
	}
	// check admin blocked this user or not
	permission, err := usrU.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.UserToken{}, err
	}
	if !permission {
		return models.UserToken{}, errors.New("user is blocked by admin")
	}
	// Get the user details in order to check password
	userDetails, err := usrU.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.UserToken{}, err
	}
	// check the password
	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password))
	if err != nil {
		return models.UserToken{}, errors.New("password incorrect")
	}
	// generate token
	tokenString, err := helper.GenerateUserToken(userDetails)
	if err != nil {
		return models.UserToken{}, errors.New("could't create token for user")
	}
	return models.UserToken{
		Username: userDetails.Username,
		Token:    tokenString,
	}, nil
}
