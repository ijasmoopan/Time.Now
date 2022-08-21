package admin

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/ijasmoopan/Time.Now/models"
	// "github.com/ijasmoopan/Time.Now/usecases"
)

// Model for implementing db connection.
type Model struct {
	DB *sql.DB
}

// DBGetAdmin for accessing admin.
func (admin Model) DBGetAdmin(adminForm models.Admin) (models.Admin, error) {

	var adminData models.Admin
	row := admin.DB.QueryRow(`SELECT admin_id,
									admin_name,
									admin_password
								FROM   admins
								WHERE  admin_name = $1 `, adminForm.Name)
	err := row.Scan(&adminData.ID,
		&adminData.Name,
		&adminData.Password)
	if err != nil {
		log.Println(err)
		return adminData, err
	}
	if adminData.Password != adminForm.Password {
		return adminData, errors.New("Incorrect Password")
	}
	return adminData, nil
}

// DBGetAdminByID for accessing admin by id.
func (admin Model) DBGetAdminByID(ID string) (models.Admin, error) {

	var adminData models.Admin
	row := admin.DB.QueryRow(`SELECT admin_id, 
								admin_name, 
								admin_password 
							FROM admins 
							WHERE admin_id = $1`, ID)
	err := row.Scan(&adminData.ID,
		&adminData.Name,
		&adminData.Password)
	if err != nil {
		log.Println(err)
		return adminData, err
	}
	return adminData, nil
}

// ----------------------------User Database------------------------------

// DBGetUsers for giving users to admin.
func (admin Model) DBGetUsers(request models.UserRequest) ([]models.User, error) {

	var users []models.User
	var param []interface{}
	var rows *sql.Rows
	var err error
	var count = 1

	sqlString := `SELECT user_id, 
						user_firstname, 
						user_secondname, 
						user_email, 
						user_phone, 
						user_status, 
						user_gender
					FROM users 
					WHERE deleted_at IS NULL`

	if request.UserID != "" {
		param = append(param, request.UserID)
		sqlString = fmt.Sprint(sqlString, ` AND user_id = $`, count)
		count++
	}
	if request.Email != "" {
		param = append(param, request.Email)
		sqlString = fmt.Sprint(sqlString, ` AND user_email = $`, count)
		count++
	}
	if request.Gender != "" {
		param = append(param, request.Gender)
		sqlString = fmt.Sprint(sqlString, ` AND user_gender = $`, count)
		count++
	}
	if request.Status != "" {
		param = append(param, request.Status)
		sqlString = fmt.Sprint(sqlString, ` AND user_status = $`, count)
		count++
	}
	rows, err = admin.DB.Query(sqlString, param...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID,
			&user.FirstName,
			&user.SecondName,
			&user.Email,
			&user.Phone,
			&user.Status,
			&user.Gender,
		)
		if err != nil {
			log.Println("User: ", user, "Error: ", err)
			return users, nil
		}
		log.Println("User: ", user)
		users = append(users, user)
	}
	return users, nil
}

// DBUpdateUserStatus for accessing users by it is active or not.
func (admin Model) DBUpdateUserStatus(userID int) error {

	var user models.User

	row := admin.DB.QueryRow(`SELECT user_status 
								FROM users 
								WHERE user_id = $1`, userID)

	err := row.Scan(&user.Status)
	if err != nil {
		log.Println(err)
		return err
	}
	var status bool
	if user.Status {
		status = false
	} else {
		status = true
	}
	_, err = admin.DB.Exec(`UPDATE users 
							SET user_status = $1 
							WHERE user_id = $2`, status,
		userID)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBUpdateUser for updating user details.
func (admin Model) DBUpdateUser(user models.User) error {

	_, err := admin.DB.Exec(`UPDATE users 
									SET user_firstname = $1, 
										user_secondname = $2, 
										user_email = $3, 
										user_phone = $4 
									WHERE user_id = $5;`, user.FirstName,
		user.SecondName,
		user.Email,
		user.Phone,
		user.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBDeleteUser for deleting a user.
func (admin Model) DBDeleteUser(userID int) error {

	result, err := admin.DB.Exec(`UPDATE users 
									SET deleted_at = $1 
									WHERE user_id = $2`, time.Now(),
		userID,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
}

// --------------------------------Product Database--------------------------

// DBUpdateProductStatus method for updating product status from database.
func (admin Model) DBUpdateProductStatus(product models.Product) error {

	row := admin.DB.QueryRow(`SELECT product_status
								WHERE product_id = $1`, product.ID)
	err := row.Scan(&product.Status)
	if err != nil {
		log.Println(err)
		return err
	}
	var status bool
	if product.Status {
		status = false
	} else {
		status = true
	}
	_, err = admin.DB.Exec(`UPDATE products 
								SET product_status = $1
								WHERE product_id = $2`, status,
		product.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBDeleteProducts method for deleting product from database.
func (admin Model) DBDeleteProducts(product models.ProductDeleteRequest) error {

	if product.ID != nil {
		_, err := admin.DB.Exec(`UPDATE products
									SET product_deleted_at = $1
									WHERE product_id = $2`, time.Now(),
			product.ID)
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = admin.DB.Exec(`UPDATE inventories 
									SET inventory_deleted_at = $1
									WHERE product_id = $2`, time.Now(),
			product.ID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if product.ImageID != nil {
		_, err := admin.DB.Exec(`DELETE FROM images 
									WHERE image_id = $1;`, product.ImageID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if product.InventoryID != nil {
		_, err := admin.DB.Exec(`UPDATE inventories 
									SET inventory_deleted_at = $1
									WHERE inventory_id = $2`, time.Now(),
			product.InventoryID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// ----------------------- DB Get Product Full Details --------------------------

// DBGetProducts for accessing products with respect to reqeust params.
func (admin Model) DBGetProducts(request models.AdminProductRequest) (map[int]models.ProductWithInventory, int, error) {

	var products = make(map[int]models.ProductWithInventory)
	filter := make([]interface{}, 0, 10)
	count := 1
	var sqlCondition string
	var rows *sql.Rows
	var err error

	sqlString := `WITH cte_getproducts (product_id, 
										product_name, 
										product_price,
										product_desc, 
										product_status, 
										product_deleted_at,
										product_brand_id,
										product_category_id,
										product_subcategory_id,
										image_id,
										product_image,
										category_name,
										category_desc,
										subcategory_name,
										subcategory_desc,
										brand_name,
										brand_desc,
										offer_id,
										offer_name,
										offer,
										offer_from,
										offer_to,
										offer_status,
										offer_price
										) 
								AS (SELECT p.product_id,
											p.product_name,
											p.product_price,
											p.product_desc,
											p.product_status,
											p.product_deleted_at,
											p.product_brand_id,
											p.product_category_id,
											p.product_subcategory_id,
											i.image_id,
											i.product_image,
											c.category_name,
											c.category_desc,
											s.subcategory_name,
											s.subcategory_desc,
											b.brand_name,
											b.brand_desc,
											o.id,
											o.name,
											o.offer,
											o.offer_from,
											o.offer_to,
											CASE 
												WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN true
												ELSE false
											END AS offer_status,
											CASE 
												WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN (p.product_price - ((p.product_price * o.offer)/100))::FLOAT
												ELSE NULL
											END AS offer_price
								FROM   products p
									INNER JOIN categories c
										ON p.product_category_id = c.category_id
									INNER JOIN subcategories s
										ON p.product_subcategory_id = s.subcategory_id
									INNER JOIN brands b
										ON p.product_brand_id = b.brand_id
									LEFT JOIN images i
										ON p.product_id = i.product_id

									LEFT JOIN category_offers o
										ON o.category_id = p.product_category_id
								WHERE  p.product_deleted_at IS NULL
									ORDER BY offer)

								SELECT * 
									FROM cte_getproducts
									WHERE product_deleted_at IS NULL`

	// sqlCount := `WITH cte_totalproducts (count)
	// 						AS (SELECT COUNT(1)
	// 					FROM   products p
	// 						INNER JOIN categories c
	// 							ON p.product_category_id = c.category_id
	// 						INNER JOIN subcategories s
	// 							ON p.product_subcategory_id = s.subcategory_id
	// 						INNER JOIN brands b
	// 							ON p.product_brand_id = b.brand_id
	// 						LEFT JOIN images i
	// 							ON p.product_id = i.product_id

	// 						LEFT JOIN category_offers o
	// 							ON o.category_id = p.product_category_id
	// 					WHERE  p.product_deleted_at IS NULL
	// 						ORDER BY offer) 
							
	// 			SELECT * 
	// 				FROM cte_totalproducts
	// 				WHERE product_deleted_at IS NULL` 

	// sqlCount = `SELECT COUNT(1) 
	// 				FROM products
	// 				WHERE product_deleted_at IS NULL`

	// -------------------------------------------------------------------------

	if request.Product != "" {
		splitProduct := strings.Split(request.Product, " ")
		for idx, value := range splitProduct {
			if idx == 0 {
				filter = append(filter, fmt.Sprint("%", value, "%"))
				sqlCondition = fmt.Sprint(` AND product_name ILIKE $`, count)
				// sqlCount = fmt.Sprint(sqlCount, ` AND product_name ILIKE $`, count)
				count++
			} else {
				filter = append(filter, fmt.Sprint("%", value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` OR product_name ILIKE $`, count)
				count++
			}
		}
	}
	if request.Category != "" {
		splitCategory := strings.Split(request.Category, " ")
		for idx, value := range splitCategory {
			if idx == 0 {
				filter = append(filter, fmt.Sprint("%", value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` AND category_name ILIKE $`, count)
				count++
			} else {
				filter = append(filter, fmt.Sprint("%", value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` OR category_name ILIKE $`, count)
				count++
			}
		}
	}
	if request.Subcategory != "" {
		splitSubcategory := strings.Split(request.Subcategory, " ")
		for idx, value := range splitSubcategory {
			if idx == 0 {
				filter = append(filter, fmt.Sprint("%", value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` AND subcategory_name ILIKE $`, count)
				count++
			} else {
				filter = append(filter, fmt.Sprint("%", value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` OR subcategory_name ILIKE $`, count)
				count++
			}
		}
	}
	if request.Brand != "" {
		splitBrand := strings.Split(request.Brand, " ")
		for idx, value := range splitBrand {
			if idx == 0 {
				filter = append(filter, fmt.Sprint("%", value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` AND brand_name ILIKE $`, count)
				count++
			} else {
				filter = append(filter, fmt.Sprint("%", value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` OR brand_name ILIKE $`, count)
				count++
			}
		}
	}
	if request.PriceMin != "" {
		filter = append(filter, request.PriceMin)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND product_price >= $`, count)
		count++
	}
	if request.PriceMax != "" {
		filter = append(filter, request.PriceMax)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND product_price <= $`, count)
		count++
	}
	if request.Status != "" {
		filter = append(filter, request.Status)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND product_status = $`, count)
		count++
	}

	var limit = 5
	sqlCondition = fmt.Sprint(sqlCondition, ` ORDER BY product_id`)
	if request.Page != "" {
		page, _ := strconv.Atoi(request.Page)
		page = (page - 1) * 5
		filter = append(filter, limit, page)
		sqlCondition = fmt.Sprint(sqlCondition, ` LIMIT $`, count)
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` OFFSET $`, count)
		count++
	} else {
		filter = append(filter, limit, 0)
		sqlCondition = fmt.Sprint(sqlCondition, ` LIMIT $`, count)
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` OFFSET $`, count)
		count++
	}

	// log.Println("sqlCondition:", sqlCondition)
	// log.Println("filter:", filter)

	var totalProducts = 35
	// row := admin.DB.QueryRow(fmt.Sprint(sqlCount, sqlCondition), filter...)
	// err = row.Scan(&totalProducts)
	// if err != nil {
	// 	log.Println(err)
	// 	return products, totalProducts, err
	// }

	rows, err = admin.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter...)

	if err != nil {
		log.Println(err)
		return products, totalProducts, err
	}
	defer rows.Close()
	for rows.Next() {
		var product models.ProductWithInventory
		var deletedAt *time.Time
		var imageID, offerID, offer sql.NullInt32
		var image, offerName sql.NullString
		var offerFrom, offerTo sql.NullTime
		var offerStatus bool
		var offerPrice sql.NullFloat64

		err := rows.Scan(&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.Status,
			&deletedAt,
			&product.Brand.ID,
			&product.Category.ID,
			&product.Subcategory.ID,
			&imageID,
			&image,
			&product.Category.Name,
			&product.Category.Description,
			&product.Subcategory.Name,
			&product.Subcategory.Description,
			&product.Brand.Name,
			&product.Brand.Description,
			&offerID,
			&offerName,
			&offer,
			&offerFrom,
			&offerTo,
			&offerStatus,
			&offerPrice,
		)
		if imageID.Valid {
			product.Image.ID = int(imageID.Int32)
			product.Image.ProductID = product.ID
		}
		if image.Valid {
			product.Image.Image = image.String
		}
		// if offerStatus.Valid {
		if offerStatus != false {
			log.Println("Offer status is ", offerStatus)
			product.Offer.Status = offerStatus

			if offerID.Valid {
				product.Offer.ID = int(offerID.Int32)
			}
			if offerName.Valid {
				product.Offer.Name = offerName.String
			}
			if offer.Valid {
				product.Offer.Offer = int(offer.Int32)
			}
			if offerFrom.Valid {
				product.Offer.From = offerFrom.Time.Format("01-02-2006")
			}
			if offerTo.Valid {
				product.Offer.To = offerTo.Time.Format("01-02-2006")
			}
			if offerPrice.Valid {
				product.OfferPrice = &offerPrice.Float64
			}
			product.Offer.Category.ID = product.Category.ID
			product.Offer.Category.Name = product.Category.Name
			product.Offer.Category.Description = product.Category.Description
		}
		if err != nil {
			log.Println("Error while getting products", err)
			return products, totalProducts, err
		}
		// log.Println("Offer:", product.Offer)
		// log.Println("Offer Price:", product.OfferPrice)
		// log.Println("Product half:", product)
		// log.Println("Product ID:", product.ID)
		// log.Println("Products Map:", products)

		var colors []models.Colors
		var inventories []models.Inventories
		count2 := 1
		var filter2 []interface{}
		var sqlCondition2 string
		filter2 = append(filter2, product.ID)
		sqlString2 := `SELECT i.inventory_id, 
							i.product_id, 
							c.color_id,
							c.color, 
							i.product_quantity 
						FROM inventories i
						INNER JOIN colors c
							ON i.color_id = c.color_id
						WHERE i.inventory_deleted_at IS NULL
							AND i.product_id = $1`

		if request.Quantity != "" {
			count2++
			filter2 = append(filter2, request.Quantity)
			sqlCondition2 = fmt.Sprint(sqlCondition2, ` AND i.product_quantity <= $`, count2)
		}
		if request.Color != "" {
			for idx, value := range request.Color {
				if idx == 0 {
					count2++
					filter2 = append(filter2, fmt.Sprint("%", value, "%"))
					sqlCondition2 = fmt.Sprint(sqlCondition2, ` AND c.color ILIKE $`, count2)
				} else {
					count2++
					filter2 = append(filter2, fmt.Sprint("%", value, "%"))
					sqlCondition2 = fmt.Sprint(sqlCondition2, ` OR c.color ILIKE $`, count2)
				}
			}
		}
		// log.Println("SQL Condition:", sqlCondition2)
		// log.Println("Filter 2:", filter2)
		// log.Println("Product ID:", product.ID)
		inventoryRows, err2 := admin.DB.Query(fmt.Sprint(sqlString2, sqlCondition2), filter2...)
		// log.Println("InventoryRows:", inventoryRows)
		if err2 == sql.ErrNoRows {
			log.Println("No rows returned", err2)
		}
		if err2 != nil {
			log.Println("Error2:", err2)
			return nil, totalProducts, err2
		}
		defer inventoryRows.Close()
		if inventoryRows.Next() {
			var inventory models.Inventories
			var color models.Colors
			err = inventoryRows.Scan(&inventory.ID,
				&inventory.ProductID,
				&inventory.Color.ID,
				&inventory.Color.Color,
				&inventory.Quantity)
			if err != nil {
				// log.Println("for loop:", err)
				return products, totalProducts, err
			}
			// log.Println("Color:", inventory)
			color.ID = inventory.Color.ID
			color.Color = inventory.Color.Color
			colors = append(colors, color)
			inventories = append(inventories, inventory)
			for inventoryRows.Next() {
				err = inventoryRows.Scan(&inventory.ID,
					&inventory.ProductID,
					&inventory.Color.ID,
					&inventory.Color.Color,
					&inventory.Quantity)
				if err != nil {
					log.Println("for loop:", err)
					return products, totalProducts, err
				}
				color.ID = inventory.Color.ID
				color.Color = inventory.Color.Color
				colors = append(colors, color)
				inventories = append(inventories, inventory)
			}
		} else {
			// rows empty
			log.Println("No colors")
			continue
		}
		product.Inventory = inventories
		product.Color = colors

		if _, ok := products[product.ID]; !ok {
			products[product.ID] = product
		}
	}
	return products, totalProducts, nil
}

// DBAddProducts for adding product to database
func (admin Model) DBAddProducts(product models.AddProduct) error {

	row := admin.DB.QueryRow(`INSERT INTO products (product_name,
												product_price,
												product_desc,
												product_created_at,
												product_brand_id,
												product_category_id,
												product_subcategory_id,
												product_status)
										VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING product_id;`, product.Name,
		product.Price,
		product.Description,
		time.Now(),
		product.BrandID,
		product.CategoryID,
		product.SubcategoryID,
		true)
	err := row.Scan(&product.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("ProductID:", product.ID)

	inventories := strings.Split(product.Inventories, ",")

	for i := 0; i < len(inventories); i++ {
		var colorID int
		row := admin.DB.QueryRow(`INSERT INTO colors (color)
										VALUES ($1)
										RETURNING color_id;`, inventories[i])
		err = row.Scan(&colorID)
		if err != nil {
			log.Println(err)
			return err
		}
		i++
		_, err = admin.DB.Exec(`INSERT INTO inventories (product_id,
														color_id,
														product_quantity,
														inventory_created_at)
												VALUES ($1, $2, $3, $4);`, product.ID,
						colorID,
						inventories[i],
						time.Now(),
						)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	log.Println("Before, Image:", product.Image)
	_, err = admin.DB.Exec(`INSERT INTO images (product_id,
												product_image)
										VALUES ($1, $2);`, product.ID,
		product.Image)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("After, Image:", product.Image)
	return nil
}

// DBUpdateProducts method for updating product from database.
func (admin Model) DBUpdateProducts(product models.ProductWithInventory) error {

	log.Println("Product:", product)
	_, err := admin.DB.Exec(`UPDATE products 
									SET product_name = $1, 
										product_desc = $2, 
										product_price = $3, 
										product_updated_at = $4, 
										product_brand_id = $5, 
										product_category_id = $6, 
										product_subcategory_id = $7 
									WHERE product_id = $8; `, product.Name,
		product.Description,
		product.Price,
		time.Now(),
		product.Brand.ID,
		product.Category.ID,
		product.Subcategory.ID,
		product.ID)
	if err != nil {
		log.Println("Updating products,", err)
		return err
	}
	log.Println("Products table updation completed")
	_, err = admin.DB.Exec(`UPDATE images  
								SET product_image = $1
								WHERE image_id = $2`, product.Image.Image,
		product.Image.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, value := range product.Inventory {

		_, err = admin.DB.Exec(`UPDATE inventories 
									SET color_id = $1,
										product_quantity = $2,
										inventory_updated_at = $3
									WHERE inventory_deleted_at IS NULL
										AND inventory_id = $4`, value.Color.ID,
			value.Quantity,
			time.Now(),
			value.ID)
		if err != nil {
			log.Println("Updating inventories, ", err)
			return err
		}

	}

	return nil
}

// ------------------Category Database----------------------

// DBGetCategories method for accessing categories from database.
func (admin Model) DBGetCategories(request models.CategoryRequest) ([]models.Categories, error) {

	var categories []models.Categories
	var rows *sql.Rows
	var err error
	var filter interface{}
	var sqlCondition string

	sqlString := `SELECT category_id,
						category_name,
						category_desc
					FROM categories 
					WHERE category_deleted_at IS NULL `

	if request.CategoryID != nil {
		filter = request.CategoryID
		sqlCondition = `AND category_id = $1`
	}
	if request.CategoryName != nil {
		filter = request.CategoryName
		sqlCondition = `AND category_name = $1`
	}
	if filter != nil {
		rows, err = admin.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter)
	} else {
		rows, err = admin.DB.Query(sqlString)
	}
	if err != nil {
		log.Println(err)
		return categories, err
	}
	defer rows.Close()
	for rows.Next() {
		var category models.Categories
		err = rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			log.Println(err)
			return categories, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// DBAddCategory method for adding new category to database.
func (admin Model) DBAddCategory(category models.Categories) error {

	_, err := admin.DB.Exec(`INSERT INTO categories (category_name, 
													category_desc, 
													category_created_at) 
											VALUES ($1, $2, $3)`, category.Name,
		category.Description,
		time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBUpdateCategory method for updating category details.
func (admin Model) DBUpdateCategory(category models.Categories) error {

	_, err := admin.DB.Exec(`UPDATE categories 
								SET category_name = $1, 
									category_desc = $2, 
									category_updated_at = $3 
								WHERE category_id = $4`, category.Name,
		category.Description,
		time.Now(),
		category.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBDeleteCategory method for deleting category from database.
func (admin Model) DBDeleteCategory(categoryID int) error {

	_, err := admin.DB.Exec(`UPDATE categories 
								SET category_deleted_at = $1 
								WHERE category_id = $2`, time.Now(),
		categoryID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ----------------Sub Category Database------------------------

// DBGetSubcategories method for accessing subcategories.
func (admin Model) DBGetSubcategories(request models.SubcategoryRequest) ([]models.Subcategories, error) {

	var subcategories []models.Subcategories
	var rows *sql.Rows
	var err error
	var filter interface{}
	var sqlCondition string

	sqlString := `SELECT subcategory_id,
						subcategory_name,
						subcategory_desc
					 FROM subcategories 
					 WHERE subcategory_deleted_at IS NULL `
	if request.SubcategoryID != nil {
		filter = request.SubcategoryID
		sqlCondition = ` AND subcategory_id = $1`
	}
	if request.SubcategoryName != nil {
		filter = request.SubcategoryName
		sqlCondition = ` AND subcategory_name = $1`
	}
	if filter != nil {
		rows, err = admin.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter)
	} else {
		rows, err = admin.DB.Query(sqlString)
	}
	if err != nil {
		log.Println(err)
		return subcategories, err
	}
	defer rows.Close()
	for rows.Next() {
		var subcategory models.Subcategories
		err = rows.Scan(&subcategory.ID, &subcategory.Name, &subcategory.Description)
		if err != nil {
			log.Println(err)
			return subcategories, err
		}
		subcategories = append(subcategories, subcategory)
	}
	return subcategories, nil
}

// DBAddSubcategory method adding subcategory to database.
func (admin Model) DBAddSubcategory(subcategory models.Subcategories) error {

	_, err := admin.DB.Exec(`INSERT INTO subcategories (subcategory_name, 
														subcategory_desc, 
														subcategory_created_at) 
										VALUES ($1, $2, $3)`, subcategory.Name,
		subcategory.Description,
		time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBUpdateSubcategory method for update category from database.
func (admin Model) DBUpdateSubcategory(subcategory models.Subcategories) error {

	_, err := admin.DB.Exec(`UPDATE subcategories 
								SET subcategory_name = $1, 
									subcategory_desc = $2, 
									subcategory_updated_at = $3 
								WHERE subcategory_id = $4`, subcategory.Name,
		subcategory.Description,
		time.Now(),
		subcategory.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBDeleteSubcategory method for deleting category from database.
func (admin Model) DBDeleteSubcategory(subcategoryID int) error {

	_, err := admin.DB.Exec(`UPDATE subcategories 
								SET subcategory_deleted_at = $1 
								WHERE subcategory_id = $2`, time.Now(), subcategoryID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ----------------Brand Database------------------------

// DBGetBrands method for accessing brands from database.
func (admin Model) DBGetBrands(request models.BrandRequest) ([]models.Brands, error) {

	var brands []models.Brands
	var rows *sql.Rows
	var err error
	var filter interface{}
	var sqlCondition string
	sqlString := `SELECT brand_id, 
						brand_name, 
						brand_desc 
					FROM brands 
					WHERE brand_deleted_at IS NULL `
	if request.BrandID != nil {
		filter = request.BrandID
		sqlCondition = ` AND brand_id = $1`
	}
	if request.BrandName != nil {
		filter = request.BrandName
		sqlCondition = ` AND brand_name = $1`
	}
	if filter != nil {
		rows, err = admin.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter)
	} else {
		rows, err = admin.DB.Query(sqlString)
	}
	if err != nil {
		log.Println(err)
		return brands, err
	}
	defer rows.Close()
	for rows.Next() {
		var brand models.Brands
		err = rows.Scan(&brand.ID, &brand.Name, &brand.Description)
		if err != nil {
			log.Println(err)
			return brands, err
		}
		brands = append(brands, brand)
	}
	return brands, nil
}

// DBAddBrand method for adding brand to database.
func (admin Model) DBAddBrand(brand models.Brands) error {

	_, err := admin.DB.Exec(`INSERT INTO brands (brand_name, 
													brand_desc, 
													brand_created_at) 
										VALUES ($1, $2, $3)`, brand.Name,
		brand.Description,
		time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBUpdateBrand method for updating brand from database.
func (admin Model) DBUpdateBrand(brand models.Brands) error {

	_, err := admin.DB.Exec(`UPDATE brands 
								SET brand_name = $1, 
									brand_desc = $2, 
									brand_updated_at = $3 
								WHERE brand_id = $4`, brand.Name,
		brand.Description,
		time.Now(),
		brand.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBDeleteBrand method for deleting brand from database.
func (admin Model) DBDeleteBrand(brandID int) error {

	_, err := admin.DB.Exec(`UPDATE brands 
								SET brand_deleted_at = $1 
								WHERE brand_id = $2`, time.Now(),
		brandID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ----------------Color Database------------------------

// DBGetColors method for accessing colors from database.
func (admin Model) DBGetColors(request models.ColorRequest) ([]models.Colors, error) {

	var colors []models.Colors
	var rows *sql.Rows
	var err error
	var filter interface{}
	var sqlCondition string
	sqlString := `SELECT color_id, 
						color 
					FROM colors 
					WHERE color_deleted_at IS NULL `
	if request.ColorID != nil {
		filter = request.ColorID
		sqlCondition = ` AND color_id = $1`
	}
	if request.Color != nil {
		filter = request.Color
		sqlCondition = ` AND color = $1`
	}
	if filter != nil {
		rows, err = admin.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter)
	} else {
		rows, err = admin.DB.Query(sqlString)
	}
	if err != nil {
		log.Println(err)
		return colors, err
	}
	defer rows.Close()
	for rows.Next() {
		var color models.Colors
		err = rows.Scan(&color.ID, &color.Color)
		if err != nil {
			log.Println(err)
			return colors, err
		}
		colors = append(colors, color)
	}
	return colors, nil
}

// DBAddColor method for adding color to database.
func (admin Model) DBAddColor(color models.Colors) error {

	_, err := admin.DB.Exec(`INSERT INTO colors (color, 
												color_created_at) 
										VALUES ($1, $2)`, color.Color,
		time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBUpdateColor method for updating color from database.
func (admin Model) DBUpdateColor(color models.Colors) error {

	_, err := admin.DB.Exec(`UPDATE colors 
								SET color = $1, 
									color_updated_at = $2 
								WHERE color_id = $3`, color.Color,
		time.Now(),
		color.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBDeleteColor method for deleting Color from database.
func (admin Model) DBDeleteColor(colorID int) error {

	_, err := admin.DB.Exec(`UPDATE colors 
								SET color_deleted_at = $1 
								WHERE color_id = $2`, time.Now(),
		colorID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// --------------------- Order Database -------------------

// DBGetOrders method for accessing order details
func (admin Model) DBGetOrders(request models.OrderRequest) ([]models.Orders, error) {

	var orders []models.Orders
	var sqlString string
	var sqlCondition string
	// var filter []int
	var filter []interface{}
	var count int

	sqlString = `SELECT o.order_id, 
						o.user_id, 
						o.product_id, 
						o.inventory_id, 
						o.quantity, 
						o.order_status,
						p.product_name,
						p.product_price,
						p.product_desc,
						p.product_status,
						p.product_brand_id,
						p.product_category_id,
						p.product_subcategory_id,
						i.image_id,
						i.product_image,
						c.category_name,
						c.category_desc,
						s.subcategory_name,
						s.subcategory_desc,
						b.brand_name,
						b.brand_desc,
						q.product_quantity,
						col.color_id,
						col.color									
					FROM orders o
						INNER JOIN products p
							ON o.product_id = p.product_id
						INNER JOIN categories c
							ON p.product_category_id = c.category_id
						INNER JOIN subcategories s
								ON p.product_subcategory_id = s.subcategory_id
						INNER JOIN brands b
								ON p.product_brand_id = b.brand_id
						LEFT JOIN images i
								ON o.product_id = i.product_id
						INNER JOIN inventories q
								ON o.inventory_id = q.inventory_id
						INNER JOIN colors col
								ON col.color_id = q.color_id
					WHERE p.product_deleted_at IS NULL `

	if request.OrderID != nil {
		count++
		sqlCondition = fmt.Sprint(` AND o.order_id = $`, count)
		filter = append(filter, *request.OrderID)
	}
	if request.OrderStatus != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` AND o.order_status = $`, count)
		filter = append(filter, *request.OrderStatus)
	}
	if request.UserID != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` AND o.user_id = $`, count)
		filter = append(filter, *request.UserID)
	}
	if request.Product != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` AND p.product_name = $`, count)
		filter = append(filter, *request.Product)
	}
	if request.PriceMin != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` AND p.product_price > $`, count)
		filter = append(filter, int(*request.PriceMin))
	}
	if request.PriceMax != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` AND p.product_price < $`, count)
		filter = append(filter, int(*request.PriceMax))
	}
	if request.Category != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` AND c.category_name = $`, count)
		filter = append(filter, *request.Category)
	}
	if request.Subcategory != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` AND c.subcategory_name = $`, count)
		filter = append(filter, *request.Subcategory)
	}
	if request.Brand != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` AND c.brand_name = $`, count)
		filter = append(filter, *request.Brand)
	}

	log.Println("Count:", count)
	log.Println("Filter:", filter)
	log.Println("Condition:", sqlCondition)

	rows, err := admin.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var imageID sql.NullInt32
		var image sql.NullString
		var order models.Orders

		err := rows.Scan(&order.ID,
			&order.UserID,
			&order.Product.ID,
			&order.Product.Inventory.ID,
			&order.Quantity,
			&order.Status,
			&order.Product.Name,
			&order.Product.Price,
			&order.Product.Description,
			&order.Product.Status,
			&order.Product.Brand.ID,
			&order.Product.Category.ID,
			&order.Product.Subcategory.ID,
			&imageID,
			&image,
			&order.Product.Category.Name,
			&order.Product.Category.Description,
			&order.Product.Subcategory.Name,
			&order.Product.Subcategory.Description,
			&order.Product.Brand.Name,
			&order.Product.Brand.Description,
			&order.Product.Inventory.Quantity,
			&order.Product.Inventory.Color.ID,
			&order.Product.Inventory.Color.Color,
		)
		if imageID.Valid {
			order.Product.Image.ID = int(imageID.Int32)
		}
		if image.Valid {
			order.Product.Image.Image = image.String
		}
		if err != nil {
			return orders, err
		}
		order.Product.Color.ID = order.Product.Inventory.Color.ID
		order.Product.Color.Color = order.Product.Inventory.Color.Color
		order.Product.Image.ProductID = order.Product.ID
		order.Product.Inventory.ProductID = order.Product.ID

		orders = append(orders, order)
	}
	return orders, nil
}

// DBChangeOrderStatus method for changing order status
func (admin Model) DBChangeOrderStatus(orderID int) error {

	var orderStatus string
	row := admin.DB.QueryRow(`SELECT order_status 
								FROM orders
								WHERE order_id = $1;`, orderID)
	err := row.Scan(&orderStatus)
	var status string
	if orderStatus == "Ordered" {
		status = "Shipped"
	} else if orderStatus == "Shipped" {
		status = "In-Transit"
	} else if orderStatus == "In-Transit" {
		status = "Deleivered"
	} else {
		return nil
	}
	if status != "Delivered" {
		_, err = admin.DB.Exec(`UPDATE orders
									SET order_status = $1
									WHERE order_id = $2;`, status,
			orderID,
		)
	} else {
		_, err = admin.DB.Exec(`UPDATE orders
									SET order_status = $1,
										delivered_at = $2
									WHERE order_id = $3;`, status,
			time.Now(),
			orderID,
		)
	}
	if err != nil {
		return err
	}
	return nil
}

// DBGetOffers function for accessing categoryOffers
func (admin Model) DBGetOffers(request models.CategoryOfferRequest) ([]models.CategoryOffer, error) {

	var offers []models.CategoryOffer
	var filter []interface{}
	var sqlString, sqlCondition string
	var count = 2
	sqlString = `SELECT c.category_id,
						c.category_name,
						c.category_desc,
						o.id,
						o.name,
						o.offer,
						o.offer_from,
						o.offer_to,
						CASE 
							WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN true
							ELSE false
						END AS offer_status
					FROM category_offers o
						INNER JOIN categories c
							ON o.category_id = c.category_id
					WHERE category_deleted_at IS NULL`

	// filter = append(filter, time.Now(), time.Now())
	if request.Name != nil {
		count++
		filter = append(filter, *request.Name)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND o.name ILIKE $`, count)
	}

	rows, err := admin.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter...)
	if err != nil {
		log.Println("Query Error:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var offer models.CategoryOffer
		var offerFrom, offerTo time.Time
		err = rows.Scan(&offer.Category.ID,
			&offer.Category.Name,
			&offer.Category.Description,
			&offer.ID,
			&offer.Name,
			&offer.Offer,
			&offerFrom,
			&offerTo,
			&offer.Status)
		if err != nil {
			log.Println("Scan Error:", err)
			return nil, err
		}
		offer.From = offerFrom.Format("01-02-2006")
		offer.To = offerTo.Format("01-02-2006")
		offers = append(offers, offer)
	}
	return offers, nil
}

// DBAddOffers function for adding category offer to a category.
func (admin Model) DBAddOffers(offer models.CategoryOfferRequest) error {

	log.Println("Offer:", offer)
	_, err := admin.DB.Exec(`INSERT INTO category_offers 
										(name, 
										offer, 
										category_id, 
										offer_from, 
										offer_to)
								VALUES ($1, $2, $3, $4, $5)`, offer.Name,
		offer.Offer,
		offer.CategoryID,
		offer.OfferFrom,
		offer.OfferTo,
	)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

// DBUpdateOffers method for updating categoryOffer
func (admin Model) DBUpdateOffers(offer models.CategoryOfferRequest) error {
	var sqlString, sqlCondition string
	var filter []interface{}
	var count = 0
	sqlString = `UPDATE category_offers
					SET `
	if offer.Name != nil {
		count++
		sqlCondition = fmt.Sprint(sqlCondition, ` name = $`, count)
		filter = append(filter, *offer.Name)
	}
	if offer.Offer != nil {
		count++
		if count > 1 {
			sqlCondition = fmt.Sprint(sqlCondition, `,`)
		}
		sqlCondition = fmt.Sprint(sqlCondition, ` offer = $`, count)
		filter = append(filter, *offer.Offer)
	}
	if offer.CategoryID != nil {
		count++
		if count > 1 {
			sqlCondition = fmt.Sprint(sqlCondition, `,`)
		}
		sqlCondition = fmt.Sprint(sqlCondition, ` category_id = $`, count)
		filter = append(filter, *offer.CategoryID)
	}
	if offer.OfferFrom != nil {
		count++
		if count > 1 {
			sqlCondition = fmt.Sprint(sqlCondition, `,`)
		}
		sqlCondition = fmt.Sprint(sqlCondition, ` offer_from = $`, count)
		filter = append(filter, *offer.OfferFrom)
	}
	if offer.OfferTo != nil {
		count++
		if count > 1 {
			sqlCondition = fmt.Sprint(sqlCondition, `,`)
		}
		sqlCondition = fmt.Sprint(sqlCondition, ` offer_to = $`, count)
		filter = append(filter, *offer.OfferTo)
	}
	count++
	sqlCondition = fmt.Sprint(sqlCondition, ` WHERE id = $`, count)
	filter = append(filter, *offer.ID)

	_, err := admin.DB.Exec(fmt.Sprint(sqlString, sqlCondition), filter...)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

// DBDeleteOffers method for deleting an offer.
func (admin Model) DBDeleteOffers(offerID *int) error {

	_, err := admin.DB.Exec(`DELETE 
								FROM category_offers 
									WHERE id = $1`, offerID)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}
