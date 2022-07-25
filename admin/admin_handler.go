package admin

import (
	// "context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ijasmoopan/Time.Now/models"
)

// ---------------------Admin Side---------------------------

// AdminLogin for sign in admin.
func (repo *Repo) AdminLogin(w http.ResponseWriter, r *http.Request) {

	var adminLogin models.Admin
	json.NewDecoder(r.Body).Decode(&adminLogin)

	adminLogin.Password = hex.EncodeToString(md5.New().Sum([]byte(adminLogin.Password)))

	admin, err := repo.admin.DBGetAdmin(adminLogin)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Incorrect Name or Password",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	id := strconv.Itoa(int(admin.ID))
	token := GeneratingToken(id)

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Successfully Logged In",
	}
	json.NewEncoder(w).Encode(&message)
}

// AdminLogout for logout admin.
func (repo *Repo) AdminLogout(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Successfully Logged Out",
	}
	json.NewEncoder(w).Encode(&message)
}

// GetAdminByID for accessing admin by id.
func (repo *Repo) GetAdminByID(w http.ResponseWriter, r *http.Request) {

	log.Println("GetAdminById...")
	adminID := chi.URLParam(r, "adminID")

	log.Println("Admin Id:", adminID)

	admin, err := repo.admin.DBGetAdminByID(adminID)
	if err != nil {
		log.Println("Error in DBGetAdminById: ", err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Incorrect Name or Password",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	log.Println("Admin by id: ", admin)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Admin by ID",
		"admin": admin,
	}
	json.NewEncoder(w).Encode(&message)
}

// AdminHome for home page of admin.
func (repo *Repo) AdminHome(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Home Page",
	}
	json.NewEncoder(w).Encode(&message)
}



// ----------------------------User Side---------------------------

// GetUsers for giving all users.
func (repo *Repo) GetUsers(w http.ResponseWriter, r *http.Request) {

	var userRequest models.UserRequest

	json.NewDecoder(r.Body).Decode(&userRequest)
	defer r.Body.Close()

	users, err := repo.users.DBGetUsers(userRequest)
	if err != nil {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"users": users,
	}
	json.NewEncoder(w).Encode(&message)
}

// UpdateUser for updating user.
func (repo *Repo) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := repo.users.DBUpdateUser(user)
	if err != nil {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Error": "err",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "User details updated",
	}
	json.NewEncoder(w).Encode(&message)
}

// UpdateUserStatus for updating user status.
func (repo *Repo) UpdateUserStatus(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := repo.users.DBUpdateUserStatus(user.ID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Error": "err",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "User Status updated",
	}
	json.NewEncoder(w).Encode(&message)
}

// DeleteUser for deleting user.
func (repo *Repo) DeleteUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := repo.users.DBDeleteUser(user.ID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Error": "err",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "User Deleted",
	}
	json.NewEncoder(w).Encode(&message)
}





// -------------------Product Side------------------------

// GetProducts for accessing products.
func (repo *Repo) GetProducts(w http.ResponseWriter, r *http.Request) {

	var request models.AdminProductRequest
	json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	products, err := repo.products.DBGetProducts(request)
	if err != nil {
		log.Println("Error:", err)
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Error": "err",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	log.Println("Handler Products:", products)
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Product": products,
	}
	json.NewEncoder(w).Encode(&message)
}

// AddProducts for adding products into database.
func (repo *Repo) AddProducts(w http.ResponseWriter, r *http.Request) {

	var newProduct models.ProductWithInventory
	json.NewDecoder(r.Body).Decode(&newProduct)
	defer r.Body.Close()

	err := repo.products.DBAddProducts(newProduct)
	if err != nil {
		log.Println(err)
		log.Println(err)
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Error": "err",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"response": "New product added",
	}
	json.NewEncoder(w).Encode(&message)
}


// UpdateProducts method for updating product from database.
func (repo *Repo) UpdateProducts(w http.ResponseWriter, r *http.Request){

	var product models.ProductWithInventory
	json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close()

	err := repo.products.DBUpdateProducts(product)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"msg": err,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Product updated",
	}
	json.NewEncoder(w).Encode(&response)
}

// DeleteProducts method for deleting a product from database.
func (repo *Repo) DeleteProducts(w http.ResponseWriter, r *http.Request){

	// var product models.Product
	var product models.ProductDeleteRequest
	json.NewDecoder(r.Body).Decode(&product)
	defer r.Body.Close()

	err := repo.products.DBDeleteProducts(product)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Response": err,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Product deleted",
	}
	json.NewEncoder(w).Encode(&response)
}

// UpdateProductStatus method for updating product status
func (repo *Repo) UpdateProductStatus(w http.ResponseWriter, r *http.Request){

	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	err := repo.products.DBUpdateProductStatus(product)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"msg": err,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Product status updated",
	}
	json.NewEncoder(w).Encode(&response)
}



// ---------------------Category Management------------------------------

// GetCategories method for accessing categories.
func (repo *Repo) GetCategories(w http.ResponseWriter, r *http.Request) {

	var request models.CategoryRequest
	json.NewDecoder(r.Body).Decode(&request)

	category, err := repo.categories.DBGetCategories(request)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Category": category,
	}
	json.NewEncoder(w).Encode(response)
}

// AddCategory for adding category by admin.
func (repo *Repo) AddCategory(w http.ResponseWriter, r *http.Request) {

	var category models.Categories
	json.NewDecoder(r.Body).Decode(&category)

	err := repo.categories.DBAddCategory(category)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Category Added",
	}
	json.NewEncoder(w).Encode(response)
}

// UpdateCategory method for updating category.
func (repo *Repo) UpdateCategory(w http.ResponseWriter, r *http.Request){

	var category models.Categories
	json.NewDecoder(r.Body).Decode(&category)

	err := repo.categories.DBUpdateCategory(category)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Category Updated",
	}
	json.NewEncoder(w).Encode(response)
}

// DeleteCategory method for deleting a category.
func (repo *Repo) DeleteCategory(w http.ResponseWriter, r *http.Request){

	var category models.Categories
	json.NewDecoder(r.Body).Decode(&category)
	defer r.Body.Close()

	err := repo.categories.DBDeleteCategory(category.ID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Category Deleted",
	}
	json.NewEncoder(w).Encode(response)
}



// --------------------Sub Category Management-----------------------

// GetSubcategories method for accessing subcategories.
func (repo *Repo) GetSubcategories(w http.ResponseWriter, r *http.Request) {

	var request models.SubcategoryRequest
	json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	subcategories, err := repo.subcategories.DBGetSubcategories(request)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Subcategory": subcategories,
	}
	json.NewEncoder(w).Encode(response)
}

// AddSubcategory method for adding new subcategory
func (repo *Repo) AddSubcategory(w http.ResponseWriter, r *http.Request) {

	var subcategory models.Subcategories
	json.NewDecoder(r.Body).Decode(&subcategory)
	defer r.Body.Close()

	err := repo.subcategories.DBAddSubcategory(subcategory)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Subcategory Added",
	}
	json.NewEncoder(w).Encode(response)
}

// UpdateSubcategory method for updating subcategory
func (repo *Repo) UpdateSubcategory(w http.ResponseWriter, r *http.Request){

	var subcategory models.Subcategories
	json.NewDecoder(r.Body).Decode(&subcategory)
	defer r.Body.Close()

	err := repo.subcategories.DBUpdateSubcategory(subcategory)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Subcategory Updated",
	}
	json.NewEncoder(w).Encode(response)
}

// DeleteSubcategory method for deleting subcategory.
func (repo *Repo) DeleteSubcategory(w http.ResponseWriter, r *http.Request){

	var subcategory models.Subcategories
	json.NewDecoder(r.Body).Decode(&subcategory)

	err := repo.subcategories.DBDeleteSubcategory(subcategory.ID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Subcategory Deleted",
	}
	json.NewEncoder(w).Encode(response)
}




// ---------------------Brand Management------------------------------

// GetBrands metod for accessing brands from database.
func (repo *Repo) GetBrands(w http.ResponseWriter, r *http.Request) {

	var request models.BrandRequest
	json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	brands, err := repo.brands.DBGetBrands(request)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Brands": brands,
	}
	json.NewEncoder(w).Encode(response)
}

// AddBrand method for adding brand to database.
func (repo *Repo) AddBrand(w http.ResponseWriter, r *http.Request) {

	var brand models.Brands
	json.NewDecoder(r.Body).Decode(&brand)
	defer r.Body.Close()

	err := repo.brands.DBAddBrand(brand)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Brands": "New brand added",
	}
	json.NewEncoder(w).Encode(response)
}

// UpdateBrand method for updating brand.
func (repo *Repo) UpdateBrand(w http.ResponseWriter, r *http.Request){

	var brand models.Brands
	json.NewDecoder(r.Body).Decode(&brand)
	defer r.Body.Close()

	err := repo.brands.DBUpdateBrand(brand)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Brands": "Brand updated",
	}
	json.NewEncoder(w).Encode(response)
}

// DeleteBrand method for deleting brand.
func (repo *Repo) DeleteBrand(w http.ResponseWriter, r *http.Request){

	var brand models.Brands
	json.NewDecoder(r.Body).Decode(&brand)
	defer r.Body.Close()

	err := repo.brands.DBDeleteBrand(brand.ID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Brands": "Brand deleted",
	}
	json.NewEncoder(w).Encode(response)
}



// -----------------------Color Management -------------------------

// GetColors method for accessing colors from a database.
func (repo *Repo) GetColors(w http.ResponseWriter, r *http.Request) {

	var request models.ColorRequest
	json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()

	colors, err := repo.colors.DBGetColors(request)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Colors": colors,
	}
	json.NewEncoder(w).Encode(response)
}

// AddColor method for adding brand to database.
func (repo *Repo) AddColor(w http.ResponseWriter, r *http.Request) {

	var color models.Colors
	json.NewDecoder(r.Body).Decode(&color)
	defer r.Body.Close()

	err := repo.colors.DBAddColor(color)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "New color added",
	}
	json.NewEncoder(w).Encode(response)
}

// UpdateColor method for updating brand.
func (repo *Repo) UpdateColor(w http.ResponseWriter, r *http.Request){

	var color models.Colors
	json.NewDecoder(r.Body).Decode(&color)
	defer r.Body.Close()

	err := repo.colors.DBUpdateColor(color)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Color updated",
	}
	json.NewEncoder(w).Encode(response)
}

// DeleteColor method for deleting brand.
func (repo *Repo) DeleteColor(w http.ResponseWriter, r *http.Request){

	var color models.Colors
	json.NewDecoder(r.Body).Decode(&color)
	defer r.Body.Close()

	err := repo.colors.DBDeleteColor(*color.ID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Error": err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Color deleted",
	}
	json.NewEncoder(w).Encode(response)
}



// --------------------------- Offer Management ---------------------------

// GetOffers method for accessing offers
func (repo *Repo) GetOffers(w http.ResponseWriter, r *http.Request){

	var request models.CategoryOfferRequest
	json.NewDecoder(r.Body).Decode(&request)

	categoryOffer, err := repo.offers.DBGetOffers(request)
	if err != nil {
		response := map[string]interface{}{
			"Response": "Can't fetch offer details",
			"Error": err,
		}
		ResponseBR(w, response)
		return
	}
	response := map[string]interface{}{
		"Response": "Category offers",
		"Offers": categoryOffer,
	}
	ResponseOK(w, response)
}

// AddOffers method for adding category offer.
func (repo *Repo) AddOffers(w http.ResponseWriter, r *http.Request){

	var categoryOffer models.CategoryOfferRequest
	json.NewDecoder(r.Body).Decode(&categoryOffer)
	defer r.Body.Close()

	log.Println("New Offer:", categoryOffer)

	err := repo.offers.DBAddOffers(categoryOffer)
	if err != nil {
		response := map[string]interface{}{
			"Response": "Can't fetch offer details",
			"Error":err,
		}
		ResponseBR(w, response)
		return
	}
	response := map[string]interface{}{
		"Response": "Category Offer Added",
	}
	ResponseOK(w, response)
}

// UpdateOffers method for updating offer
func (repo *Repo) UpdateOffers(w http.ResponseWriter, r *http.Request){

	var offer models.CategoryOfferRequest
	json.NewDecoder(r.Body).Decode(&offer)
	defer r.Body.Close()
	err := repo.offers.DBUpdateOffers(offer)
	if err != nil {
		response := map[string]interface{}{
			"Response": "Can't fetch offer details",
			"Error":err,
		}
		ResponseBR(w, response)
		return
	}
	response := map[string]interface{}{
		"Response": "Category Offer Updated",
	}
	ResponseOK(w, response)
}

// DeleteOffers method for deleting offer
func (repo *Repo) DeleteOffers(w http.ResponseWriter, r *http.Request){

	var offer models.CategoryOfferRequest
	json.NewDecoder(r.Body).Decode(&offer)
	defer r.Body.Close()

	if offer.ID != nil {
		err := repo.offers.DBDeleteOffers(offer.ID)
		if err != nil {
			response := map[string]interface{}{
				"Response": "Can't fetch offer details",
				"Error":err,
			}
			ResponseBR(w, response)
			return
		}
		response := map[string]interface{}{
			"Response": "Category Offer Deleted",
		}
		ResponseOK(w, response)
	} else {
		response := map[string]interface{}{
			"Response": "Insert offer ID",
		}
		ResponseOK(w, response)
	}
}





// ---------------------------- Order Management -------------------------

// GetOrders method for accessing orders
func (repo *Repo) GetOrders(w http.ResponseWriter, r *http.Request){

	var request models.OrderRequest
	json.NewDecoder(r.Body).Decode(&request)

	orders, err := repo.orders.DBGetOrders(request)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Response": "Can't fetch order details",
			"Error": err,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "All orders",
		"Order": orders,
	}
	json.NewEncoder(w).Encode(&response)
}

// ChangeOrderStatus method for changing status of orders
func (repo *Repo) ChangeOrderStatus(w http.ResponseWriter, r *http.Request){

	var order models.Orders
	json.NewDecoder(r.Body).Decode(&order)

	err := repo.orders.DBChangeOrderStatus(order.ID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Response": "Can't fetch order details",
			"Error": err,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": "Order status updated",
	}
	json.NewEncoder(w).Encode(&response)
}