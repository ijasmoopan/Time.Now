package admin

import (
	"errors"
	"fmt"

	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/repository"
)

func ValidateAdmin(adminForm models.Admin) (models.Admin, error) {

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	var admin models.Admin

	if result := db.First(&admin, "admin_name = ? ", adminForm.Admin_name); result.Error != nil {
		return admin, result.Error
	}
	if admin.Admin_password != adminForm.Admin_password {
		return admin, errors.New("Incorrect Password")
	}
	return admin, nil
}

func DBGetUsers()([]models.User, []models.User, error){

	var users []models.User
	var allUsers []models.User

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	if result := db.Find(&users); result.Error != nil {
		return users, allUsers, result.Error
	}
	if result := db.Unscoped().Find(&allUsers); result.Error != nil {
		return users, allUsers, result.Error
	}
	fmt.Println("Processing user list...")
	return users, allUsers, nil
}

func DBGetUser(user_id string) (models.User, error) {

	var user models.User
	db := repository.ConnectDB()
	defer repository.CloseDB(db)
	if result := db.First(&user, "user_id = ? ", user_id); result.Error != nil {
		return user, result.Error
	} 
	fmt.Println("Database User: ", user)
	return user, nil
}

func DBGetUserStatus(user_id string)(error){

	var user models.User
	db := repository.ConnectDB()
	defer repository.CloseDB(db)
	fmt.Println("Changing user status, User ID: ", user_id)
	db.First(&user, "user_id = ?", user_id)
	if user.User_status {
		user.User_status = false
	} else {
		user.User_status = true
	}
	fmt.Println("User status: ", user.User_status)
	db.Save(&user)
	return nil
}

func UpdatingUser(user_id string, newUser models.User) error {

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	var user models.User
	db.First(&user, "user_id = ?", user_id)

	user.User_firstname = newUser.User_firstname
	user.User_secondname = newUser.User_secondname
	user.User_email = newUser.User_email
	user.User_phone = newUser.User_phone
	user.User_referral = newUser.User_referral

	db.Save(&user)

	return nil
}
func DBDeleteUser(user_id string) error {

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	var user models.User
	if result := db.Delete(&user, "user_id = ?", user_id); result.Error != nil {
		return result.Error
	}
	return nil
}



// --------------------------------Product Side--------------------------

func DBGetAllProducts() ([]models.Product, error){

	var products []models.Product
	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	if result := db.Find(&products); result.Error != nil {
		return products, result.Error
	}
	return products, nil
}

func DBGetProduct(product_id string) (models.Product, error){

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	var product models.Product
	if result := db.First(&product, "product_id = ?", product_id); result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

func DBUpdateProduct(product_id string, editProduct models.Product) error {

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	var product models.Product
	db.First(&product, "product_id", product_id)

	product.Product_name = editProduct.Product_name
	product.Product_color = editProduct.Product_color
	product.Product_desc = editProduct.Product_desc
	product.Product_price = editProduct.Product_price
	product.Product_quantity = editProduct.Product_quantity
	
	db.Save(&product)
	
	return nil
}