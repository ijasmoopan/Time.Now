package user

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ijasmoopan/Time.Now/models"
)

// ----------------------Home-------------------------

// GetProducts method will give products corresponding to request params.
func (repo *Repo) GetProducts(w http.ResponseWriter, r *http.Request){

	var request models.ProductRequest
	json.NewDecoder(r.Body).Decode(&request)

	products, err := repo.products.DBGetProducts(request)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"Response": http.StatusBadRequest,
			"Error": err,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"Response": http.StatusOK,
		"products": products,
	}
	json.NewEncoder(w).Encode(&response)
}

// HomeSingleProduct will get single product, its colors and it's similar products.
func (repo *Repo) HomeSingleProduct(w http.ResponseWriter, r *http.Request){

	var request models.ProductRequest
	json.NewDecoder(r.Body).Decode(&request)

	product, err := repo.products.DBGetProducts(request)
	if err != nil {
		log.Println(err)
		message := map[string]interface{}{
			"msg": "Single Product Details",
			"error": err,
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		
		json.NewEncoder(w).Encode(&message)
		return
	}

	colors, err := repo.products.DBGetAllColorsOfAProduct(product[0].ID)
	if err != nil {
		log.Println(err)
		message := map[string]interface{}{
			"msg": "Single Product Colors",
			"error": err,
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	// ----------------------------- Similar Products---------------------------
	recommendations, err := repo.products.DBGetRecommendedProducts(product[0].ID, product[0].Category.ID, product[0].Subcategory.ID)
	if err == sql.ErrNoRows {
		log.Println(err)
		message := map[string]interface{}{
			"msg": "Single Product Details",
			"error": "No recommended products",
			"product": product,
			"colors": colors,
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&message)
	}
	if err != nil {
		log.Println(err)
		message := map[string]interface{}{
			"msg": "Can't fetch recommended products",
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
	}

	message := map[string]interface{}{
		"msg": "Single Product Details",
		"product": product,
		"colors": colors,
		"recommendations": recommendations,
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&message)
}



// ---------------------User-------------------------------

// UserLogin describes login for user.
func (repo *Repo) UserLogin(w http.ResponseWriter, r *http.Request){

	var loginUser models.UserLogin
	json.NewDecoder(r.Body).Decode(&loginUser)
	defer r.Body.Close()

	log.Println("Login user:", loginUser)
	loginUser.Password = hex.EncodeToString(md5.New().Sum([]byte(loginUser.Password)))
	log.Println(loginUser)
	
	user, err := repo.user.DBValidateUser(loginUser)
	if err != nil {
		log.Println(err)
	}
	if user.Password != loginUser.Password {
		message := map[string]interface{}{
			"response": http.StatusBadRequest,
			"msg": "Incorrect Username or Password",
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}

	id := strconv.Itoa(user.ID)
	token := GeneratingToken(id)
	
	cookie := http.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	log.Println("Token generated..Redirecting to Home...")

	message := map[string]interface{}{
		"response": http.StatusOK,
		"msg": "User Validated",
		"user": user.Email,
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&message)
}

// UserRegistration for registering a new user by their own.
func (repo *Repo) UserRegistration(w http.ResponseWriter, r *http.Request){

	var newUser models.UserRegister
	json.NewDecoder(r.Body).Decode(&newUser)

	if newUser.Password != newUser.ConfirmPassword {
		message := map[string]interface{}{
			"error": "Different passwords",
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}
	newUser.Password = hex.EncodeToString(md5.New().Sum([]byte(newUser.Password)))
	
	if newUser.Referral != nil {
		// -----------------------------------------------------------

		// Check it is someone's referral or not...I want to do this..Add points to wallet...
	} else {
		newuuid := uuid.New()
		newUser.Referral = &newuuid
	}
		
	err := repo.user.DBUserRegistration(newUser)
	if err != nil {
		log.Println(err)
		message := map[string]interface{}{
			"response": "User already exist",
			"error": err,
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&message)
		return
	}
	message := map[string]interface{}{
		"msg": "Registration Completed",
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&message)
}

// UserLogout for logoutting user account.
func (repo *Repo) UserLogout(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Logged Out Successfully",
	}
	json.NewEncoder(w).Encode(&message)
}

// UserProfile is for providing profile page for user.
func (repo *Repo) UserProfile(w http.ResponseWriter, r *http.Request){
	
	userID := chi.URLParam(r, "userID")
	log.Println("UserID:", userID)
	user, err := repo.user.DBGetUser(userID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't access user details",
		}
		json.NewEncoder(w).Encode(&message)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"user details": user,
	}
	json.NewEncoder(w).Encode(&message)
}

// UpdateUserProfile for updating user details.
func (repo *Repo) UpdateUserProfile(w http.ResponseWriter, r *http.Request){

	userID := chi.URLParam(r, "userID")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID, _ = strconv.Atoi(userID)

	err := repo.user.DBUpdateUser(user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch user details",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "user details updated",
	}
	json.NewEncoder(w).Encode(&message)
}

// DeleteUserProfile for deleting user account for users.
func (repo *Repo) DeleteUserProfile(w http.ResponseWriter, r *http.Request){

	userID := chi.URLParam(r, "userID")
	
	err := repo.user.DBDeleteUser(userID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch user details",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "user deleted",
	}
	json.NewEncoder(w).Encode(&message)
}



// -----------------------User Address Management-----------------------------

// GetUserAddress method for accessing address.
func (repo *Repo) GetUserAddress(w http.ResponseWriter, r *http.Request){

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	address, err := repo.address.DBGetAddress(user.ID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Response": "Can't fetch any address",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Response": "User Address",
		"Address": address,
	}
	json.NewEncoder(w).Encode(&message)
}

// AddUserAddress method for adding new address.
func (repo *Repo) AddUserAddress(w http.ResponseWriter, r *http.Request){

	var address models.Address
	json.NewDecoder(r.Body).Decode(&address)

	err := repo.address.DBAddAddress(address)
	if err != nil {
		log.Println("handler", err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch any address",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Addresses Added",
	}
	json.NewEncoder(w).Encode(&message)
}

// UpdateUserAddress method for updating address
func (repo *Repo) UpdateUserAddress(w http.ResponseWriter, r *http.Request){

	var address models.Address
	json.NewDecoder(r.Body).Decode(&address)

	err := repo.address.DBUpdateAddress(address)
	if err != nil {
		log.Println("handler", err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch any address",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Address Updated",
	}
	json.NewEncoder(w).Encode(&message)
}

// DeleteUserAddress method for deleting address.
func (repo *Repo) DeleteUserAddress(w http.ResponseWriter, r *http.Request){

	var address models.Address
	json.NewDecoder(r.Body).Decode(&address)

	err := repo.address.DBDeleteAddress(address.ID)
	if err != nil {
		log.Println("handler", err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Response": "Can't fetch any address",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Response": "Address Deleted",
	}
	json.NewEncoder(w).Encode(&message)
}



// ---------------------Cart Management------------------------

// AddCart method to add product to cart
func (repo *Repo) AddCart(w http.ResponseWriter, r *http.Request){

	var cart models.Cart
	json.NewDecoder(r.Body).Decode(&cart)
	defer r.Body.Close()

	err := repo.cart.DBAddCart(cart)
	if err != nil {
		log.Println("Error:", err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch cart details",
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Response": "Added to cart",
	}
	json.NewEncoder(w).Encode(&message)
}

// GetCart method access cart of the user.
func (repo *Repo) GetCart(w http.ResponseWriter, r *http.Request){

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	cartproducts, countOfProducts, err := repo.cart.DBGetCart(user.ID)
	if len(cartproducts) == 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		message := map[string]interface{}{
			"Response": "No products in cart",
			"Cart Products": cartproducts,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Response": "Can't fetch cart details",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Response": "All products in cart",
		"Cart Products": cartproducts,
		"Products in cart": countOfProducts,
	}
	json.NewEncoder(w).Encode(&message)
}

// UpdateCart method for updating product quantity from cart.
func (repo *Repo) UpdateCart(w http.ResponseWriter, r *http.Request){

	var cart models.Cart
	json.NewDecoder(r.Body).Decode(&cart)

	err := repo.cart.DBUpdateCart(cart)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch cart details",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Cart Updated",
	}
	json.NewEncoder(w).Encode(&message)
}

// DeleteCart method to delete product from cart.
func (repo *Repo) DeleteCart(w http.ResponseWriter, r *http.Request){

	var cart models.Cart
	json.NewDecoder(r.Body).Decode(&cart)

	err := repo.cart.DBDeleteCart(cart.ID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch cart details",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Removed the product from cart",
	}
	json.NewEncoder(w).Encode(&message)
}




// ----------------------------Wishlist Management---------------------------------

// GetWishlist method for accessing user wishlist.
func (repo *Repo) GetWishlist(w http.ResponseWriter, r *http.Request){

	var wishlist models.Wishlist
	json.NewDecoder(r.Body).Decode(&wishlist)
	defer r.Body.Close()

	wishlistproducts, err := repo.wishlist.DBGetWishlist(wishlist.UserID)
	if err == sql.ErrNoRows {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		message := map[string]interface{}{
			"Response": "No products in wishlist",
			"Wishlist Products": wishlistproducts,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Response": "Can't fetch wishlist details",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "All products in wishlist",
		"Wishlist Products": wishlistproducts,
		"Products in wishlist": len(wishlistproducts),
	}
	json.NewEncoder(w).Encode(&message)
}

// AddWishlist method for adding products into wishist.
func (repo *Repo) AddWishlist(w http.ResponseWriter, r *http.Request){

	var wishlist models.Wishlist
	json.NewDecoder(r.Body).Decode(&wishlist)

	err := repo.wishlist.DBAddWishlist(wishlist)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Response": "Can't fetch wishlist details",
			"Error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Response": "Product added to wishlist",
	}
	json.NewEncoder(w).Encode(&message)
}

// DeleteWishlist method for deleting a product from wishlist.
func (repo *Repo) DeleteWishlist(w http.ResponseWriter, r *http.Request){

	var wishlist models.Wishlist
	json.NewDecoder(r.Body).Decode(&wishlist)

	err := repo.wishlist.DBDeleteWishlist(wishlist.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch wishlist details",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Product removed from wishlist",
	}
	json.NewEncoder(w).Encode(&message)
}



// ----------------------------------User Checkout-----------------------------------------

// CartCheckout method for cart product checkout
func (repo *Repo) CartCheckout(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	cartCheckout, countOfProducts, totalPrice, err := repo.checkout.DBCartCheckout(user.ID) 
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Response": "Can't fetch checkout data",
			"Error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Response": "Product checkout",
		"Products": countOfProducts,
		"Total Price": totalPrice,
		"Products for checkout": cartCheckout,
	}
	json.NewEncoder(w).Encode(&message)
}

// ProductCheckout method for instant buy
func (repo *Repo) ProductCheckout(w http.ResponseWriter, r *http.Request){

	var productCheckout models.ProductCheckout
	json.NewDecoder(r.Body).Decode(&productCheckout)

	productCheckout, totalPrice, err := repo.checkout.DBProductCheckout(productCheckout)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Response": "Can't fetch checkout data",
			"Error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Response": "Product checkout",
		"Product for checkout": productCheckout,
		"Total Price": totalPrice,
	}
	json.NewEncoder(w).Encode(&message)
}




// ------------------------- Payment Management ------------------------

// func (repo *Repo) PayPal(w http.ResponseWriter, r *http.Request){
// }

// CashOnDelivery method for COD payment
func (repo *Repo) CashOnDelivery(w http.ResponseWriter, r *http.Request){

	var COD models.COD
	json.NewDecoder(r.Body).Decode(&COD)

	

}



// --------------------------Order Management------------------------------

// PlaceOrder method for ordering placing.
func (repo *Repo) PlaceOrder(w http.ResponseWriter, r *http.Request){

	var placeOrder models.PlaceOrder
	json.NewDecoder(r.Body).Decode(&placeOrder)

	err := repo.checkout.DBPlaceOrder(placeOrder)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"msg": "Can't fetch any Products",
			"error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"msg": "Order Placed",
	}
	json.NewEncoder(w).Encode(&message)
}


// ----------------------------- Orders -----------------------------

// GetOrders method for accessing all orders
func (repo *Repo) GetOrders(w http.ResponseWriter, r *http.Request){

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	orders, err := repo.orders.DBGetOrders(user.ID)
	if err != nil {
		log.Println(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		message := map[string]interface{}{
			"Response": "Can't fetch order details",
			"Error": err,
		}
		json.NewEncoder(w).Encode(&message)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := map[string]interface{}{
		"Response": "All Orders",
		"Orders": orders,
	}
	json.NewEncoder(w).Encode(&message)
}

// CancelOrder method cancelling an order.
func (repo *Repo) CancelOrder(w http.ResponseWriter, r *http.Request){

	// var order models.Order

	
}