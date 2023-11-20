package usecase

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
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

func (usrU *userUsecase) SignUp(user models.UserDetails) (models.UserToken, error) {
	// check the user exist or not,if exist show the error(its a signup function)

	userExist := usrU.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.UserToken{}, errors.New("user already exist please sign in")
	}
	if user.Password != user.ConfirmPassword {
		return models.UserToken{}, errors.New("password does't match")
	}
	// hash the password
	hashedPass, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return models.UserToken{}, err
	}
	user.Password = hashedPass
	// insert the user into database
	userData, err := usrU.userRepo.SignUp(user)
	if err != nil {
		return models.UserToken{}, err
	}
	// create jwt token for user
	tokenString, err := helper.GenerateUserToken(userData)
	if err != nil {
		return models.UserToken{}, errors.New("couldn't create token for user due to some internal error")
	}
	return models.UserToken{
		Username: user.Username,
		Token:    tokenString,
	}, nil

}

func (usrU *userUsecase) AddAddress(id int, address models.AddAddress) error {
	rslt := usrU.userRepo.CheckIfFirstAddress(id)
	var checkAddress bool

	if !rslt {
		checkAddress = true
	} else {
		checkAddress = false
	}
	if err := usrU.userRepo.AddAddress(id, address, checkAddress); err != nil {
		return err
	}
	return nil
}

func (usrU *userUsecase) GetAddresses(id int) ([]domain.Address, error) {
	addresses, err := usrU.userRepo.GetAddresses(id)
	if err != nil {
		return []domain.Address{}, err
	}
	return addresses, nil
}

func (usrU *userUsecase) GetUserDetails(id int) (models.UserResponse, error) {
	userDetails, err := usrU.userRepo.GetUserDetails(id)
	if err != nil {
		return models.UserResponse{}, err
	}
	return userDetails, nil
}
