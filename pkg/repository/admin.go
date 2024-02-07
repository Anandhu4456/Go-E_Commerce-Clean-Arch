package repository

import (
	"fmt"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB)interfaces.AdminRepository{
	return &adminRepository{
		DB:DB,
	}
}

func (ar *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {
	var adminCompareDetail domain.Admin
	err := ar.DB.Raw("select * from admins where user_name = ?", adminDetails.Email).Scan(&adminCompareDetail).Error
	if err != nil  {
		return domain.Admin{}, err
	}
	fmt.Println("admincompare details from repo",adminCompareDetail)
	return adminCompareDetail, nil
}

func (ar *adminRepository)GetUserById(id string)(domain.User,error){
	userId,err:=strconv.Atoi(id)
	if err!=nil{
		return domain.User{},err
	}
	query:=fmt.Sprintf("select * from users where id = '%d'",userId)
	var userDetails domain.User
	err =ar.DB.Raw(query).Scan(&userDetails).Error
	if err!=nil{
		return domain.User{},err
	}
	return userDetails,nil
}

// This function will both block and unblock user
func (ar *adminRepository)UpdateBlockUserById(user domain.User)error{
	err:=ar.DB.Exec("update users set permission = ? where id = ?",user.Permission,user.ID).Error
	if err!=nil{
		return err
	}
	return nil
}

func (ar *adminRepository)GetUsers(page,limit int)([]models.UserDetailsAtAdmin,error){
	if page== 0{
		page = 1
	}
	if limit==0{
		limit =10
	}
	offset:=(page-1)*limit
	var userDetails []models.UserDetailsAtAdmin
	err:=ar.DB.Raw("select id,name,email,phone,permission from users limit ? offset ?",limit,offset).Scan(&userDetails).Error
	if err!=nil{
		return []models.UserDetailsAtAdmin{},err
	}
	return userDetails,nil
}