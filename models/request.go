package models

// ProductRequest struct describes requests for distributing products.
type ProductRequest struct {
	Product     []*string `json:"product"`
	Category    []*string `json:"category"`
	Subcategory []*string `json:"subcategory"`
	Brand       []*string `json:"brand"`
	Color       []*string `json:"color"`
	PriceMin    *int      `json:"priceMin"`
	PriceMax    *int      `json:"priceMax"`
	UserID      *int      `json:"userID"`
	Wishlist    *bool     `json:"wishlist"`
	OfferMin    *int      `json:"offerMin"`
	OfferMax    *int      `json:"offerMax"`
	Page        *int      `json:"page"`
}

// AdminRequest struct describes requests for distrbuting admin.
type AdminRequest struct {
	AdminID   *int    `json:"adminID"`
	AdminName *string `json:"adminName"`
}

// UserRequest struct describes requests for distributing users.
type UserRequest struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

// AdminProductRequest describes requests for distributing products for admin.
type AdminProductRequest struct {
	Product     string `json:"product"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Brand       string `json:"brand"`
	Color       string `json:"color"`
	Status      string `json:"status"`
	PriceMin    string `json:"priceMin"`
	PriceMax    string `json:"priceMax"`
	Quantity    string `json:"quantity"`
	Page        string `json:"page"`
	Offer       string `json:"offer"`
}

// CategoryRequest describes requests for getting categories for admin.
type CategoryRequest struct {
	CategoryID   *int    `json:"categoryID"`
	CategoryName *string `json:"categoryName"`
}

// SubcategoryRequest describes requests for getting subcategories for admin.
type SubcategoryRequest struct {
	SubcategoryID   *int    `json:"subcategoryID"`
	SubcategoryName *string `json:"subcategoryName"`
}

// BrandRequest describes requests for getting brands for admin.
type BrandRequest struct {
	BrandID   *int    `json:"brandID"`
	BrandName *string `json:"brandName"`
}

// ColorRequest describes request for getting colors for admin.
type ColorRequest struct {
	ColorID *int    `json:"colorID"`
	Color   *string `json:"color"`
}

// OrderRequest describes model for accessing filtered orders.
type OrderRequest struct {
	OrderID     *int     `json:"orderID"`
	UserID      *int     `json:"userID"`
	Product     *string  `json:"product"`
	Category    *string  `json:"category"`
	Subcategory *string  `json:"subcategory"`
	Brand       *string  `json:"brand"`
	OrderStatus *string  `json:"orderStatus"`
	PriceMin    *float64 `json:"priceMin"`
	PriceMax    *float64 `json:"priceMax"`
}
