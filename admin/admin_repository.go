package admin

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/usecases"
)

type AdminModel struct {
	DB *sql.DB
}



func (a AdminModel) DBGetAdmin(adminForm models.Admin) (models.Admin, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var admin models.Admin

	row := a.DB.QueryRow(`SELECT admin_id,
								admin_name,
								admin_password
							FROM   admins
							WHERE  admin_name = $1 `, adminForm.Admin_name)

	row.Scan(&admin.Admin_id, &admin.Admin_name, &admin.Admin_password)
	log.Println("Output from database: ", admin)

	if admin.Admin_password != adminForm.Admin_password {
		return admin, errors.New("Incorrect Password")
	}

	log.Println("Output from database: ", admin)
	return admin, nil
}

func (a AdminModel) DBGetAdminByName(admin_name string) (models.Admin, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var admin models.Admin

	row := a.DB.QueryRow("SELECT admin_id, admin_name FROM admins WHERE admin_id=$1", admin_name)
	err := row.Scan(&admin.Admin_id, &admin.Admin_name)
	if err != nil {
		log.Println("DBGetAdminByName, error: ", err)
	}
	return admin, nil
}

func (a AdminModel) DBGetAdminById(admin_id string) (models.Admin, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var admin models.Admin
	row := a.DB.QueryRow("SELECT admin_id, admin_name, admin_password FROM admins WHERE admin_id = $1", admin_id)
	err := row.Scan(&admin.Admin_id, &admin.Admin_name, &admin.Admin_password)

	if err != nil {
		log.Println("Redirecting to admin... ", admin, err)
		return admin, err
	}
	return admin, nil
}





// ----------------------------User Database------------------------------

func (a AdminModel) DBGetUsers() ([]models.User, error) {

	file := usecases.Logger()
	log.SetOutput(file)
	var users []models.User

	rows, err := a.DB.Query("SELECT user_id, user_firstname, user_secondname, user_email, user_phone, user_status, user_gender, deleted_at FROM users")
	if err != nil {
		log.Println("Query wrote...", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.User_id, &user.User_firstname, &user.User_secondname, &user.User_email, &user.User_phone, &user.User_status, &user.User_gender, &user.Deleted_at)
		if err != nil {
			log.Println("User: ", user, "Error: ", err)
			return users, nil
		}
		log.Println("User: ", user)
		users = append(users, user)
	}

	log.Println("Processing user list...")
	return users, nil
}

func (a AdminModel) DBGetUser(user_id string) (models.User, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var user models.User
	log.Println("DBGetUser, user_id:", user_id)

	row := a.DB.QueryRow("SELECT user_id, user_firstname, user_secondname, user_email, user_phone, user_referral, user_status, user_gender, deleted_at FROM users WHERE user_id=$1", user_id)

	err := row.Scan(&user.User_id, &user.User_firstname, &user.User_secondname, &user.User_email, &user.User_phone, &user.User_referral, &user.User_status, &user.User_gender, &user.Deleted_at)
	if err != nil {
		log.Println("User: ", user, "Error: ", err)
		return user, err
	}

	log.Println("Database User: ", user)
	return user, nil
}

func (a AdminModel) DBGetUserStatus(user_id string) error {

	file := usecases.Logger()
	log.SetOutput(file)

	var user models.User

	row := a.DB.QueryRow("SELECT user_status FROM users WHERE user_id=$1", user_id)
	err := row.Scan(&user.User_status)
	if err != nil {
		log.Println(err)
		return err
	}
	var status bool
	if user.User_status {
		status = false
	} else {
		status = true
	}
	_, err = a.DB.Exec("UPDATE users SET user_status=$1 WHERE user_id=$2", status, user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a AdminModel) DBUpdatingUser(user_id string, newUser models.User) error {

	file := usecases.Logger()
	log.SetOutput(file)

	log.Println("Updating user:", user_id)
	result, err := a.DB.Exec("UPDATE users SET user_firstname=$1, user_secondname=$2, user_email=$3, user_phone=$4, user_referral=$5 WHERE user_id=$6", newUser.User_firstname, newUser.User_secondname, newUser.User_email, newUser.User_phone, newUser.User_referral, user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Updated result:", result)

	return nil
}

func (a AdminModel) DBDeleteUser(user_id string) error {

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("UPDATE users SET deleted_at=$1 WHERE user_id=$2", time.Now(), user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
}





// --------------------------------Product Database--------------------------

func (a AdminModel) DBGetAllProducts() ([]models.SampleProduct, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var products []models.SampleProduct

	rows, err := a.DB.Query("SELECT product_id, product_name, product_price, product_desc FROM products")
	if err != nil {
		log.Println(err)
		return products, err
	}
	defer rows.Close()
	for rows.Next() {
		var product models.SampleProduct
		err := rows.Scan(&product.Product_id, &product.Product_name, &product.Product_price, &product.Product_desc)
		if err != nil {
			log.Println("Error while getting products", err)
			return products, err
		}
		log.Println("DB Product:", product)
		products = append(products, product)
	}

	log.Println("Products:", products)
	return products, nil
}

func (a AdminModel) DBGetProduct(product_id string) (models.ListProduct, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var product models.ListProduct

	row := a.DB.QueryRow(`SELECT products.product_id,
								products.product_name,
								products.product_price,
								products.product_desc,
								products.product_brand_id,
								brands.brand_name,
								products.product_category_id,
								categories.category_name,
								products.product_subcategory_id,
								subcategories.subcategory_name,
								products.product_inventory_id,
								inventories.product_quantity,
								inventories.product_color
							FROM   products
								inner join brands
										ON products.product_brand_id = brands.brand_id
								inner join categories
										ON products.product_category_id = categories.category_id
								inner join subcategories
										ON products.product_subcategory_id = subcategories.subcategory_id
								inner join inventories
										ON products.product_inventory_id = inventories.inventory_id
							WHERE  products.product_id = $1`, product_id)
							
	err := row.Scan(&product.Product_id, &product.Product_name, &product.Product_price, &product.Product_desc, &product.Product_brand.Brand_id, &product.Product_brand.Brand_name, &product.Product_category.Category_id, &product.Product_category.Category_name, &product.Product_subcategory.Subcategory_id, &product.Product_subcategory.Subcategory_name, &product.Product_inventory.Inventory_id, &product.Product_inventory.Product_quantity, &product.Product_inventory.Product_color)
	if err != nil {
		log.Println("Error in DB", err)
		log.Println(err)
		return product, err
	}
	log.Println("Product in detail:", product)
	return product, nil
}

func (a AdminModel) DBGetProductColors(product_id string) ([]models.Inventories, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var colors []models.Inventories
	rows, err := a.DB.Query("SELECT * FROM colors WHERE product_id = $1", product_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var color models.Inventories
		err = rows.Scan(&color.Inventory_id, &color.Product_color, &color.Product_id)
		if err != nil {
			log.Println(err)
			return colors, err
		}
		log.Println("Color:", color)
		colors = append(colors, color)
	}
	return colors, nil
}

func (a AdminModel) DBAddProduct(product models.ListProduct) error {

	file := usecases.Logger()
	log.SetOutput(file)

	log.Println("In DBAddProduct..")

	var inventory_id int
	err := a.DB.QueryRow("INSERT INTO inventories (product_quantity, product_color, inventory_created_at) VALUES ($1, $2, $3) RETURNING inventory_id", product.Product_inventory.Product_quantity, product.Product_inventory.Product_color, time.Now()).Scan(&inventory_id)
	if err != nil {
		log.Println("Error in inventories")
		log.Println(err)
		return err
	}
	err = a.DB.QueryRow("INSERT INTO products (product_name, product_desc, product_price, product_brand_id, product_category_id, product_subcategory_id, product_inventory_id, product_created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING product_id", product.Product_name, product.Product_desc, product.Product_price, product.Product_brand.Brand_id, product.Product_category.Category_id, product.Product_subcategory.Subcategory_id, inventory_id, time.Now()).Scan(&product.Product_id)
	if err != nil {
		log.Println("Error in products")
		log.Println(err)
		return err
	}
	result, err := a.DB.Exec("INSERT INTO colors (color, product_id) VALUES ($1, $2)", product.Product_inventory.Product_color, product.Product_id)
	if err != nil {
		log.Println("Error in colors")
		log.Println(err)
		return err
	}
	log.Println("New Color:", result)
	return nil
}

func (a AdminModel) DBEditProduct(product models.ListProduct) error {

	file := usecases.Logger()
	log.SetOutput(file)

	log.Println("Name:", product.Product_name)
	log.Println("Desc:", product.Product_desc) 
	log.Println("Price:", product.Product_price)
	log.Println("Time:", time.Now())
	log.Println("Brand:", product.Product_brand.Brand_id)
	log.Println("Category:", product.Product_category.Category_id)
	log.Println("Sibcategory:", product.Product_subcategory.Subcategory_id) 
	log.Println("Quantity:", product.Product_inventory.Product_quantity)
	log.Println("Product_id:", product.Product_id)

	// row := a.DB.QueryRow("SELECT product_id, product_name, product_category_id, product_inventory_id FROM products WHERE product_id=$1", product.Product_id)
	// err := row.Scan(&product.Product_id, &product.Product_name, &product.Product_category.Category_id, &product.Product_inventory.Inventory_id)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(product)
	// _, _ = a.DB.Exec("UPDATE categories SET category_name=$1, category_desc=$2, category_updated_at=$3 WHERE category_id=$4", category.Category_name, category.Category_desc, time.Now(), category.Category_id)

	// result, err := a.DB.Exec("UPDATE inventories SET product_quantity=$1, inventory_updated_at=$2 WHERE inventory_id=$3", product.Product_inventory.Product_quantity, time.Now(), product.Product_inventory.Inventory_id)
	// if err != nil {
	// 	log.Println("Error in inventories")
	// 	log.Println(err)
	// 	return err
	// }
	// log.Println("Inventory Result:", result)

	result, err := a.DB.Exec("UPDATE products SET product_name=$1, product_desc=$2, product_price=$3, product_updated_at=$4, product_brand_id=$5, product_category_id=$6, product_subcategory_id=$7, WHERE product_id=$8", product.Product_name, product.Product_desc, product.Product_price, time.Now(), product.Product_brand.Brand_id, product.Product_category.Category_id, product.Product_subcategory.Subcategory_id, product.Product_id)
	if err != nil {
		log.Println(err)
		return err
	}

	// result, err = a.DB.QueryRow("SELECT  FROM products WHERE p")

	log.Println("Products Result:", result)

	result, err = a.DB.Exec("UPDATE inventories SET product_color=$1 WHERE invetory_id=$2", product.Product_inventory.Product_color, product.Product_inventory.Inventory_id)
	if err != nil {
		log.Println("Error in colors")
		log.Println(err)
		return err
	}
	log.Println("Color Results:", result)

	return nil
}





// ------------------Category Database----------------------

func (a AdminModel) DBGetAllCategories() ([]models.Categories, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var categories []models.Categories
	rows, err := a.DB.Query("SELECT * FROM categories WHERE category_deleted_at IS NULL")
	if err != nil {
		log.Println(err)
		return categories, err
	}
	defer rows.Close()
	for rows.Next() {
		var category models.Categories
		err = rows.Scan(&category.Category_id, &category.Category_name, &category.Category_desc, &category.Category_created_at, &category.Category_updated_at, &category.Category_deleted_at)
		if err != nil {
			log.Println(err)
			return categories, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (a AdminModel) DBAddCategory(category models.Categories) error {

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("INSERT INTO categories (category_name, category_desc, category_created_at) VALUES ($1, $2, $3)", category.Category_name, category.Category_desc, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("New Category:", result)
	return nil
}

func (a AdminModel) DBGetCategoryProducts(category_id string)([]models.SampleProduct, error){

	file := usecases.Logger()
	log.SetOutput(file)

	var products []models.SampleProduct
	rows, err := a.DB.Query("SELECT product_id, product_name, product_desc, product_price FROM products WHERE product_category_id=$1", category_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var product models.SampleProduct
		err = rows.Scan(&product.Product_id, &product.Product_name, &product.Product_desc, &product.Product_price)
		if err != nil {
			log.Println(err)
			return products, err
		}
		log.Println("Product:", product)
		products = append(products, product)
	}
	return products, nil
}

func (a AdminModel) DBGetCategory(category_id string)(models.Categories, error){

	file := usecases.Logger()
	log.SetOutput(file)

	var category models.Categories
	row := a.DB.QueryRow("SELECT category_id, category_name, category_desc FROM categories WHERE category_id=$1", category_id)
	err := row.Scan(&category.Category_id, &category.Category_name, &category.Category_desc)
	if err != nil {
		log.Println("Error in Get Category", err)
		return category, err
	}
	return category, nil
}

func (a AdminModel) DBEditCategory(category models.Categories) (error){

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("UPDATE categories SET category_name=$1, category_desc=$2, category_updated_at=$3 WHERE category_id=$4", category.Category_name, category.Category_desc, time.Now(), category.Category_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Result:", result)
	return nil
}

func (a AdminModel) DBDeleteCategory(category_id string) (error){

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("UPDATE categories SET category_deleted_at=$1 WHERE category_id=$2", time.Now(), category_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
} 




// ----------------Sub Category Database------------------------

func (a AdminModel) DBGetAllSubcategories() ([]models.Subcategories, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var subcategories []models.Subcategories
	rows, err := a.DB.Query("SELECT * FROM subcategories WHERE subcategory_deleted_at IS NULL")
	if err != nil {
		log.Println(err)
		return subcategories, err
	}
	defer rows.Close()
	for rows.Next() {
		var subcategory models.Subcategories
		err = rows.Scan(&subcategory.Subcategory_id, &subcategory.Subcategory_name, &subcategory.Subcategory_desc, &subcategory.Subcategory_created_at, &subcategory.Subcategory_updated_at, &subcategory.Subcategory_deleted_at)
		if err != nil {
			log.Println(err)
			return subcategories, err
		}
		subcategories = append(subcategories, subcategory)
	}
	return subcategories, nil
}

func (a AdminModel) DBAddSubcategory(subcategory models.Subcategories) error {

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("INSERT INTO subcategories (subcategory_name, subcategory_desc, subcategory_created_at) VALUES ($1, $2, $3)", subcategory.Subcategory_name, subcategory.Subcategory_desc, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("New Subcategory:", result)

	return nil
}

func (a AdminModel) DBGetSubcategoryProducts(subcategory_id string)([]models.SampleProduct, error){

	file := usecases.Logger()
	log.SetOutput(file)

	var products []models.SampleProduct
	rows, err := a.DB.Query("SELECT product_id, product_name, product_desc, product_price FROM products WHERE product_subcategory_id=$1", subcategory_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var product models.SampleProduct
		err = rows.Scan(&product.Product_id, &product.Product_name, &product.Product_desc, &product.Product_price)
		if err != nil {
			log.Println(err)
			return products, err
		}
		log.Println("Product:", product)
		products = append(products, product)
	}
	return products, nil
}

func (a AdminModel) DBGetSubcategory(subcategory_id string)(models.Subcategories, error){

	file := usecases.Logger()
	log.SetOutput(file)

	var subcategory models.Subcategories
	row := a.DB.QueryRow("SELECT subcategory_id, subcategory_name, subcategory_desc FROM subcategories WHERE subcategory_id=$1", subcategory_id)
	err := row.Scan(&subcategory.Subcategory_id, &subcategory.Subcategory_name, &subcategory.Subcategory_desc)
	if err != nil {
		log.Println("Error in Get Subcategory", err)
		return subcategory, err
	}
	return subcategory, nil
}

func (a AdminModel) DBEditSubcategory(subcategory models.Subcategories) (error){

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("UPDATE subcategories SET subcategory_name=$1, subcategory_desc=$2, subcategory_updated_at=$3 WHERE subcategory_id=$4", subcategory.Subcategory_name, subcategory.Subcategory_desc, time.Now(), subcategory.Subcategory_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Result:", result)
	return nil
}

func (a AdminModel) DBDeleteSubcategory(subcategory_id string) (error){

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("UPDATE subcategories SET subcategory_deleted_at=$1 WHERE subcategory_id=$2", time.Now(), subcategory_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
} 




// ----------------Brand Database------------------------

func (a AdminModel) DBGetAllBrands() ([]models.Brands, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var brands []models.Brands
	rows, err := a.DB.Query("SELECT * FROM brands WHERE brand_deleted_at IS NULL")
	if err != nil {
		log.Println(err)
		return brands, err
	}
	defer rows.Close()
	for rows.Next() {
		var brand models.Brands
		err = rows.Scan(&brand.Brand_id, &brand.Brand_name, &brand.Brand_desc, &brand.Brand_created_at, &brand.Brand_updated_at, &brand.Brand_deleted_at)
		if err != nil {
			log.Println(err)
			return brands, err
		}
		brands = append(brands, brand)
	}
	log.Println("DBGetAllProducts..")
	return brands, nil
}

func (a AdminModel) DBAddBrand(brand models.Brands) error {

	file := usecases.Logger()
	log.SetOutput(file)

	log.Println("New Brand is inserting:", brand)

	result, err := a.DB.Exec("INSERT INTO brands (brand_name, brand_desc, brand_created_at) VALUES ($1, $2, $3)", brand.Brand_name, brand.Brand_desc, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("New Brand:", result)
	return nil
}

func (a AdminModel) DBGetBrandProducts(brand_id string)([]models.SampleProduct, error){

	file := usecases.Logger()
	log.SetOutput(file)

	var products []models.SampleProduct
	rows, err := a.DB.Query("SELECT product_id, product_name, product_desc, product_price FROM products WHERE product_brand_id=$1", brand_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var product models.SampleProduct
		err = rows.Scan(&product.Product_id, &product.Product_name, &product.Product_desc, &product.Product_price)
		if err != nil {
			log.Println(err)
			return products, err
		}
		log.Println("Product:", product)
		products = append(products, product)
	}
	return products, nil
}

func (a AdminModel) DBGetBrand(brand_id string)(models.Brands, error){

	file := usecases.Logger()
	log.SetOutput(file)

	var brand models.Brands
	row := a.DB.QueryRow("SELECT brand_id, brand_name, brand_desc FROM brands WHERE brand_id=$1", brand_id)
	err := row.Scan(&brand.Brand_id, &brand.Brand_name, &brand.Brand_desc)
	if err != nil {
		log.Println("Error in Get Category", err)
		return brand, err
	}
	return brand, nil
}

func (a AdminModel) DBEditBrand(brand models.Brands) (error){

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("UPDATE brands SET brand_name=$1, brand_desc=$2, brand_updated_at=$3 WHERE brand_id=$4", brand.Brand_name, brand.Brand_desc, time.Now(), brand.Brand_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Result:", result)
	return nil
}

func (a AdminModel) DBDeleteBrand(brand_id string) (error){

	file := usecases.Logger()
	log.SetOutput(file)

	result, err := a.DB.Exec("UPDATE brands SET brand_deleted_at=$1 WHERE brand_id=$2", time.Now(), brand_id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
} 
