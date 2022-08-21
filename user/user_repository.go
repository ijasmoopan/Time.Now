package user

import (
	// "context"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/usecases"
)

// Model describing struct for database connection.
type Model struct {
	DB *sql.DB
}

// DBGetProducts for accessing products with respect to reqeust params.
func (user Model) DBGetProducts(request models.ProductRequest) ([]models.ProductWithInventory, error) {

	file := usecases.Logger()
	log.SetOutput(file)

	var products []models.ProductWithInventory
	filter := make([]interface{}, 0, 10)
	count := 1
	var sqlCondition, sqlString string
	var rows *sql.Rows
	var err error

	sqlString = `WITH cte_getproducts (product_id, 
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
										offer_price,
										wishlist
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
								END AS offer_price`

	if request.UserID != nil {
		sqlString = fmt.Sprint(sqlString, `,
										CASE 
											WHEN p.product_id = w.product_id AND w.user_id = $`, count, `THEN true
											ELSE false
										END AS wishlist `)
		filter = append(filter, *request.UserID)
		count++
	} else {
		sqlString = fmt.Sprint(sqlString, `,
											CASE 
												WHEN p.product_id = w.product_id THEN false
												ELSE false
											END AS wishlist `)
	}
	sqlString = fmt.Sprint(sqlString, ` FROM   products p
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
											
											LEFT JOIN wishlist w
												ON w.product_id = p.product_id
											LEFT JOIN users u
												ON u.user_id = w.user_id
										WHERE  p.product_deleted_at IS NULL
											ORDER BY offer)

										SELECT * 
											FROM cte_getproducts
											WHERE product_deleted_at IS NULL `)
	if request.Product != nil {
		for idx, value := range request.Product {
			if idx == 0 {
				filter = append(filter, fmt.Sprint("%", *value, "%"))
				sqlCondition = fmt.Sprint(` AND product_name ILIKE $`, count)
				count++
			} else {
				filter = append(filter, fmt.Sprint("%", *value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` OR product_name ILIKE $`, count)
				count++
			}
		}
	}
	if request.Category != nil {
		for idx, value := range request.Category {
			if idx == 0 {
				filter = append(filter, fmt.Sprint("%", *value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` AND category_name ILIKE $`, count)
				count++
			} else {
				filter = append(filter, fmt.Sprint("%", *value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` OR category_name ILIKE $`, count)
				count++
			}
		}
	}
	if request.Subcategory != nil {
		for idx, value := range request.Subcategory {
			if idx == 0 {
				filter = append(filter, fmt.Sprint("%", *value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` AND subcategory_name ILIKE $`, count)
				count++
			} else {
				filter = append(filter, fmt.Sprint("%", *value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` OR subcategory_name ILIKE $`, count)
				count++
			}
		}
	}
	if request.Brand != nil {
		for idx, value := range request.Brand {
			if idx == 0 {
				filter = append(filter, fmt.Sprint("%", *value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` AND brand_name ILIKE $`, count)
				count++
			} else {
				filter = append(filter, fmt.Sprint("%", *value, "%"))
				sqlCondition = fmt.Sprint(sqlCondition, ` OR brand_name ILIKE $`, count)
				count++
			}
		}
	}
	if request.PriceMin != nil {
		filter = append(filter, *request.PriceMin)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND product_price >= $`, count)
		count++
	}
	if request.PriceMax != nil {
		filter = append(filter, *request.PriceMax)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND product_price <= $`, count)
		count++
	}
	// wishlist
	if request.Wishlist != nil {
		filter = append(filter, *request.Wishlist)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND wishlist = $`, count)
		count++
	}
	// offer
	if request.OfferMin != nil {
		filter = append(filter, *request.OfferMin)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND offer >= $`, count)
		count++
	}
	if request.OfferMax != nil {
		filter = append(filter, *request.OfferMax)
		sqlCondition = fmt.Sprint(sqlCondition, ` AND offer <= $`, count)
		count++
	}
	// paging
	var limit = 5
	sqlCondition = fmt.Sprint(sqlCondition, ` ORDER BY product_id`)
	if request.Page != nil {
		*request.Page = (*request.Page - 1) * 5
		filter = append(filter, limit, *request.Page)
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

	log.Println("sqlCondition:", sqlCondition)
	log.Println("filter:", filter)

	rows, err = user.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter...)

	if err != nil {
		log.Println(err)
		return products, err
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
			&product.Wishlist,
		)
		if imageID.Valid {
			product.Image.ID = int(imageID.Int32)
			product.Image.ProductID = product.ID
		}
		if image.Valid {
			product.Image.Image = image.String
		}
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
			return products, err
		}
		log.Println("DB Product:", product)

		var colors []models.Colors
		var inventories []models.Inventories
		filter2 := make([]interface{}, 0, 10)
		count2 := 0
		var sqlCondition2 string
		inventoryrows, err := user.DB.Query(`SELECT i.inventory_id, 
													i.product_id, 
													c.color_id,
													c.color, 
													i.product_quantity 
												FROM inventories i
												INNER JOIN colors c
													ON i.color_id = c.color_id
												WHERE i.inventory_deleted_at IS NULL
													AND i.product_id = $1;`, product.ID)
		if request.Color != nil {
			for idx, value := range request.Color {
				if idx == 0 {
					count2++
					filter2 = append(filter2, fmt.Sprint("%", *value, "%"))
					sqlCondition2 = fmt.Sprint(sqlCondition2, ` AND c.color ILIKE $`, count2)
				} else {
					count2++
					filter2 = append(filter2, fmt.Sprint("%", *value, "%"))
					sqlCondition2 = fmt.Sprint(sqlCondition2, ` OR c.color ILIKE $`, count2)
				}
			}
		}
		if err != nil {
			log.Println("Inventoryrows:", err)
			return nil, err
		}
		defer inventoryrows.Close()
		for inventoryrows.Next() {
			var inventory models.Inventories
			var color models.Colors
			err = inventoryrows.Scan(&inventory.ID,
				&inventory.ProductID,
				&inventory.Color.ID,
				&inventory.Color.Color,
				&inventory.Quantity)
			if err != nil {
				log.Println("for loop:", err)
				return products, err
			}
			log.Println("Color:", inventory)
			color.ID = inventory.Color.ID
			color.Color = inventory.Color.Color
			colors = append(colors, color)
			inventories = append(inventories, inventory)
		}
		product.Inventory = inventories
		product.Color = colors

		log.Println("Product:", product)
		products = append(products, product)
	}
	return products, nil
}

// DBGetProductsWithColors method access all products corresponding to request params.
// func (user Model) DBGetProductsWithColors(request models.ProductRequest) ([]models.ProductWithColor, error) {
// 	var products []models.ProductWithColor
// 	var sqlCondition string
// 	sqlString := `SELECT p.product_id,
// 						p.product_name,
// 						p.product_price,
// 						p.product_desc,
// 						i.image_id,
// 						p.product_brand_id,
// 						p.product_category_id,
// 						p.product_subcategory_id,
// 						i.product_id,
// 						i.product_image,
// 						c.category_name,
// 						c.category_desc,
// 						s.subcategory_name,
// 						s.subcategory_desc,
// 						b.brand_name,
// 						b.brand_desc
// 					FROM   products p
// 						INNER JOIN categories c
// 								ON p.product_category_id = c.category_id
// 						INNER JOIN subcategories s
// 								ON p.product_subcategory_id = s.subcategory_id
// 						INNER JOIN brands b
// 								ON p.product_brand_id = b.brand_id
// 						LEFT JOIN images i
// 								ON p.product_id = i.product_id
// 					WHERE  p.product_deleted_at IS NULL `
// 	var rows *sql.Rows
// 	var err error
// 	var filter []interface{}
// 	count := 1
// 	if request.Product != nil {
// 		filter = append(filter, *request.Product)
// 		sqlCondition = fmt.Sprint(` AND p.product_name = $`, count)
// 		count++
// 	}
// 	if request.Category != nil {
// 		filter = append(filter, *request.Category)
// 		sqlCondition = fmt.Sprint(sqlCondition, ` AND p.product_category_name = $`, count)
// 		count++
// 	}
// 	if request.Subcategory != nil {
// 		filter = append(filter, *request.Subcategory)
// 		sqlCondition = fmt.Sprint(sqlCondition, ` AND p.product_subcategory_name = $`, count)
// 		count++
// 	}
// 	if request.Brand != nil {
// 		filter = append(filter, *request.Brand)
// 		sqlCondition = fmt.Sprint(sqlCondition, ` AND p.product_brand_name = $`, count)
// 		count++
// 	}
// 	if request.PriceMin != nil {
// 		filter = append(filter, *request.PriceMin)
// 		sqlCondition = fmt.Sprint(sqlCondition, ` AND p.product_price > $`, count)
// 		count++
// 	}
// 	if request.PriceMax != nil {
// 		filter = append(filter, *request.PriceMax)
// 		sqlCondition = fmt.Sprint(sqlCondition, ` AND p.product_price < $`, count)
// 		count++
// 	}
// 	rows, err = user.DB.Query(fmt.Sprint(sqlString, sqlCondition), filter...)
// 	if err != nil {
// 		log.Println(err)
// 		return products, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var product models.ProductWithColor
// 		var imageID sql.NullInt32
// 		var productID sql.NullInt32
// 		var image sql.NullString
// 		err := rows.Scan(&product.ID,
// 			&product.Name,
// 			&product.Price,
// 			&product.Description,
// 			&imageID,
// 			&product.Brand.ID,
// 			&product.Category.ID,
// 			&product.Subcategory.ID,
// 			&productID,
// 			&image,
// 			&product.Category.Name,
// 			&product.Category.Description,
// 			&product.Subcategory.Name,
// 			&product.Subcategory.Description,
// 			&product.Brand.Name,
// 			&product.Brand.Description,
// 		)
// 		if imageID.Valid {
// 			product.Image.ID = int(imageID.Int32)
// 		}
// 		if productID.Valid {
// 			product.Image.ProductID = int(productID.Int32)
// 		}
// 		if image.Valid {
// 			product.Image.Image = image.String
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			return products, err
// 		}
// 		log.Println(product)
// 		// ---------------------- Colors -------------------------
// 		colorRows, err := user.DB.Query(`SELECT c.color_id,
// 											c.color
// 										FROM colors c
// 										INNER JOIN inventories i
// 											ON i.color_id = c.color_id
// 										WHERE i.product_id = $1`, product.ID)
// 		if err != nil {
// 			log.Println(err)
// 			return products, err
// 		}
// 		defer colorRows.Close()
// 		for colorRows.Next() {
// 			var color models.Colors
// 			err = colorRows.Scan(&color.ID,
// 				&color.Color)
// 			product.Color = append(product.Color, color)
// 		}
// 		// ------------------------ END --------------------------
// 		products = append(products, product)
// 	}
// 	return products, nil
// }

// DBGetAllColorsOfAProduct get all colors of a product.
func (user Model) DBGetAllColorsOfAProduct(productID int) ([]models.Inventories, error) {

	var inventories []models.Inventories
	rows, err := user.DB.Query(`SELECT inventory_id,
									product_color
								FROM   inventories
								WHERE  product_id = $1
									AND product_quantity > $2
									AND inventory_deleted_at IS NULL; `, productID, "0")

	if err != nil {
		log.Println(err)
		return inventories, err
	}
	defer rows.Close()

	for rows.Next() {
		var inventory models.Inventories
		err = rows.Scan(&inventory.ID, &inventory.Color)
		if err != nil {
			log.Println(err)
			return inventories, err
		}
		inventories = append(inventories, inventory)
		log.Println(inventory)
	}
	return inventories, err
}

// DBGetRecommendedProducts get similar products of productID.
func (user Model) DBGetRecommendedProducts(productID, categoryID, SubcategoryID int) ([]models.Product, error) {

	var products []models.Product
	rows, err := user.DB.Query(`SELECT p.product_id,
									p.product_name,
									p.product_price,
									p.product_desc,
									p.product_image_id,
									p.product_brand_id,
									p.product_category_id,
									p.product_subcategory_id,
									i.product_image,
									c.category_name,
									c.category_desc,
									s.subcategory_name,
									s.subcategory_desc,
									b.brand_name,
									b.brand_desc
								FROM   products p
									INNER JOIN categories c
											ON p.product_category_id = c.category_id
									INNER JOIN subcategories s
											ON p.product_subcategory_id = s.subcategory_id
									INNER JOIN brands b
											ON p.product_brand_id = b.brand_id
									LEFT JOIN images i
										ON p.product_image_id = i.image_id
								WHERE  p.product_deleted_at IS NULL
									AND p.product_category_id = $1
									AND p.product_subcategory_id = $2
									AND p.product_id <> $3;`, categoryID, SubcategoryID, productID)

	if err != nil {
		log.Println(err)
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		var imageID sql.NullInt32
		var image sql.NullString

		err := rows.Scan(&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
			&imageID,
			&product.Brand.ID,
			&product.Category.ID,
			&product.Subcategory.ID,
			&image,
			&product.Category.Name,
			&product.Category.Description,
			&product.Subcategory.Name,
			&product.Subcategory.Description,
			&product.Brand.Name,
			&product.Brand.Description,
		)

		if imageID.Valid {
			product.Image.ID = int(imageID.Int32)
		}
		if imageID.Valid {
			product.Image.Image = image.String
		}
		if err != nil {
			return products, err
		}
		log.Println(err)
		products = append(products, product)
	}
	return products, nil
}

// -------------------------------User management-------------------------

// DBGetUser method for accessing a single user.
func (user Model) DBGetUser(userID string) (models.User, error) {

	var usr models.User
	var image sql.NullString
	row := user.DB.QueryRow(`SELECT u.user_id, 
								u.user_firstname, 
								u.user_secondname, 
								u.user_phone, 
								u.user_email, 
								u.user_gender,
								u.user_image 
							FROM users u
							WHERE u.deleted_at IS NULL 
								AND u.user_id = $1;`, userID)

	err := row.Scan(&usr.ID,
		&usr.FirstName,
		&usr.SecondName,
		&usr.Phone,
		&usr.Email,
		&usr.Gender,
		&image,
	)
	if image.Valid {
		usr.Image = image.String
	}
	if err != nil {
		log.Println(err)
		return usr, err
	}
	return usr, nil
}

// DBAuthUser for authenticating user by middleware.
func (user Model) DBAuthUser(userID string) (models.UserLogin, error) {

	var usr models.UserLogin
	row := user.DB.QueryRow(`SELECT user_id, 
								user_email, 
								user_password 
							FROM users 
							WHERE user_id = $1`, userID)

	err := row.Scan(&usr.ID, &usr.Email, &usr.Password)
	if err != nil {
		log.Println(err)
		return usr, err
	}
	return usr, nil
}

// DBValidateUser method for validating user input details.
func (user Model) DBValidateUser(loginUser models.UserLogin) (models.UserLogin, error) {

	var usr models.UserLogin

	row := user.DB.QueryRow(`SELECT user_id, 
								user_email, 
								user_password 
							FROM users 
							WHERE user_email = $1`, loginUser.Email)

	err := row.Scan(&usr.ID, &usr.Email, &usr.Password)
	if err != nil {
		log.Println(err)
		return usr, err
	}
	log.Println("user repo:", usr)
	return usr, nil
}

// DBCheckUserStatus method for checking if the user is blocked by admin or not.
func (user Model) DBCheckUserStatus(usr models.UserLogin)(error) {
	var checkStatus bool

	row := user.DB.QueryRow(`SELECT user_status 
								FROM users
								WHERE user_id = $1`, usr.ID)
	err := row.Scan(&checkStatus)
	if err != nil {
		return err
	}
	if !checkStatus {
		return err
	} else {
		return nil
	}
}

// DBUserRegistration method for registering a new user.
func (user Model) DBUserRegistration(newUser models.UserRegister) error {

	_, err := user.DB.Exec(`INSERT INTO users
								(user_firstname,
								user_secondname,
								user_password,
								user_phone,
								user_email,
								user_gender,
								user_referral,
								created_at)
							VALUES      ($1,
								$2,
								$3,
								$4,
								$5,
								$6,
								$7, 
								$8);`, 
		newUser.FirstName,
		newUser.SecondName,
		newUser.Password,
		newUser.Phone,
		newUser.Email,
		newUser.Gender,
		newUser.Referral,
		time.Now())
	// time.Now().Format("15:04:05 02-01-2006"))

	if err != nil {
		return err
	}
	return nil
}

// DBUpdateUser method is for updating user details.
func (user Model) DBUpdateUser(usr models.User) error {

	_, err := user.DB.Exec(`UPDATE users 
								SET user_firstname = $1,
									user_secondname = $2,
									user_phone = $3,
									user_email = $4,
									user_image = $5
									updated_at = $6
								WHERE user_id = $7;`, usr.FirstName,
		usr.SecondName,
		usr.Phone,
		usr.Email,
		usr.Image,
		time.Now(),
		usr.ID,
		// time.Now().Format("15:04:05 02-01-2006"))
	)

	if err != nil {
		return err
	}
	return nil
}

// DBDeleteUser method for deleting a user.
func (user Model) DBDeleteUser(userID string) error {

	_, err := user.DB.Exec(`UPDATE users 
							SET deleted_at= $1
							WHERE user_id = $2;`, time.Now(), userID)
	if err != nil {
		return err
	}
	return nil
}

// ------------------ User Address --------------------

// DBGetAddress method for accessing address of a user.
func (user Model) DBGetAddress(userID int) ([]models.Address, error) {

	var addresses []models.Address

	rows, err := user.DB.Query(`SELECT address_id,
									user_id,
									address_name,
									address_phone,
									address_pincode,
									address_housename,
									address_streetname,
									address_city,
									address_state,
									address_desc
								FROM   address
								WHERE  user_id = $1; `, userID)
	if err != nil {
		log.Println("Query", err)
		return addresses, err
	}
	defer rows.Close()

	for rows.Next() {
		var address models.Address
		var addressDesc sql.NullString
		err := rows.Scan(&address.ID,
			&address.UserID,
			&address.Name,
			&address.Phone,
			&address.Pincode,
			&address.HouseName,
			&address.StreetName,
			&address.City,
			&address.State,
			&addressDesc)
		if addressDesc.Valid {
			address.Description = addressDesc.String
		}
		if err != nil {
			log.Println("Scan", err)
			return addresses, err
		}
		addresses = append(addresses, address)
		log.Println("Address:", address)
	}
	return addresses, nil
}

// DBAddAddress method for adding new address.
func (user Model) DBAddAddress(address models.Address) error {

	_, err := user.DB.Exec(`INSERT INTO address (user_id,
											address_name,
											address_phone,
											address_pincode,
											address_housename,
											address_streetname,
											address_city,
											address_state,
											address_desc,
											address_created_at)
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10); `, address.UserID,
		address.Name,
		address.Phone,
		address.Pincode,
		address.HouseName,
		address.StreetName,
		address.City,
		address.State,
		address.Description,
		time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBUpdateAddress method for updating address.
func (user Model) DBUpdateAddress(address models.Address) error {

	_, err := user.DB.Exec(`UPDATE address
							SET address_name = $1,
								address_phone = $2,
								address_pincode = $3,
								address_housename = $4,
								address_streetname = $5,
								address_city = $6,
								address_state = $7,
								address_desc = $8,
								address_updated_at = $9;`, address.Name,
		address.Phone,
		address.Pincode,
		address.HouseName,
		address.StreetName,
		address.City,
		address.State,
		address.Description,
		time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DBDeleteAddress method for deleting an address.
func (user Model) DBDeleteAddress(addressID int) error {

	_, err := user.DB.Exec(`UPDATE address
							SET address_deleted_at = $1 
							WHERE address_id = $2;`, addressID, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// --------------------------Cart Management-----------------------

// DBGetCart method is using for accessing cart products.
func (user Model) DBGetCart(userID int) ([]models.Cart, int, error) {

	var countOfProducts int
	var cartProducts []models.Cart
	rows, err := user.DB.Query(`SELECT p.product_id,
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
									q.inventory_id,
									q.product_quantity,
									col.color_id,
									col.color,
									cart.cart_id,
									cart.quantity,
									cart.user_id
								FROM cart
									INNER JOIN products p
											ON cart.product_id = p.product_id
									INNER JOIN categories c
											ON p.product_category_id = c.category_id
									INNER JOIN subcategories s
											ON p.product_subcategory_id = s.subcategory_id
									INNER JOIN brands b
											ON p.product_brand_id = b.brand_id
									LEFT JOIN images i
											ON p.product_id = i.product_id
									INNER JOIN inventories q
											ON q.inventory_id = cart.inventory_id
									INNER JOIN colors col
											ON col.color_id = q.color_id
								WHERE  cart.user_id = $1
									AND p.product_deleted_at IS NULL
									AND cart.deleted_at IS NULL;`, userID)

	if err != nil {
		log.Println("Error:", err)
		return cartProducts, countOfProducts, err
	}
	defer rows.Close()
	log.Println("for loop")
	for rows.Next() {
		var cartProduct models.Cart
		var productImageID sql.NullInt32
		var productImage sql.NullString
		err := rows.Scan(&cartProduct.Product.ID,
			&cartProduct.Product.Name,
			&cartProduct.Product.Price,
			&cartProduct.Product.Description,
			&cartProduct.Product.Status,
			&cartProduct.Product.Brand.ID,
			&cartProduct.Product.Category.ID,
			&cartProduct.Product.Subcategory.ID,
			&productImageID,
			&productImage,
			&cartProduct.Product.Category.Name,
			&cartProduct.Product.Category.Description,
			&cartProduct.Product.Subcategory.Name,
			&cartProduct.Product.Subcategory.Description,
			&cartProduct.Product.Brand.Name,
			&cartProduct.Product.Brand.Description,
			&cartProduct.Product.Inventory.ID,
			&cartProduct.Product.Inventory.Quantity,
			&cartProduct.Product.Inventory.Color.ID,
			&cartProduct.Product.Inventory.Color.Color,
			&cartProduct.ID,
			&cartProduct.Quantity,
			&cartProduct.UserID)
		if productImageID.Valid {
			cartProduct.Product.Image.ID = int(productImageID.Int32)
		}
		if productImage.Valid {
			cartProduct.Product.Image.Image = productImage.String
		}
		if err != nil {
			log.Println("Error in for loop:", err)
			return cartProducts, countOfProducts, err
		}
		cartProduct.Product.Color.ID = cartProduct.Product.Inventory.Color.ID
		cartProduct.Product.Color.Color = cartProduct.Product.Inventory.Color.Color
		cartProduct.Product.Image.ProductID = cartProduct.Product.ID
		cartProduct.Product.Inventory.ProductID = cartProduct.Product.ID

		cartProducts = append(cartProducts, cartProduct)
		countOfProducts = countOfProducts + cartProduct.Quantity
	}
	return cartProducts, countOfProducts, nil
}

// DBAddCart method for adding product into cart.
func (user Model) DBAddCart(cartProduct models.Cart) error {
	log.Println("DBAdd")
	var err error
	row := user.DB.QueryRow(`INSERT INTO cart
								(user_id, 
								product_id, 
								inventory_id, 
								quantity,
								created_at)
							VALUES ($1, $2, $3, $4, $5) 
								RETURNING cart_id;`, cartProduct.UserID,
		cartProduct.Product.ID,
		cartProduct.Product.Inventory.ID,
		cartProduct.Quantity,
		time.Now(),
	)
	err = row.Scan(&cartProduct.ID)
	if err != nil {
		log.Println("Err", err)
		return err
	}
	// Reduce product quantity in inventories
	// _, err = user.DB.Exec(`UPDATE inventories
	// 							SET product_quantity = product_quantity - $1
	// 							WHERE inventory_id = $2;`, cartProduct.Quantity,
	// 		cartProduct.Product.Inventory.ID,
	// 	)
	// if err != nil {
	// 	log.Println("Err", err)
	// 	return err
	// }
	// durationOfTime := time.Duration(15 * time.Minute)
	// f := func() {
	// 	log.Println("After 15 minutes... ")
	// 	// Buisness logic
	// 	log.Println("Checking order_id is in order table or not")
	// 	row := user.DB.QueryRow(`SELECT order_id
	// 								FROM orders
	// 								WHERE cart_id = $1;`, cartProduct.ID)
	// 	err = row.Scan(&cartProduct.ID)
	// 	if err == sql.ErrNoRows {
	// 		log.Println("Checking cart_id is still in cart table or not")
	// 		row = user.DB.QueryRow(`SELECT cart_id
	// 									FROM cart
	// 									WHERE cart_id = $1
	// 										AND deleted_at IS NULL;`, cartProduct.ID)
	// 		err = row.Scan(&cartProduct.ID)
	// 		if err == sql.ErrNoRows {
	// 		// Increase product quantity in inventories
	// 		log.Println("No rows.. so increasing quantity in inventories")
	// 		_, err = user.DB.Exec(`UPDATE inventories
	// 								SET product_quantity = product_quantity + $1
	// 								WHERE inventory_id = $2;`, cartProduct.Quantity,
	// 				cartProduct.Product.Inventory.ID,
	// 			)
	// 		}
	// 	}
	// 	if err != nil {
	// 		log.Println("Err", err)
	// 		return
	// 	}
	// }
	// timer := time.AfterFunc(durationOfTime, f)
	// defer timer.Stop()

	return nil
}

// DBUpdateCart method for updating product from cart.
func (user Model) DBUpdateCart(cart models.Cart) error {

	// var quantity int
	// log.Println("Checking cart_id is still in cart table or not")
	// row := user.DB.QueryRow(`SELECT quantity
	// 							FROM cart
	// 							WHERE cart_id = $1
	// 								AND deleted_at IS NULL;`, cart.ID)
	// err := row.Scan(&quantity)
	// if cart.Quantity == quantity {
	_, err := user.DB.Exec(`UPDATE cart
										SET inventory_id = $1,
										quantity = $2,
										updated_at = $3
									WHERE cart_id = $4;`, cart.Product.Inventory.ID,
		cart.Quantity,
		time.Now(),
		cart.ID,
	)
	// }
	// ----------------------------------------------------------------------------------------------------
	// _, err = user.DB.Exec(`UPDATE inventories
	// 						SET product_quantity = product_quantity + $1
	// 						WHERE inventory_id = $2;`, cart.Quantity,
	// 				cart.Product.Inventory.ID,
	// 			)

	if err != nil {
		return err
	}
	return nil
}

// DBDeleteCart method for deleting product from cart
func (user Model) DBDeleteCart(cartID int) error {
	_, err := user.DB.Exec(`UPDATE cart 
								SET deleted_at = $1
								WHERE cart_id = $2;`, time.Now(),
		cartID,
	)
	if err != nil {
		return err
	}
	return nil
}

// ---------------------------------- Wishlist Management ---------------------------------

// DBGetWishlist method for accessing wishlist details of a user.
func (user Model) DBGetWishlist(userID int) ([]models.Wishlist, error) {

	var wishlist []models.Wishlist

	rows, err := user.DB.Query(`SELECT p.product_id,
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
									q.inventory_id,
									q.product_quantity,
									col.color_id,
									col.color,
									wishlist.wishlist_id,
									wishlist.user_id
								FROM wishlist
									INNER JOIN products p
											ON wishlist.product_id = p.product_id
									INNER JOIN categories c
											ON p.product_category_id = c.category_id
									INNER JOIN subcategories s
											ON p.product_subcategory_id = s.subcategory_id
									INNER JOIN brands b
											ON p.product_brand_id = b.brand_id
									LEFT JOIN images i
											ON i.product_id = p.product_id
									INNER JOIN inventories q
											ON q.inventory_id = wishlist.inventory_id
									INNER JOIN colors col
											ON col.color_id = q.color_id
								WHERE  wishlist.user_id= $1
									AND p.product_deleted_at IS NULL;`, userID)

	if err != nil {
		log.Println("Error:", err)
		return wishlist, err
	}
	log.Println("No error")
	defer rows.Close()

	for rows.Next() {
		var product models.Wishlist
		var imageID sql.NullInt32
		var image sql.NullString
		err := rows.Scan(&product.Product.ID,
			&product.Product.Name,
			&product.Product.Price,
			&product.Product.Description,
			&product.Product.Status,
			&product.Product.Brand.ID,
			&product.Product.Category.ID,
			&product.Product.Subcategory.ID,
			&imageID,
			&image,
			&product.Product.Category.Name,
			&product.Product.Category.Description,
			&product.Product.Subcategory.Name,
			&product.Product.Subcategory.Description,
			&product.Product.Brand.Name,
			&product.Product.Brand.Description,
			&product.Product.Inventory.ID,
			&product.Product.Inventory.Quantity,
			&product.Product.Inventory.Color.ID,
			&product.Product.Inventory.Color.Color,
			&product.ID,
			&product.UserID)
		if imageID.Valid {
			product.Product.Image.ID = int(imageID.Int32)
		}
		if image.Valid {
			product.Product.Image.Image = image.String
		}
		if err != nil {
			return wishlist, err
		}
		product.Product.Color.ID = product.Product.Inventory.Color.ID
		product.Product.Color.Color = product.Product.Inventory.Color.Color
		product.Product.Image.ProductID = product.Product.ID
		product.Product.Inventory.ProductID = product.Product.ID

		wishlist = append(wishlist, product)
	}
	return wishlist, nil
}

// DBAddWishlist method for adding product into wishlist
func (user Model) DBAddWishlist(wishlist models.Wishlist) error {

	_, err := user.DB.Exec(`INSERT INTO wishlist
								(user_id,
								product_id,
								inventory_id)
							VALUES ($1, $2, $3);`, wishlist.UserID,
		wishlist.Product.ID,
		wishlist.Product.Inventory.ID)

	if err != nil {
		return err
	}
	return nil
}

// DBDeleteWishlist method for deleting product from wishlist
func (user Model) DBDeleteWishlist(wishlistID int) error {

	_, err := user.DB.Exec(`DELETE 
							FROM wishlist 
							WHERE wishlist_id = $1;`, wishlistID)
	if err != nil {
		return err
	}
	return nil
}

// ---------------------------------Checkout---------------------------

// DBCartCheckout method for cart checkout
func (user Model) DBCartCheckout(userID int) (models.CartCheckout, int, float64, error) {

	var checkout models.CartCheckout
	var countOfProducts int
	var totalPrice float64
	rows, err := user.DB.Query(`SELECT p.product_id,
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
										q.inventory_id,
										q.product_quantity,
										col.color_id,
										col.color,
										cart.cart_id,
										cart.quantity,
										cart.user_id,
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
									FROM cart
										INNER JOIN products p
												ON cart.product_id = p.product_id
										INNER JOIN categories c
												ON p.product_category_id = c.category_id
										INNER JOIN subcategories s
												ON p.product_subcategory_id = s.subcategory_id
										INNER JOIN brands b
												ON p.product_brand_id = b.brand_id
										LEFT JOIN images i
												ON p.product_id = i.product_id
										INNER JOIN inventories q
												ON q.inventory_id = cart.inventory_id
										INNER JOIN colors col
												ON col.color_id = q.color_id
										LEFT JOIN category_offers o
												ON o.category_id = p.product_category_id
									WHERE  cart.user_id = $1
										AND p.product_deleted_at IS NULL
										AND cart.deleted_at IS NULL;`, userID)

	if err != nil {
		log.Println("Error:", err)
		return checkout, countOfProducts, totalPrice, err
	}
	defer rows.Close()
	log.Println("for loop")
	for rows.Next() {
		var cartProduct models.Cart
		var productImageID, offerID, offer sql.NullInt32
		var productImage, offerName sql.NullString
		var offerFrom, offerTo sql.NullTime
		var offerStatus bool
		var offerPrice sql.NullFloat64

		err := rows.Scan(&cartProduct.Product.ID,
			&cartProduct.Product.Name,
			&cartProduct.Product.Price,
			&cartProduct.Product.Description,
			&cartProduct.Product.Status,
			&cartProduct.Product.Brand.ID,
			&cartProduct.Product.Category.ID,
			&cartProduct.Product.Subcategory.ID,
			&productImageID,
			&productImage,
			&cartProduct.Product.Category.Name,
			&cartProduct.Product.Category.Description,
			&cartProduct.Product.Subcategory.Name,
			&cartProduct.Product.Subcategory.Description,
			&cartProduct.Product.Brand.Name,
			&cartProduct.Product.Brand.Description,
			&cartProduct.Product.Inventory.ID,
			&cartProduct.Product.Inventory.Quantity,
			&cartProduct.Product.Inventory.Color.ID,
			&cartProduct.Product.Inventory.Color.Color,
			&cartProduct.ID,
			&cartProduct.Quantity,
			&cartProduct.UserID,
			&offerID,
			&offerName,
			&offer,
			&offerFrom,
			&offerTo,
			&offerStatus,
			&offerPrice,
		)
		if productImageID.Valid {
			cartProduct.Product.Image.ID = int(productImageID.Int32)
		}
		if productImage.Valid {
			cartProduct.Product.Image.Image = productImage.String
		}
		if offerStatus != false {
			log.Println("Offer status is ", offerStatus)
			cartProduct.Product.Offer.Status = offerStatus

			if offerID.Valid {
				cartProduct.Product.Offer.ID = int(offerID.Int32)
			}
			if offerName.Valid {
				cartProduct.Product.Offer.Name = offerName.String
			}
			if offer.Valid {
				cartProduct.Product.Offer.Offer = int(offer.Int32)
			}
			if offerFrom.Valid {
				cartProduct.Product.Offer.From = offerFrom.Time.Format("01-02-2006")
			}
			if offerTo.Valid {
				cartProduct.Product.Offer.To = offerTo.Time.Format("01-02-2006")
			}
			if offerPrice.Valid {
				cartProduct.Product.OfferPrice = offerPrice.Float64
			}
			cartProduct.Product.Offer.Category.ID = cartProduct.Product.Category.ID
			cartProduct.Product.Offer.Category.Name = cartProduct.Product.Category.Name
			cartProduct.Product.Offer.Category.Description = cartProduct.Product.Category.Description
		}
		if err != nil {
			log.Println("Error in for loop:", err)
			return checkout, countOfProducts, totalPrice, err
		}
		cartProduct.Product.Color.ID = cartProduct.Product.Inventory.Color.ID
		cartProduct.Product.Color.Color = cartProduct.Product.Inventory.Color.Color
		cartProduct.Product.Image.ProductID = cartProduct.Product.ID
		cartProduct.Product.Inventory.ProductID = cartProduct.Product.ID

		checkout.Cart = append(checkout.Cart, cartProduct)
		countOfProducts = countOfProducts + cartProduct.Quantity
		totalPrice = totalPrice + (cartProduct.Product.Price * float64(cartProduct.Quantity))
	}

	rows, err = user.DB.Query(`SELECT a.address_id,
									a.user_id,
									a.address_name,
									a.address_phone,
									a.address_pincode,
									a.address_housename,
									a.address_streetname,
									a.address_city,
									a.address_state,
									a.address_desc
								FROM address a
									INNER JOIN cart 
										ON a.user_id = cart.user_id 
								WHERE a.user_id = $1;`, userID)
	if err != nil {
		return checkout, countOfProducts, totalPrice, err
	}
	defer rows.Close()
	var address models.Address
	var description sql.NullString

	if rows.Next() {
		err = rows.Scan(&address.ID,
			&address.UserID,
			&address.Name,
			&address.Phone,
			&address.Pincode,
			&address.HouseName,
			&address.StreetName,
			&address.City,
			&address.State,
			&description,
		)
		if description.Valid {
			address.Description = description.String
		}
		if err != nil {
			return checkout, countOfProducts, totalPrice, err
		}
		checkout.Address = append(checkout.Address, address)

		for rows.Next() {
			err = rows.Scan(&address.ID,
				&address.UserID,
				&address.Name,
				&address.Phone,
				&address.Pincode,
				&address.HouseName,
				&address.StreetName,
				&address.City,
				&address.State,
				&description,
			)
			if description.Valid {
				address.Description = description.String
			}
			if err != nil {
				return checkout, countOfProducts, totalPrice, err
			}
			checkout.Address = append(checkout.Address, address)
		}
	} else {
		return checkout, countOfProducts, totalPrice, errors.New("Add new address")
	}
	return checkout, countOfProducts, totalPrice, nil
}

// DBProductCheckout method for instant buy
func (user Model) DBProductCheckout(productCheckout models.ProductCheckout) (models.ProductCheckout, float64, error) {

	var productImage sql.NullString
	var productImageID sql.NullInt32
	var totalPrice float64
	row := user.DB.QueryRow(`SELECT p.product_id,
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
									q.inventory_id,
									q.product_quantity,
									col.color_id,
									col.color
								FROM products p
									INNER JOIN categories c
											ON p.product_category_id = c.category_id
									INNER JOIN subcategories s
											ON p.product_subcategory_id = s.subcategory_id
									INNER JOIN brands b
											ON p.product_brand_id = b.brand_id
									LEFT JOIN images i
											ON p.product_id = i.product_id
									INNER JOIN inventories q
											ON p.product_id = q.product_id
									INNER JOIN colors col
											ON col.color_id = q.color_id
								WHERE  q.inventory_id = $1
									AND p.product_deleted_at IS NULL;`, productCheckout.Product.Inventory.ID)

	err := row.Scan(&productCheckout.Product.ID,
		&productCheckout.Product.Name,
		&productCheckout.Product.Price,
		&productCheckout.Product.Description,
		&productCheckout.Product.Status,
		&productCheckout.Product.Brand.ID,
		&productCheckout.Product.Category.ID,
		&productCheckout.Product.Subcategory.ID,
		&productImageID,
		&productImage,
		&productCheckout.Product.Category.Name,
		&productCheckout.Product.Category.Description,
		&productCheckout.Product.Subcategory.Name,
		&productCheckout.Product.Subcategory.Description,
		&productCheckout.Product.Brand.Name,
		&productCheckout.Product.Brand.Description,
		&productCheckout.Product.Inventory.ID,
		&productCheckout.Product.Inventory.Quantity,
		&productCheckout.Product.Inventory.Color.ID,
		&productCheckout.Product.Inventory.Color.Color,
	)
	if productImageID.Valid {
		productCheckout.Product.Image.ID = int(productImageID.Int32)
	}
	if productImage.Valid {
		productCheckout.Product.Image.Image = productImage.String
	}
	if err != nil {
		log.Println("Error in for loop:", err)
		return productCheckout, totalPrice, err
	}
	productCheckout.Product.Color.ID = productCheckout.Product.Inventory.Color.ID
	productCheckout.Product.Color.Color = productCheckout.Product.Inventory.Color.Color
	productCheckout.Product.Image.ProductID = productCheckout.Product.ID
	productCheckout.Product.Inventory.ProductID = productCheckout.Product.ID
	totalPrice = productCheckout.Product.Price

	rows, err := user.DB.Query(`SELECT a.address_id,
										a.user_id,
										a.address_name,
										a.address_phone,
										a.address_pincode,
										a.address_housename,
										a.address_streetname,
										a.address_city,
										a.address_state,
										a.address_desc
									FROM address a
										INNER JOIN cart 
											ON a.user_id = cart.user_id 
									WHERE a.user_id = $1;`, productCheckout.UserID)
	if err != nil {
		return productCheckout, totalPrice, err
	}
	defer rows.Close()
	var address models.Address
	var description sql.NullString
	for rows.Next() {
		err = rows.Scan(&address.ID,
			&address.UserID,
			&address.Name,
			&address.Phone,
			&address.Pincode,
			&address.HouseName,
			&address.StreetName,
			&address.City,
			&address.State,
			&description,
		)
		if description.Valid {
			address.Description = description.String
		}
		if err != nil {
			return productCheckout, totalPrice, err
		}
		productCheckout.UserID = address.UserID
		productCheckout.Address = append(productCheckout.Address, address)
	}
	return productCheckout, totalPrice, nil
}

// ------------------------------- Placing Order ------------------------------------

// DBGetPayment method for cash on delivery
func (user Model) DBGetPayment(request models.PaymentRequest) (models.PaymentResponse, error) {

	var payment models.PaymentResponse
	payment.UserID = *request.UserID

	var sqlString string
	var param int
	sqlString = `SELECT p.product_price,`
	if request.ProductID == nil {
		sqlString = fmt.Sprint(sqlString, `cart.quantity,
											cart.user_id,`)
	}
	sqlString = fmt.Sprint(sqlString, `CASE 
											WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN true
											ELSE false
										END AS offer_status,
										CASE 
											WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN (p.product_price - ((p.product_price * o.offer)/100))::FLOAT
											ELSE NULL
										END AS offer_price`)
	if request.ProductID != nil {
		sqlString = fmt.Sprint(sqlString, `	FROM inventories q
												INNER JOIN products p
														ON q.product_id = p.product_id`)
	} else {
		sqlString = fmt.Sprint(sqlString, ` FROM cart
												INNER JOIN products p
														ON cart.product_id = p.product_id`)
	}
	sqlString = fmt.Sprint(sqlString, ` LEFT JOIN category_offers o
												ON o.category_id = p.product_category_id
									WHERE p.product_deleted_at IS NULL
											AND p.product_status = true`)

	if request.ProductID != nil {
		sqlString = fmt.Sprint(sqlString, ` AND q.inventory_id = $1`)
		param = *request.ProductID
	} else {
		sqlString = fmt.Sprint(sqlString, ` AND cart.user_id = $1
											AND cart.deleted_at IS NULL`)
		param = *request.UserID
	}

	rows, err := user.DB.Query(sqlString, param)

	if err != nil {
		log.Println("Error:", err)
		return payment, err
	}
	defer rows.Close()
	for rows.Next() {

		var offerStatus bool
		var offerPrice sql.NullFloat64
		var productPrice float64
		var productQuantity int

		if request.ProductID != nil {
			err = rows.Scan(&productPrice,
				&offerStatus,
				&offerPrice,
			)
			productQuantity = 1
		} else {
			err = rows.Scan(&productPrice,
				&productQuantity,
				&payment.UserID,
				&offerStatus,
				&offerPrice,
			)
		}
		if offerStatus != false {
			if offerPrice.Valid {
				payment.TotalPrice = payment.TotalPrice + (productPrice * float64(productQuantity))

				payment.OfferPrice = payment.OfferPrice + (offerPrice.Float64 * float64(productQuantity))

				payment.Savings = payment.Savings + (productPrice - offerPrice.Float64)

			}
		} else {
			payment.TotalPrice = payment.TotalPrice + (productPrice * float64(productQuantity))

			payment.OfferPrice = payment.OfferPrice + productPrice*float64(productQuantity)
		}
		if err != nil {
			log.Println(err)
			return payment, err
		}
	}
	payment.PaymentType = append(payment.PaymentType, "Cash On Delivery", "Bank")
	return payment, nil
}

// DBPayPayment for paying payment
func (user Model) DBPayPayment(payment models.Payment) (models.Payment, error) {

	if *payment.PaymentType == "COD" {
		// insert into payment table
		err := user.DB.QueryRow(`INSERT INTO payments (user_id,
													total_price,
													payment_type,
													payment_status,
													created_at)
										VALUES ($1, $2, $3, $4, $5) 
										RETURNING payment_id, 
												payment_status `, payment.UserID,
			payment.Amount,
			payment.PaymentType,
			false,
			time.Now()).Scan(&payment.ID,
			&payment.PaymentStatus)
		if err != nil {
			log.Println(err)
			return payment, err
		}
	} else if *payment.PaymentType == "Bank" {
		// insert into payment table
		err := user.DB.QueryRow(`INSERT INTO payments (user_id,
													total_price,
													payment_type,
													payment_status,
													created_at, 
													paid_at)
										VALUES ($1, $2, $3, $4, $5, $6) 
										RETURNING payment_id, 
												payment_status`, payment.UserID,
			payment.Amount,
			payment.PaymentType,
			true,
			time.Now(),
			time.Now()).Scan(&payment.ID,
			&payment.PaymentStatus)
		if err != nil {
			log.Println(err)
			return payment, err
		}
	} else {
		return payment, errors.New("Enter a valid payment type")
	}
	return payment, nil
}

// DBPlaceOrder method for placing an order
func (user Model) DBPlaceOrder(placeOrder models.PlaceOrder) error {

	ctx := context.Background()
	tx, err := user.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	for i := 0; i < len(placeOrder.ProductID); i++ {

		var cartID, quantity sql.NullInt32
		row := tx.QueryRow(`SELECT cart_id,
									quantity 
								FROM cart 
								WHERE user_id = $1 
									AND inventory_id = $2;`, placeOrder.UserID,
			placeOrder.InventoryID)

		err := row.Scan(&cartID,
			&quantity)

		if err == sql.ErrNoRows {
			// Reduce product quantity in inventories
			_, err = tx.Exec(`UPDATE inventories 
									SET product_quantity = product_quantity - $1
								WHERE inventory_id = $2;`, placeOrder.Quantity,
				placeOrder.InventoryID[i],
			)
		} else if err != nil {
			tx.Rollback()
			return err
		}
		if cartID.Valid && quantity.Valid {
			placeOrder.CartID = int(cartID.Int32)
			placeOrder.Quantity = int(quantity.Int32)
			// Reduce product quantity in inventories
			_, err = tx.Exec(`UPDATE inventories 
									SET product_quantity = product_quantity - $1
								WHERE inventory_id = $2;`, placeOrder.Quantity,
				placeOrder.InventoryID[i],
			)

		}

		// if placeOrder.CartID != nil {
		// var timeStatus time.Time
		// var currentTime = time.Now()
		// row := tx.QueryRow(`SELECT CASE WHEN updated_at IS NOT NULL THEN updated_at
		// 								WHEN updated_at IS NULL THEN created_at
		// 							END AS status
		// 						FROM cart
		// 						WHERE cart_id = $1
		// 							AND deleted_at IS NULL; `, placeOrder.CartID[i])
		// err = row.Scan(&timeStatus)
		// if err != nil {
		// 	tx.Rollback()
		// 	return err
		// }
		// duration := currentTime.Sub(timeStatus)
		// if duration > 15 {
		// Reduce product quantity in inventories
		// _, err = tx.Exec(`UPDATE inventories
		// 							SET product_quantity = product_quantity - $1
		// 							WHERE inventory_id = $2;`, placeOrder.Quantity,
		// 	placeOrder.InventoryID[i],
		// )
		// }
		// If it is else case, then no need to reduce product quantity in inventories
		// } else {
		// 	// Reduce product quantity in inventories
		// 	_, err = tx.Exec(`UPDATE inventories
		// 								SET product_quantity = product_quantity - $1
		// 								WHERE inventory_id = $2;`, placeOrder.Quantity,
		// 		placeOrder.InventoryID[i],
		// 	)
		// }

		if err != nil {
			tx.Rollback()
			return err
		}

		// --------------------------- END -------------------------

		// var payment models.Payment
		// row := tx.QueryRow(`SELECT payment_id,
		// 							user_id,
		// 							total_price,
		// 							payment_type,
		// 							payment_status
		// 						FROM payments
		// 						WHERE cancelled_at IS NULL
		// 						AND payment_id = $1;`, placeOrder.PaymentID)
		// err = row.Scan(&payment.ID,
		// 	&payment.UserID,
		// 	&payment.Amount,
		// 	&payment.PaymentType,
		// 	&payment.PaymentStatus,
		// )
		// if err != nil {
		// 	tx.Rollback()
		// 	return err
		// }

		_, err = tx.ExecContext(ctx, `INSERT INTO orders (user_id,
														address_id, 
														product_id, 
														inventory_id, 
														quantity,
														payment_id,
														ordered_at)
								VALUES ($1, $2, $3, $4, $5, $6);`, placeOrder.UserID,
			placeOrder.AddressID,
			placeOrder.ProductID[i],
			placeOrder.InventoryID[i],
			placeOrder.Quantity,
			placeOrder.PaymentID,
			time.Now(),
		)
		if err != nil {
			tx.Rollback()
			return err
		}

		// // I wanna write this above -----------------------------------------------------
		// _, err = tx.ExecContext(ctx, `UPDATE inventories
		// 								SET product_quantity = product_quantity - $1
		// 								WHERE inventory_id = $2;`, placeOrder.Quantity[i],
		// 	placeOrder.InventoryID[i],
		// )
		// if err != nil {
		// 	tx.Rollback()
		// 	return err
		// }

		_, err = tx.ExecContext(ctx, `DELETE FROM cart 
											WHERE cart_id = $1;`, placeOrder.CartID,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ------------------------------- Orders ------------------------------

// DBGetOrders method for accessing orders of a user.
func (user Model) DBGetOrders(userID int) ([]models.Orders, error) {

	var orders []models.Orders
	rows, err := user.DB.Query(`SELECT o.order_id, 
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
									WHERE o.order_status <> 'Delivered' 
										AND p.product_deleted_at IS NULL
										AND o.user_id = $1;`, userID)
	if err != nil {
		return orders, err
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
		log.Println("Order:", orders)
	}
	return orders, nil
}
