package routes

import (
	// "net/http"

	"net/http"

	"github.com/ijasmoopan/Time.Now/admin"
	"github.com/ijasmoopan/Time.Now/repository"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {

	db := repository.ConnectDB()

	a := admin.InterfaceHandler(db)

	router := chi.NewRouter()


// Admin side routes
// -----------------------Admin Authentication-----------------------

	adminRouter1 := router.Group(nil)

	adminRouter1.Get("/admin", a.AdminLoginPage)
	adminRouter1.Post("/admin/validating", a.AdminLoginValidating)

	adminRouter1.With(a.DeletingJWT, a.IsAdminAuthorized).Get("/admin/logout", a.AdminLogout)



// ----------------------Admin User Management--------------------- 

	adminRouter2 := router.Group(nil)
	adminRouter2.Use(a.IsAdminAuthorized)

	adminRouter2.Get("/admin/home/{adminname}", a.AdminHome)
	adminRouter2.Get("/admin/home/{adminname}/userlist", a.UserListTable)

	adminRouter2.With(a.UserCtx).Get("/admin/home/userlist/{userid}/viewuser", a.ViewUser)

	adminRouter2.With(a.UserCtx).Get("/admin/home/userlist/{userid}/edituser", a.EditUser)
	adminRouter2.Post("/admin/home/userlist/{userid}/editinguser", a.EditingUser)

	adminRouter2.With(a.UserCtx).Get("/admin/home/userlist/{userid}/statususer", a.BlockUser)
	adminRouter2.With(a.UserCtx).Get("/admin/home/userlist/{userid}/deleteuser", a.DeleteUser)


// ----------------------Admin Product Management------------------

	adminRouter2.Get("/admin/home/{adminname}/productlist", a.ProductList)

	adminRouter2.Get("/admin/home/{adminname}/productlist/addproduct", a.AddProduct)
	adminRouter2.Post("/admin/home/{adminname}/productlist/addingproduct", a.AddingProduct)

	adminRouter2.Get("/admin/home/productlist/{product_id}/viewproduct", a.ViewProduct)

	adminRouter2.Get("/admin/home/productlist/{product_id}/editproduct", a.EditProduct)
	adminRouter2.Post("/admin/home/productlist/{product_id}/editingproduct", a.EditingProduct)

	// adminRouter2.Get("/admin/home/productlist/{product_id}/deleteproduct", a.DeleteProduct)




// ---------------------Admin Category Management--------------------

	adminRouter2.Get("/admin/home/{adminname}/categorylist", a.CategoryList)

	adminRouter2.Get("/admin/home/{adminname}/categorylist/addcategory", a.AddCategory)
	adminRouter2.Post("/admin/home/{adminname}/categorylist/addingcategory", a.AddingCategory)

	adminRouter2.Get("/admin/home/categorylist/{category_id}/viewcategoryproducts", a.ViewCategoryProducts)

	adminRouter2.Get("/admin/home/categorylist/{category_id}/editcategory", a.EditCategory)
	adminRouter2.Post("/admin/home/categorylist/{category_id}/editingcategory", a.EditingCategory)

	adminRouter2.Get("/admin/home/categorylist/{category_id}/deletecategory", a.DeleteCategory)



// -------------------Admin Sub Category Management------------------

	adminRouter2.Get("/admin/home/{adminname}/subcategorylist", a.SubcategoryList)

	adminRouter2.Get("/admin/home/{adminname}/subcategorylist/addsubcategory", a.AddSubcategory)
	adminRouter2.Post("/admin/home/{adminname}/subcategorylist/addingsubcategory", a.AddingSubcategory)

	
	adminRouter2.Get("/admin/home/subcategorylist/{subcategory_id}/viewsubcategoryproducts", a.ViewSubcategoryProducts)

	adminRouter2.Get("/admin/home/subcategorylist/{subcategory_id}/editsubcategory", a.EditSubcategory)
	adminRouter2.Post("/admin/home/subcategorylist/{subcategory_id}/editingsubcategory", a.EditingSubcategory)

	adminRouter2.Get("/admin/home/subcategorylist/{subcategory_id}/deletesubcategory", a.DeleteSubcategory)



// --------------------Admin Brand Management------------------------

	adminRouter2.Get("/admin/home/{adminname}/brandlist", a.BrandList)

	adminRouter2.Get("/admin/home/{adminname}/brandlist/addbrand", a.AddBrand)
	adminRouter2.Post("/admin/home/{adminname}/brandlist/addingbrand", a.AddingBrand)

	adminRouter2.Get("/admin/home/brandlist/{brand_id}/viewbrandproducts", a.ViewBrandProducts)

	adminRouter2.Get("/admin/home/brandlist/{brand_id}/editbrand", a.EditBrand)
	adminRouter2.Post("/admin/home/brandlist/{brand_id}/editingbrand", a.EditingBrand)

	adminRouter2.Get("/admin/home/brandlist/{brand_id}/deletebrand", a.DeleteBrand)



	// User side routes
// --------------------------------Account Authentication---------------------------

	// r.Get("/login", user.UserLoginPage)
	// r.Post("/loginvalidating", user.UserLoginValidating)
	// r.Get("/register", user.UserRegistration)
	// r.Post("/registering", user.UserRegistering)
	// r.With(middleware.DeletingJWT).Get("/home/logout", user.UserLogout)
	

// ------------------------------Product Side-------------------------------------
	// r.With(middleware.IsUserAuthorized).Get("/home/{userfirstname}", user.HomePage) 
	// r.Get("/", user.Home)
	// r.Get("/{product_id}/view", user.ProductView)

	// Trail route....
	router.Get("/admin/{admin_id}", a.GetAdminById)
	
	// Admin CSS Files
	router.Handle("/adminstyles/*", http.StripPrefix("/adminstyles/", http.FileServer(http.Dir("./adminTemplates/assets"))))
		
	// User CSS Files
	router.Handle("/userstyles/*", http.StripPrefix("/userstyles/", http.FileServer(http.Dir("./userTemplates/css"))))
	router.Handle("/userjscript/*", http.StripPrefix("/userjscript/", http.FileServer(http.Dir("./userTemplates/js"))))
	router.Handle("/userimg/*", http.StripPrefix("/userimg/", http.FileServer(http.Dir("./userTemplates/images"))))
	router.Handle("/userfont/*", http.StripPrefix("/userfont/", http.FileServer(http.Dir("./userTemplates/fonts"))))

	return router
}

