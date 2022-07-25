package models

// Cart struct for describing cart details.
type Cart struct {
	ID       int     `json:"cartID"`
	UserID   int     `json:"userID"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

// CartInDetail struct for describing cart details.
type CartInDetail struct {
	ID       int     `json:"cartID"`
	User     User    `json:"user"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

// AddToCart struct for describing cart while adding products into cart.
type AddToCart struct {
	ID          int `json:"userID"`
	ProductID   int `json:"productID"`
	InventoryID int `json:"inventoryID"`
	Quantity    int `json:"quantity"`
}

// Wishlist struct describing wishlist details.
type Wishlist struct {
	ID      int     `json:"wishlistID"`
	UserID  int     `json:"userID"`
	Product Product `json:"product"`
}

// WishlistInDetail struct describing wishlist details.
type WishlistInDetail struct {
	ID      int     `json:"wishlistID"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}

// AddToWishlist struct for describing wishlist while adding products into wishlist.
type AddToWishlist struct {
	UserID      int `json:"userID"`
	ProductID   int `json:"productID"`
	InventoryID int `json:"inventoryID"`
}


