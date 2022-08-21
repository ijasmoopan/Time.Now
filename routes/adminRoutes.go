package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/ijasmoopan/Time.Now/admin"
	"github.com/ijasmoopan/Time.Now/repository"
)

// Router function for handling routes.
func adminRouter(router *chi.Mux) *chi.Mux {

	db := repository.ConnectDB()
	admindb := admin.InterfaceHandler(db)

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

	// ------------------------- Admin Offer Management ------------------------

	adminRouter.Get("/admin/offers", admindb.GetOffers)
	adminRouter.Post("/admin/offers", admindb.AddOffers)
	adminRouter.Patch("/admin/offers", admindb.UpdateOffers)
	adminRouter.Delete("/admin/offers", admindb.DeleteOffers)

	// ---------------------------- Admin Order Management --------------------------------

	adminRouter.Get("/admin/orders", admindb.GetOrders)
	adminRouter.Post("/admin/orders/status", admindb.ChangeOrderStatus)

	return router
}