package admin

import (
	"database/sql"

	"github.com/ijasmoopan/Time.Now/models"
)

// Repo is a struct that has a bunch of interfaces as fields.
// @property admin - interface{}
// @property users - This is a slice of User structs.
// @property adminbyid - This is a property that is used to get an admin by their id.
// @property adminbyname - This is a property that returns an interface that has a method called
// DBGetAdminByName.
// @property user - This is a property that is used to get a user by their ID.
type Repo struct {
	
	admin interface {
		DBGetAdmin(models.Admin)(models.Admin, error)
	}
	
	adminbyid interface {
		DBGetAdminById(string)(models.Admin, error)
	}
	adminbyname interface {
		DBGetAdminByName(string)(models.Admin, error)
	}

	users interface {
		DBGetUsers()([]models.User, error)
	}
	user interface {
		DBGetUser(string)(models.User, error)
	}
	updateuser interface {
		DBUpdatingUser(string, models.User)(error)
	}
	userStatus interface {
		DBGetUserStatus(string)(error)
	}
	deleteUser interface {
		DBDeleteUser(string)(error)
	}



	products interface {
		DBGetAllProducts()([]models.SampleProduct, error)
	}
	product interface {
		DBGetProduct(string)(models.ListProduct, error)
	}
	productcolors interface {
		DBGetProductColors(string)([]models.Inventories, error)
	}
	addproduct interface {
		DBAddProduct(models.ListProduct)(error)
	}
	editproduct interface {
		DBEditProduct(models.ListProduct)(error)
	}
	deleteproduct interface {
		DBDeleteProduct(string)(error)
	}



	categories interface {
		DBGetAllCategories()([]models.Categories, error)
	}
	category interface {
		DBGetCategory(string)(models.Categories, error)
	}
	addcategory interface {
		DBAddCategory(models.Categories)(error)
	}
	editcategory interface {
		DBEditCategory(models.Categories)(error)
	}
	deletecategory interface{
		DBDeleteCategory(string)(error)
	}
	categoryproducts interface {
		DBGetCategoryProducts(string)([]models.SampleProduct, error)
	}



	subcategories interface {
		DBGetAllSubcategories()([]models.Subcategories, error)
	}
	subcategory interface {
		DBGetSubcategory(string)(models.Subcategories, error)
	}
	addsubcategory interface {
		DBAddSubcategory(models.Subcategories)(error)
	}
	editsubcategory interface {
		DBEditSubcategory(models.Subcategories)(error)
	}
	deletesubcategory interface{
		DBDeleteSubcategory(string)(error)
	}
	subcategoryproducts interface {
		DBGetSubcategoryProducts(string)([]models.SampleProduct, error)
	}



	brands interface {
		DBGetAllBrands()([]models.Brands, error)
	}
	brand interface {
		DBGetBrand(string)(models.Brands, error)
	}
	addbrand interface {
		DBAddBrand(models.Brands)(error)
	}
	editbrand interface {
		DBEditBrand(models.Brands)(error)
	}
	deletebrand interface{
		DBDeleteBrand(string)(error)
	}
	brandproducts interface {
		DBGetBrandProducts(string)([]models.SampleProduct, error)
	}
}

func InterfaceHandler(db *sql.DB) *Repo {

	return &Repo{
		admin: AdminModel{DB: db},
		adminbyid: AdminModel{DB: db},
		adminbyname: AdminModel{DB: db},

		user: AdminModel{DB: db},
		users: AdminModel{DB: db},
		updateuser: AdminModel{DB: db},
		userStatus: AdminModel{DB: db},
		deleteUser: AdminModel{DB: db},


		product: AdminModel{DB: db},
		products: AdminModel{DB: db},
		productcolors: AdminModel{DB: db},
		addproduct: AdminModel{DB: db},
		editproduct: AdminModel{DB: db},
		// deleteproduct: AdminModel{DB: db},

		category: AdminModel{DB: db},
		categories: AdminModel{DB: db},
		addcategory: AdminModel{DB: db},
		editcategory: AdminModel{DB: db},
		deletecategory: AdminModel{DB: db},
		categoryproducts: AdminModel{DB: db},

		subcategory: AdminModel{DB: db},
		subcategories: AdminModel{DB: db},
		addsubcategory: AdminModel{DB: db},
		editsubcategory: AdminModel{DB: db},
		deletesubcategory: AdminModel{DB: db},
		subcategoryproducts: AdminModel{DB: db},

		brand: AdminModel{DB: db},
		brands: AdminModel{DB: db},
		addbrand: AdminModel{DB: db},
		editbrand: AdminModel{DB: db},
		deletebrand: AdminModel{DB: db},
		brandproducts: AdminModel{DB: db},
	}
}