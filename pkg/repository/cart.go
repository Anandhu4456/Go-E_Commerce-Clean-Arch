package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type cartRepository struct{
	DB *gorm.DB
}

// constructor function
func NewCartRepository (DB *gorm.DB)interfaces.CartRepository{
	return &cartRepository{
		DB:DB,
	}
}

func (cr *cartRepository)GetAddresses(id int)([]domain.Address,error){
	var addresses []domain.Address
	err:=cr.DB.Raw("select * from addresses where id = ?",id).Scan(&addresses).Error
	if err!=nil{
		return []domain.Address{},err
	}
	return addresses,nil
}

func (cr *cartRepository)GetCartId(user_id int)(int,error){
	var userId int
	err:=cr.DB.Raw("select id from carts where user_id = ?",user_id).Scan(&userId).Error
	if err!=nil{
		return 0,err
	}
	return userId,nil
}

func (cr *cartRepository)CreateNewCart(user_id int)(int,error){
	var id int
	err:=cr.DB.Exec(`INSERT INTO carts (user_id) VALUES (?)`,user_id).Error
	if err!=nil{
		return 0,err
	}
	err = cr.DB.Raw("select id from carts where user_id ",user_id).Scan(&id).Error
	if err!=nil{
		return 0,err
	}
	return id,nil
}
func (cr *cartRepository)AddLineItems(inventory_id ,cart_id int)error{
	err:=cr.DB.Exec(`INSERT INTO line_items (inventory_id,cart_id) VALUES(?,?)`,inventory_id,cart_id).Error
	if err!=nil{
		return err
	}
	return nil
}