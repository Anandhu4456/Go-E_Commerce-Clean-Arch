package domain

type User struct {
	ID         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone" gorm:"unique"`
	Permission bool   `json:"permission" gorm:"default:true"`
}

type Address struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Default   bool   `json:"default" gorm:"default:false"`
}
