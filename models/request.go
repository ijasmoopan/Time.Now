package models

// ProductRequest struct describes requests for distributing products.
type ProductRequest struct {
	Product     *string `json:"product"`
	Category    *string `json:"category"`
	Subcategory *string `json:"subcategory"`
	Brand       *string `json:"brand"`
	PriceMin    *int    `json:"priceMin"`
	PriceMax    *int    `json:"priceMax"`
}

// AdminRequest struct describes requests for distrbuting admin.
type AdminRequest struct {
	AdminID   *int    `json:"adminID"`
	AdminName *string `json:"adminName"`
}

// UserRequest struct describes requests for distributing users.
type UserRequest struct {
	UserID *int    `json:"userID"`
	Email  *string `json:"email"`
	Gender *string `json:"gender"`
	Status *bool   `json:"status"`
}

// AdminProductRequest describes requests for distributing products for admin.
type AdminProductRequest struct {
	Product     *string `json:"product"`
	Category    *string `json:"category"`
	Subcategory *string `json:"subcategory"`
	Brand       *string `json:"brand"`
	Color       *string `json:"color"`
	Status      *bool   `json:"status"`
	PriceMin    *int    `json:"priceMin"`
	PriceMax    *int    `json:"priceMax"`
	Quantity    *int    `json:"quantity"`
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
