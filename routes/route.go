package routes

import (
	"net/http"

	"github.com/ijasmoopan/Time.Now/admin"
	"github.com/ijasmoopan/Time.Now/user"
	"github.com/ijasmoopan/Time.Now/middleware"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	// Admin side routes
// -----------------------Admin Authentication-----------------------

	r.Get("/admin", admin.AdminLoginPage)
	r.Post("/adminvalidating", admin.AdminLoginValidating)
	r.With(middleware.IsAdminAuthorized).Get("/admin/home/{adminname}", admin.AdminHome)
	

// ----------------------Admin User Listing-------------------------

	r.With(middleware.IsAdminAuthorized).Get("/admin/userlist", admin.UserListTable)
	r.Route("/admin/userlist/{user_id}", func(r chi.Router){
		r.Use(middleware.UserCtx)
		r.Get("/view", admin.ViewUser)
		r.Get("/edit", admin.EditUser)
		r.Delete("/delete", admin.DeleteUser)
	})
	r.With(middleware.IsAdminAuthorized).Post("/admin/userlist/{user_id}/editing", admin.EditingUser)
	r.With(middleware.IsAdminAuthorized).Get("/admin/userlist/{user_id}/status", admin.BlockUser)
	r.With(middleware.IsAdminAuthorized).Get("/admin/userlist/{user_id}/delete", admin.DeleteUser)

	r.With(middleware.DeletingJWT).Get("/admin/logout", admin.AdminLogout)
	


// ---------------------Admin Product Side----------------------

	r.With(middleware.IsAdminAuthorized).Get("/admin/productlist", admin.ProductList)
	r.With(middleware.IsAdminAuthorized).Get("/admin/productlist/{product_id}/view", admin.ProductView)
	r.With(middleware.IsAdminAuthorized).Get("/admin/productlist/{product_id}/edit", admin.EditProduct)
	r.With(middleware.IsAdminAuthorized).Get("/admin/productlist/{product_id}/editing", admin.EditingProduct)




	// User side routes
// --------------------------------Account Authentication---------------------------

	r.Get("/login", user.UserLoginPage)
	r.Post("/loginvalidating", user.UserLoginValidating)
	r.Get("/register", user.UserRegistration)
	r.Post("/registering", user.UserRegistering)
	r.With(middleware.DeletingJWT).Get("/home/logout", user.UserLogout)
	

// ------------------------------Product Side-------------------------------------
	r.With(middleware.IsUserAuthorized).Get("/home/{userfirstname}", user.HomePage) 
	r.Get("/", user.Home)
	r.Get("/{product_id}/view", user.ProductView)



	// Admin CSS Files
	r.Handle("/adminstyles/*", http.StripPrefix("/adminstyles/", http.FileServer(http.Dir("./adminTemplates/assets"))))
		
	// User CSS Files
	r.Handle("/userstyles/*", http.StripPrefix("/userstyles/", http.FileServer(http.Dir("./userTemplates/css"))))
	r.Handle("/userjscript/*", http.StripPrefix("/userjscript/", http.FileServer(http.Dir("./userTemplates/js"))))
	r.Handle("/userimg/*", http.StripPrefix("/userimg/", http.FileServer(http.Dir("./userTemplates/images"))))
	r.Handle("/userfont/*", http.StripPrefix("/userfont/", http.FileServer(http.Dir("./userTemplates/fonts"))))

	return r
}

