package models

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	Username string
	Token    string
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
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
