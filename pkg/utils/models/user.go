package models

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User Signup
type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"required"`
	Username        string `json:"username"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserToken struct {
	User  UserDetailsResponse
	Token string
}

// user details shown after loggin
type UserDetailsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	Phone string `json:"phone"`
}

type UserSignInResponse struct{
	Id       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}


type Address struct {
	Id        uint   `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

type AddAddress struct {
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	RePassword  string `json:"re_password"`
}

type GetCart struct {
	Id            int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	CategoryId    int     `json:"category_id"`
	Quantity      int     `json:"quantity"`
	Total         float64 `json:"total"`
	DiscountPrice float64 `json:"discount_price"`
}

type GetCartResponse struct {
	Id     int
	Values []GetCart
}

type CheckOut struct {
	CartId         int
	Addresses      []Address
	Products       []GetCart
	PaymentMethods []PaymentMethod
	TotalPrice     float64
	DiscountPrice  float64
}
type EditUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserKey string

func (k UserKey) String() string {
	return string(k)
}
