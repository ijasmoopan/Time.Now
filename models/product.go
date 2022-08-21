package models

import (
	"github.com/shopspring/decimal"
	"time"
)

// Product describes details of products.
type Product struct {
	ID          int           `json:"productID"`
	Name        string        `json:"productName"`
	Description string        `json:"productDesc"`
	Price       float64       `json:"productPrice"`
	OfferPrice  float64       `json:"offerPrice"`
	Status      bool          `json:"productStatus"`
	Color       Colors        `json:"productColor"`
	Image       Images        `json:"productImage"`
	Brand       Brands        `json:"productBrand"`
	Category    Categories    `json:"productCategory"`
	Subcategory Subcategories `json:"productSubcategory"`
	Inventory   Inventories   `json:"productInventory"`
	Offer       CategoryOffer `json:"offer"`
}

// ProductWithInventory describes details of products.
type ProductWithInventory struct {
	ID          int           `json:"productID"`
	Name        string        `json:"productName"`
	Description string        `json:"productDesc"`
	Price       float64       `json:"productPrice"`
	OfferPrice  *float64      `json:"offerPrice"`
	Status      bool          `json:"productStatus"`
	Color       []Colors      `json:"productColor"`
	Image       Images        `json:"productImage"`
	Brand       Brands        `json:"productBrand"`
	Category    Categories    `json:"productCategory"`
	Subcategory Subcategories `json:"productSubcategory"`
	Inventory   []Inventories `json:"productInventory"`
	Offer       CategoryOffer `json:"offer"`
	Wishlist    *bool         `json:"wishlist"`
}

// AddProduct struct for adding a product.
type AddProduct struct {
	ID            int
	Name          string
	Description   string
	Price         float64
	Status        bool
	Inventories   string
	Image         string
	CategoryID    int
	BrandID       int
	SubcategoryID int
}

// Products describes details of products.
type Products struct {
	ID          int           `json:"productID"`
	Name        string        `json:"productName"`
	Description string        `json:"productDesc"`
	Price       float64       `json:"productPrice"`
	Status      bool          `json:"productStatus"`
	Color       Colors        `json:"productColor"`
	Image       Images        `json:"productImage"`
	Brand       Brands        `json:"productBrand"`
	Category    Categories    `json:"productCategory"`
	Subcategory Subcategories `json:"productSubcategory"`
	Inventory   Inventories   `json:"productInventory"`
}

// ProductDeleteRequest struct for deleting a product or its inventory
type ProductDeleteRequest struct {
	ID          *int `json:"productID"`
	ColorID     *int `json:"colorID"`
	InventoryID *int `json:"inventoryID"`
	ImageID     *int `json:"imageID"`
}

// ProductWithColor describes details of products.
type ProductWithColor struct {
	ID          int           `json:"productID"`
	Name        string        `json:"productName"`
	Description string        `json:"productDesc"`
	Price       float64       `json:"productPrice"`
	Status      bool          `json:"productStatus"`
	Color       []Colors      `json:"productColor"`
	Image       Images        `json:"productImage"`
	Brand       Brands        `json:"productBrand"`
	Category    Categories    `json:"productCategory"`
	Subcategory Subcategories `json:"productSubcategory"`
}

// ProductInDetail describes details of products.
type ProductInDetail struct {
	ID          int           `json:"product_id"`
	Name        string        `json:"product_name"`
	Description string        `json:"product_desc"`
	Price       float64       `json:"product_price"`
	Status      bool          `json:"product_status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   *time.Time    `json:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at"`
	Color       Colors        `json:"product_color"`
	Image       Images        `json:"product_image"`
	Brand       Brands        `json:"product_brand"`
	Category    Categories    `json:"product_category"`
	Subcategory Subcategories `json:"product_subcategory"`
	Inventory   Inventories   `json:"product_inventory"`
}

// HomeProduct describes product when printing details in front-end.
type HomeProduct struct {
	ID          int           `json:"product_id"`
	Name        string        `json:"product_name"`
	Description string        `json:"product_desc"`
	Price       float64       `json:"product_price"`
	Status      bool          `json:"product_status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   *time.Time    `json:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at"`
	Color       Colors        `json:"product_color"`
	Image       Images        `json:"product_image"`
	Brand       Brands        `json:"product_brand"`
	Category    Categories    `json:"product_category"`
	Subcategory Subcategories `json:"product_subcategory"`
}

// Colors struct describes colors of product.
type Colors struct {
	ID    *int    `json:"colorID"`
	Color *string `json:"color"`
}

// Inventories struct describes inventory details of products.
type Inventories struct {
	ID        int    `json:"inventoryID"`
	ProductID int    `json:"productID"`
	Color     Colors `json:"productColor"`
	Quantity  int    `json:"productQuantity"`
}

// InventoriesInDetail struct describes inventory details of products.
type InventoriesInDetail struct {
	ID        int        `json:"inventoryID"`
	ProductID int        `json:"productID"`
	Color     Colors     `json:"productColorID"`
	Quantity  int        `json:"productQuantity"`
	CreatedAt time.Time  `json:"inventory_created_at"`
	UpdatedAt *time.Time `json:"inventory_updated_at"`
	DeletedAt *time.Time `json:"inventory_deleted_at"`
}

// ProductInDetails struct for describing detailed details about products.
type ProductInDetails struct {
	ID              int
	Name            string
	Description     string
	Price           float64
	Status          bool
	BrandID         int
	BrandName       string
	CategoryID      int
	CategoryName    string
	SubcategoryID   int
	SubcategoryName string
	InventoryID     int
	Quantity        int
	ColorID         int
	Color           string
}

// ListProduct struct for listing some details of products.
type ListProduct struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Status      bool
	Inventory   Inventories
	Brand       Brands
	Category    Categories
	Subcategory Subcategories
	// Product_color       Colors
}

// SampleProduct struct describing basic details of products.
type SampleProduct struct {
	ID          int
	Name        string
	Description string
	Status      bool
	Price       decimal.Decimal
}

// Images struct for describing images of products.
type Images struct {
	ID        int    `json:"imageID"`
	ProductID int    `json:"productID"`
	Image     string `json:"productImage"`
}

// Brands struct for describing brands of products.
type Brands struct {
	ID          int    `json:"brandID"`
	Name        string `json:"brandName"`
	Description string `json:"brandDesc"`
}

// BrandsInDetail struct for describing brands of products.
type BrandsInDetail struct {
	ID          int        `json:"brandID"`
	Name        string     `json:"brandName"`
	Description string     `json:"brandDesc"`
	CreatedAt   time.Time  `json:"brand_created_at"`
	UpdatedAt   *time.Time `json:"brand_updated_at"`
	DeletedAt   *time.Time `json:"brand_deleted_at"`
}

// Categories struct for describing categories of products.
type Categories struct {
	ID          int    `json:"categoryID"`
	Name        string `json:"categoryName"`
	Description string `json:"categoryDesc"`
}

// CategoriesInDetail struct for describing categories of products.
type CategoriesInDetail struct {
	ID          int        `json:"category_id"`
	Name        string     `json:"category_name"`
	Description string     `json:"category_desc"`
	CreatedAt   time.Time  `json:"category_created_at"`
	UpdatedAt   *time.Time `json:"category_updated_at"`
	DeletedAt   *time.Time `json:"category_deleted_at"`
}

// Subcategories struct for describing subcategories of products.
type Subcategories struct {
	ID          int    `json:"subcategoryID"`
	Name        string `json:"subcategoryName"`
	Description string `json:"subcategoryDesc"`
}

// SubcategoriesInDetail struct for describing subcategories of products.
type SubcategoriesInDetail struct {
	ID          int        `json:"subcategory_id"`
	Name        string     `json:"subcategory_name"`
	Description string     `json:"subcategory_desc"`
	CreatedAt   time.Time  `json:"subcategory_created_at"`
	UpdatedAt   *time.Time `json:"subcategory_updated_at"`
	DeletedAt   *time.Time `json:"subcategory_deleted_at"`
}
