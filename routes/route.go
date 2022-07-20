package routes

import (
	"github.com/ijasmoopan/Time.Now/admin"
	"github.com/ijasmoopan/Time.Now/repository"
	"github.com/ijasmoopan/Time.Now/user"
	"github.com/go-chi/chi/v5"
)

// Router function for handling routes.
func Router() *chi.Mux {

	db := repository.ConnectDB()
	admindb := admin.InterfaceHandler(db)
	userdb := user.InterfaceHandler(db)

	router := chi.NewRouter()

// Admin side routes
// -----------------------Admin Authentication-----------------------

	router.Post("/admin/login", admindb.AdminLogin)
	router.With(admindb.DeletingJWT, admindb.IsAdminAuthorized).Post("/admin/logout", admindb.AdminLogout)


// ----------------------Admin User Management--------------------- 

	adminRouter := router.Group(nil)
	adminRouter.Use(admindb.IsAdminAuthorized)

	adminRouter.Get("/admin/{adminName}", admindb.AdminHome)
	
	adminRouter.Get("/admin/users", admindb.GetUsers)
	adminRouter.Patch("/admin/users", admindb.UpdateUser)
	adminRouter.Delete("/admin/users", admindb.DeleteUser)

	adminRouter.Patch("/admin/users/status", admindb.UpdateUserStatus)


// ----------------------Admin Product Management------------------

	adminRouter.Get("/admin/products", admindb.GetProducts)
	adminRouter.Post("/admin/products", admindb.AddProducts)
	adminRouter.Patch("/admin/products", admindb.UpdateProducts)
	adminRouter.Delete("/admin/products", admindb.DeleteProducts)

	adminRouter.Patch("/admin/products/status", admindb.UpdateProductStatus)

// ---------------------Admin Category Management--------------------

	adminRouter.Get("/admin/categories", admindb.GetCategories)
	adminRouter.Post("/admin/categories", admindb.AddCategory)
	adminRouter.Patch("/admin/categories", admindb.UpdateCategory)
	adminRouter.Delete("/admin/categories", admindb.DeleteCategory)


// -------------------Admin Sub Category Management------------------

	adminRouter.Get("/admin/subcategories", admindb.GetSubcategories)
	adminRouter.Post("/admin/subcategories", admindb.AddSubcategory)
	adminRouter.Patch("/admin/subcategories", admindb.UpdateSubcategory)
	adminRouter.Delete("/admin/subcategories", admindb.DeleteSubcategory)


// --------------------Admin Brand Management------------------------

	adminRouter.Get("/admin/brands", admindb.GetBrands)
	adminRouter.Post("/admin/brands", admindb.AddBrand)
	adminRouter.Patch("/admin/brands", admindb.UpdateBrand)
	adminRouter.Delete("/admin/brands", admindb.DeleteBrand)


// --------------------Admin Color Management------------------------

	adminRouter.Get("/admin/colors", admindb.GetColors)
	adminRouter.Post("/admin/colors", admindb.AddColor)
	adminRouter.Patch("/admin/colors", admindb.UpdateColor)
	adminRouter.Delete("/admin/colors", admindb.DeleteColor)


// ------------------------- Admin Category Offer Management ------------------------

	adminRouter.Get("/admin/offers", admindb.GetOffers)


	// adminRouter.Get("/admin/home/offerlist/productoffer", admindb.ProductOffer)
	// CRUD

	// adminRouter.Get("/admin/home/offerlist/categoryoffer", admindb.CategoryOffer)
	// CRUD

	// adminRouter.Get("/admin/home/offerlist/subcategoryoffer", admindb.SubcategoryOffer)
	// CRUD


// ---------------------------- Admin Order Management --------------------------------

	adminRouter.Get("/admin/orders", admindb.GetOrders)
	adminRouter.Post("/admin/orders/status", admindb.ChangeOrderStatus)
	


// -------------------------User Home Page------------------------------------

	router.Get("/products", userdb.GetProducts)

	router.Get("/products/{product_id}", userdb.HomeSingleProduct)

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


// // ------------------------User Cart Management-----------------------------

// 	userRouter.Get("/user/cart", userdb.UserCart)
// 	userRouter.Post("/user/{product_id}-{inventory_id}-{quantity}/addtocart", userdb.AddToCart)

// 	userRouter.Patch("/user/cart/{cart_id}-{inventory_id}-{quantity}/updatecartproduct", userdb.UpdateProductFromCart)
// 	userRouter.Delete("/user/cart/{cart_id}/deletecart", userdb.DeleteProductFromCart)



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

// 	userRouter.Get("/user/checkout/payment/paypal", userdb.PayPal)
// 	userRouter.Get("/user/checkout/payment/cod", userdb.CashOnDelivery)


// ----------------------------- Order Management ------------------------------

	userRouter.Post("/user/checkout/placeorder", userdb.PlaceOrder)

	userRouter.Get("/user/orders", userdb.GetOrders)
	userRouter.Get("/user/cancelorder", userdb.CancelOrder)








	// Trail route....
	// router.Get("/admin/{admin_id}", admindb.GetAdminById)
	
	// Admin CSS Files
	// router.Handle("/adminstyles/*", http.StripPrefix("/adminstyles/", http.FileServer(http.Dir("./adminTemplates/assets"))))
		
	// User CSS Files
	// router.Handle("/userstyles/*", http.StripPrefix("/userstyles/", http.FileServer(http.Dir("./userTemplates/css"))))
	// router.Handle("/userjscript/*", http.StripPrefix("/userjscript/", http.FileServer(http.Dir("./userTemplates/js"))))
	// router.Handle("/userimg/*", http.StripPrefix("/userimg/", http.FileServer(http.Dir("./userTemplates/images"))))
	// router.Handle("/userfont/*", http.StripPrefix("/userfont/", http.FileServer(http.Dir("./userTemplates/fonts"))))

	return router
}

