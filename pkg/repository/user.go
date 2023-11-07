package repository

import (
	"errors"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// constructor funciton

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (ur *userRepository) CheckUserAvailability(email string) bool {
	var userCount int

	err := ur.DB.Raw("SELECT COUNT(*) FROM users WHERE email=?", email).Scan(&userCount).Error
	if err != nil {
		return false
	}
	// if count greater than 0, user already exist
	return userCount > 0
}

func (ur *userRepository) UserBlockStatus(email string) (bool, error) {
	var permission bool
	err := ur.DB.Raw("SELECT permission FROM users WHERE email=?", email).Scan(&permission).Error
	if err != nil {
		return false, err
	}
	return permission, nil
}

func (ur *userRepository) FindUserByEmail(user models.UserLogin) (models.UserResponse, error) {
	var userResponse models.UserResponse
	err := ur.DB.Raw("SELECT * FROM users WHERE email=? AND permission=true", user.Email).Scan(&userResponse).Error
	if err != nil {
		return models.UserResponse{}, errors.New("no user found")
	}
	return userResponse, nil
}

func (ur *userRepository) FindUserIDByOrderID(orderID int) (int, error) {
	var userId int
	err := ur.DB.Raw("SELECT user_id FROM orders WHERE order_id=?", orderID).Scan(&userId).Error
	if err != nil {
		return 0, errors.New("user id not found")
	}
	return userId, nil
}

func (ur *userRepository) SignUp(user models.UserDetails) (models.UserResponse, error) {
	var userResponse models.UserResponse
	err := ur.DB.Exec("INSERT INTO users(name,email,username,phone,password)VALUES(?,?,?,?,?)RETURNING id,name,email,phone", user.Name, user.Email, user.Username, user.Phone, user.Password).Scan(&userResponse).Error
	if err != nil {
		return models.UserResponse{}, err
	}
	return userResponse, nil
}

func (ur *userRepository) AddAddress(id int, address models.AddAddress, result bool) error {
	query := `
	
	INSERT INTO addresses(user_id,name,house_name,street,city,state,pin,"default")
	VALUES($1,$2,$3,$4,$5,$6,$7,$8)
	RETURNING id
	`
	err := ur.DB.Exec(query, id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin, result).Error
	if err != nil {
		return errors.New("adding address failed")
	}
	return nil
}

func (ur *userRepository) CheckIfFirstAddress(id int) bool {
	var addressCount int
	err := ur.DB.Raw("SELECT COUNT(*)FROM addresses WHERE user_id=?", id).Scan(&addressCount).Error
	if err != nil {
		return false
	}
	// if addresscount >0 there is already a address
	return addressCount > 0
}

func (ur *userRepository) GetAddresses(id int) ([]domain.Address, error) {
	var getAddress []domain.Address
	err := ur.DB.Raw("SELECT * FROM addresses WHERE id=?", id).Scan(&getAddress).Error
	if err != nil {
		return []domain.Address{}, errors.New("failed to getting address")
	}
	return getAddress, nil
}

func (ur *userRepository) GetUserDetails(id int) (models.UserResponse, error) {
	var userDetails models.UserResponse

	err := ur.DB.Raw("SELECT * FROM users WHERE id=?", id).Scan(&userDetails).Error
	if err != nil {
		return models.UserResponse{}, errors.New("error while getting user details")
	}
	return userDetails, nil
}

func (ur *userRepository) ChangePassword(id int, password string) error {
	err := ur.DB.Exec("UPDATE users SET password=? WHERE id=?", password, id).Error
	if err != nil {
		return errors.New("password changing failed")
	}
	return nil
}

func (ur *userRepository) GetPassword(id int) (string, error) {
	var password string
	err := ur.DB.Raw("SELECT password FROM users WHERE id=?", id).Scan(&password).Error
	if err != nil {
		return "", errors.New("password getting failed")
	}
	return password, nil
}

func (ur *userRepository) FindIdFromPhone(phone string) (int, error) {
	var userid int
	err := ur.DB.Raw("SELECT id FROM users WHERE phone=?", phone).Scan(&userid).Error
	if err != nil {
		return 0, err
	}
	return userid, nil
}

func (ur *userRepository) EditName(id int, name string) error {
	err := ur.DB.Exec("UPDATE users SET name=? WHERE id=?", name, id).Error
	if err != nil {
		return errors.New("error while editing name")
	}
	return nil
}
