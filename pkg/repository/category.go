package repository

import (
	"errors"
	"strconv"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

func (cat *categoryRepository) AddCategory(category string) (domain.Category, error) {
	var b string
	err := cat.DB.Raw("INSERT INTO categories(category)VALUES(?)RETURNING category", category).Scan(&b).Error
	if err != nil {
		return domain.Category{}, err
	}
	var categoryResponse domain.Category
	err = cat.DB.Raw(`
		SELECT 

		cat.id,
		cat.category
		
		FROM
		categories cat
		WHERE
		cat.category = ?
	`, b).Scan(&categoryResponse).Error
	if err != nil {
		return domain.Category{}, err
	}
	return categoryResponse, nil
}

func (cat *categoryRepository)CheckCategory(current string)(bool,error){
	var response int
	err:=cat.DB.Raw("SELECT COUNT (*) FROM categories WHERE category= ?",current).Scan(&response).Error
	if err!=nil{
		return false,err
	}
	if response == 0{
		return false,err
	}
	return true,nil
}

func (cat *categoryRepository)UpdateCategory(current, new string)(domain.Category,error){

	// check database connection 

	if cat.DB == nil{
		return domain.Category{},errors.New("database connection failed while update category")
	}

	// update category
	err:=cat.DB.Exec("UPDATE categories SET category=? WHERE category=?",new,current).Error
	if err!=nil{
		return domain.Category{},err
	}
	// Retrieve updated category
	var updatedCat domain.Category
	err =cat.DB.First(&updatedCat,"category=?",new).Error
	if err!=nil{
		return domain.Category{},err
	}
	return updatedCat,nil
}

func (cat *categoryRepository)DeleteCategory(categoryId string)error{
	id,err:=strconv.Atoi(categoryId) 
	if err!=nil{
		return errors.New("string to int conversion failed")
	}
	deleteRes:=cat.DB.Exec("DELETE FROM categories WHERE id=?",id)
	if deleteRes.RowsAffected <1{
		return errors.New("No record exists with this id")
	}
	return nil
}