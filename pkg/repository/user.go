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

func (ur *userRepository) EditEmail(id int, email string) error {
	if err := ur.DB.Exec("UPDATE users SET email=? WHERE id=?", email, id).Error; err != nil {
		return errors.New("error while changing email")
	}
	return nil
}

func (ur *userRepository) EditPhone(id int, phone string) error {
	if err := ur.DB.Exec("UPDATE users SET phone=?,WHERE id=?", phone, id).Error; err != nil {
		return errors.New("error while changing phone number")
	}
	return nil
}

func (ur *userRepository) EditUsername(id int, username string) error {
	if err := ur.DB.Exec("UPDATE users SET username=? WHERE id=?", username, id).Error; err != nil {
		return errors.New("error while changing username")
	}
	return nil
}

func (ur *userRepository) RemoveFromCart(id int, inventoryID int) error {
	if err := ur.DB.Exec("DELETE FROM line_items WHERE id=? AND inventory_id=?", id, inventoryID).Error; err != nil {
		return errors.New("item not removed")
	}
	return nil
}

func (ur *userRepository) ClearCart(cartID int) error {
	if err := ur.DB.Exec("DELETE FROM line_items WHERE cart_id=?", cartID).Error; err != nil {
		return errors.New("cart not cleared")
	}
	return nil
}

func (ur *userRepository) UpdateQuantityAdd(id, inv_id int) error {
	query := `
	UPDATE line_items SET quantity=quantity+1 
	WHERE cart_id=? AND inventory_id=?
	`
	if err := ur.DB.Exec(query, id, inv_id).Error; err != nil {
		return errors.New("failed to add quantity")
	}
	return nil
}

func (ur *userRepository) UpdateQuantityLess(id, inv_id int) error {
	query :=
		`
	UPDATE line_items SET quantity=quantity-1
	WHERE cart_id=? AND inventory_id=?
	`
	if err := ur.DB.Exec(query, id, inv_id).Error; err != nil {
		return errors.New("failed to decrease quantity")
	}
	return nil
}

func (ur *userRepository) FindUserByOrderID(orderId string) (domain.User, error) {
	var userDetails domain.User
	err := ur.DB.Raw("SELECT users.name,users.email,users.phone FROM users JOIN orders ON orders.user_id=users.id WHERE order_id=?", orderId).Scan(&userDetails).Error
	if err != nil {
		return domain.User{}, errors.New("user not found with order id")
	}
	return userDetails, nil
}

func (ur *userRepository) GetCartID(id int) (int, error) {
	var cartid int
	if err := ur.DB.Raw("SELECT id FROM cart_id WHERE user_id=?", id).Scan(&cartid).Error; err != nil {
		return 0, errors.New("cart id not found")
	}
	return cartid, nil
}

func (ur *userRepository) GetProductsInCart(cart_id, page, limit int) ([]int, error) {
	var cartProducts []int

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	err := ur.DB.Raw("SELECT inventory_id FROM line_items WHERE cart_id=?,limit=?,offset=?", cart_id, limit, offset).Scan(&cartProducts).Error
	if err != nil {
		return []int{}, err
	}
	return cartProducts, nil
}

func (ur *userRepository) FindProductNames(inventory_id int) (string, error) {
	var productName string
	if err := ur.DB.Raw("SELECT product_name FROM inventories WHERE id=?", inventory_id).Scan(&productName).Error; err != nil {
		return "", errors.New("product name not found")
	}
	return productName, nil
}
