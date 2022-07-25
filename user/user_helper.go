package user

import (
	"database/sql"

	"github.com/ijasmoopan/Time.Now/models"
)

// Repo struct describes interfaces for database accessing.
type Repo struct {

	user interface {
		DBGetUser(string)(models.User, error)
		DBAuthUser(string)(models.UserLogin, error)
		DBValidateUser(models.UserLogin)(models.UserLogin, error)
		DBUserRegistration(models.UserRegister)(error)
		DBUpdateUser(models.User)(error)
		DBDeleteUser(string)(error)
	}

	address interface {
		DBGetAddress(int)([]models.Address, error)
		DBAddAddress(models.Address)(error)
		DBUpdateAddress(models.Address)(error)
		DBDeleteAddress(int)(error)
	}

	products interface {
		// DBGetProducts(models.ProductRequest)([]models.Product, error)
		DBGetProducts(models.ProductRequest)([]models.ProductWithInventory, error)
		// DBGetProduct(string)(models.Product, error)
		DBGetAllColorsOfAProduct(int)([]models.Inventories, error)
		DBGetRecommendedProducts(int, int, int)([]models.Product, error)
		// DBGetProductsWithColors(models.ProductRequest)([]models.ProductWithColor, error)
	}

	cart interface {
		DBGetCart(int)([]models.Cart, int, error)
		DBAddCart(models.Cart)(error)
		DBUpdateCart(models.Cart)(error)
		DBDeleteCart(int)(error)
	}

	wishlist interface {
		DBGetWishlist(int)([]models.Wishlist, error)
		DBAddWishlist(models.Wishlist)(error)
		// DBUpdateWishlist(models.Wishlist)(error)
		DBDeleteWishlist(int)(error)
	}

	checkout interface {
		DBCartCheckout(int)(models.CartCheckout, int, float64, error)
		DBProductCheckout(models.ProductCheckout)(models.ProductCheckout, float64, error)

		DBGetPayment(models.PaymentRequest)(models.PaymentResponse, error)
		DBPayPayment(models.Payment)(models.Payment, error)
		DBPlaceOrder(models.PlaceOrder)(error)
	}

	orders interface {
		DBGetOrders(int)([]models.Orders, error)
	}

}

// InterfaceHandler function helps to make dependency injection for connecting database.
func InterfaceHandler(db *sql.DB) *Repo {

	return &Repo{
		user: Model{DB: db},

		address: Model{DB: db},

		products: Model{DB: db},

		cart: Model{DB: db},

		wishlist: Model{DB: db},

		checkout: Model{DB: db},

		orders: Model{DB: db},
	}
}