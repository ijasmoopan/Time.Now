package admin

import (
	"database/sql"

	"github.com/ijasmoopan/Time.Now/models"
)

// Repo is a struct that has a bunch of interfaces as fields.
type Repo struct {
	
	admin interface {
		DBGetAdmin(models.Admin)(models.Admin, error)
		DBGetAdminByID(string)(models.Admin, error)
	}

	users interface {
		DBGetUsers(models.UserRequest)([]models.User, error)
		DBUpdateUser(models.User)(error)
		DBUpdateUserStatus(int)(error)
		DBDeleteUser(int)(error)
	}

	products interface {
		DBGetProducts(models.AdminProductRequest)(map[int]models.ProductWithInventory, int, error)
		// DBAddProducts(models.ProductWithInventory)(error)
		DBAddProducts(models.AddProduct)(error)

		DBUpdateProducts(models.ProductWithInventory)(error)
		DBDeleteProducts(models.ProductDeleteRequest)(error)
		DBUpdateProductStatus(models.Product)(error)
	}

	categories interface {
		DBGetCategories(models.CategoryRequest)([]models.Categories, error)
		DBAddCategory(models.Categories)(error)
		DBUpdateCategory(models.Categories)(error)
		DBDeleteCategory(int)(error)
	}

	subcategories interface {
		DBGetSubcategories(models.SubcategoryRequest)([]models.Subcategories, error)
		DBAddSubcategory(models.Subcategories)(error)
		DBUpdateSubcategory(models.Subcategories)(error)
		DBDeleteSubcategory(int)(error)
	}

	brands interface {
		DBGetBrands(models.BrandRequest)([]models.Brands, error)
		DBAddBrand(models.Brands)(error)
		DBUpdateBrand(models.Brands)(error)
		DBDeleteBrand(int)(error)
	}

	colors interface {
		DBGetColors(models.ColorRequest)([]models.Colors, error)
		DBAddColor(models.Colors)(error)
		DBUpdateColor(models.Colors)(error)
		DBDeleteColor(int)(error)
	}

	orders interface {
		DBGetOrders(models.OrderRequest)([]models.Orders, error)
		DBChangeOrderStatus(int)(error)
	}

	offers interface {
		DBGetOffers(models.CategoryOfferRequest)([]models.CategoryOffer, error)
		DBAddOffers(models.CategoryOfferRequest)(error)
		DBUpdateOffers(models.CategoryOfferRequest)(error)
		DBDeleteOffers(*int)(error)
	}
}

// InterfaceHandler for handling interface for db connection.
func InterfaceHandler(db *sql.DB) *Repo {

	return &Repo{

		admin: Model{DB: db},

		users: Model{DB: db},

		products: Model{DB: db},

		categories: Model{DB: db},

		subcategories: Model{DB: db},

		brands: Model{DB: db},

		colors: Model{DB: db},

		orders: Model{DB: db},

		offers: Model{DB: db},
	}
}