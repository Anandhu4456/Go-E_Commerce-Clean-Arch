package models

import "time"

type AdminLogin struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=6,max=12"`
}

type AdminDetailsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminToken struct {
	Username string
	Token    string
}

type UserDetailsAtAdmin struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Permission bool   `json:"permission"`
}

type CustomDates struct {
	StartingDate time.Time `json:"starting_date"`
	EndDate      time.Time `json:"end_date"`
}
