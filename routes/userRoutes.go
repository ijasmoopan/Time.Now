package routes

import (
	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	// "github.com/ijasmoopan/Time.Now/admin"
	"github.com/ijasmoopan/Time.Now/repository"
	"github.com/ijasmoopan/Time.Now/user"
)

// Router function for handling routes.
func userRouter(router *chi.Mux) *chi.Mux {

	db := repository.ConnectDB()
	userdb := user.InterfaceHandler(db)

	// -------------------------User Home Page------------------------------------

	homeRouter := router.Group(nil)
	homeRouter.Use(userdb.IsHomeUserAuthorized)

	homeRouter.Get("/products", userdb.GetProducts)

	homeRouter.Get("/products/{product_id}", userdb.HomeSingleProduct)

	router.Post("/login", userdb.UserLogin)
	router.Post("/signup", userdb.UserRegistration)
	router.With(userdb.DeleteToken).Post("/logout", userdb.UserLogout)

	userRouter := router.Group(nil)
	userRouter.Use(userdb.IsUserAuthorized)

	userRouter.Get("/user/{userID}/user", userdb.UserProfile)
	userRouter.Patch("/user/{userID}/user", userdb.UpdateUserProfile)
	userRouter.Delete("/user/{userID}/user", userdb.DeleteUserProfile)

	// -------------------------------User Address Management------------------------

	userRouter.Get("/user/address", userdb.GetUserAddress)
	userRouter.Post("/user/address", userdb.AddUserAddress)
	userRouter.Patch("/user/address", userdb.UpdateUserAddress)
	userRouter.Delete("/user/address", userdb.DeleteUserAddress)

	// ------------------------User Cart Management-----------------------------

	userRouter.Get("/user/cart", userdb.GetCart)
	userRouter.Post("/user/cart", userdb.AddCart)
	userRouter.Patch("/user/cart", userdb.UpdateCart)
	userRouter.Delete("/user/cart", userdb.DeleteCart)

	// -----------------------User Wishlist Management---------------------------

	userRouter.Get("/user/wishlist", userdb.GetWishlist)
	userRouter.Post("/user/wishlist", userdb.AddWishlist)
	userRouter.Delete("/user/wishlist", userdb.DeleteWishlist)

	// --------------------------------User Checkout-------------------------------

	userRouter.Get("/user/cart/checkout", userdb.CartCheckout)
	userRouter.Get("/user/product/checkout", userdb.ProductCheckout)

	// ---------------------------- Payment Management ---------------------------

	userRouter.Get("/user/checkout/payment", userdb.GetPayment)
	userRouter.Post("/user/checkout/payment", userdb.PayPayment)

	// ----------------------------- Order Management ------------------------------

	userRouter.Post("/user/checkout/placeorder", userdb.PlaceOrder)

	userRouter.Get("/user/orders", userdb.GetOrders)
	userRouter.Get("/user/cancelorder", userdb.CancelOrder)

	// Trail route....
	// router.Get("/admin/{admin_id}", admindb.GetAdminById)

	return router
}