package db

import (
	"fmt"

	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Inventory{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.WishList{})
	db.AutoMigrate(&domain.WishlistItems{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.OrderItem{})
	db.AutoMigrate(&domain.LineItems{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.Offer{})
	db.AutoMigrate(&domain.Coupon{})
	db.AutoMigrate(&domain.Wallet{})
	db.AutoMigrate(&domain.WalletHistory{})
	db.AutoMigrate(&domain.Image{})

	return db, dbErr
}
