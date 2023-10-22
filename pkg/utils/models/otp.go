package models

type VerifyData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
}

type OTPData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
}
