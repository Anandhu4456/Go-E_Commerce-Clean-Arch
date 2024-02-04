package db

import (
	"fmt"

	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err:=db.AutoMigrate(&domain.Inventory{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Category{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Admin{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.User{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Cart{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Wishlist{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Address{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Order{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.OrderItem{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.LineItems{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.PaymentMethod{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Offer{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Coupon{});err!=nil{
		return db,err
	}
	if err:=db.AutoMigrate(&domain.Wallet{});err!=nil{
		return db,err
	}
	

	CheckAndCreateAdmin(db)
	return db, dbErr
}

func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "adminpass"
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		admin := domain.Admin{
			ID:       1,
			Name: "admin",
			Username: "yoursstore@gmail.com",
			Password: string(hashedPass),
		}
		db.Create(&admin)
	}
}
