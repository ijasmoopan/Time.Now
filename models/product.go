package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	Product_id          int             `json:"product_id"`
	Product_name        string          `json:"user_firstname"`
	Product_desc        string          `json:"product_desc"`
	Product_price       decimal.Decimal `json:"product_price"`
	Product_created_at  time.Time       `json:"created_at"`
	Product_updated_at  *time.Time      `json:"updated_at"`
	Product_deleted_at  *time.Time      `json:"deleted_at"`
	Product_image       Images          `json:"product_image"`
	Product_brand       Brands          `json:"product_brand"`
	Product_category    Categories      `json:"product_category"`
	Product_subcategory Subcategories   `json:"product_subcategory"`
	Product_inventory   Inventories     `json:"product_inventory"`
	// Product_color    Colors          `json:"product_color"`
}

type Inventories struct {
	Inventory_id         int        `json:"inventory_id"`
	Product_id           int        `json:"product_id"`
	Product_color        string     `json:"product_color"`
	Product_quantity     int        `json:"product_quantity"`
	Inventory_created_at time.Time  `json:"inventory_created_at"`
	Inventory_updated_at *time.Time `json:"inventory_updated_at"`
	Inventory_deleted_at *time.Time `json:"inventory_deleted_at"`
}

type ProductInDetail struct {
	Product_id               int
	Product_name             string
	Product_desc             string
	Product_price            decimal.Decimal
	Product_brand         int
	Product_brand_name       string
	Product_category_id      int
	Product_category_name    string
	Product_subcategory_id   int
	Product_subcategory_name string
	Product_inventory_id     int
	Product_quantity         int
	Color_id                 int
	Color                    string
}

type ListProduct struct {
	Product_id          int
	Product_name        string
	Product_desc        string
	Product_price       decimal.Decimal
	Product_inventory   Inventories
	Product_brand       Brands
	Product_category    Categories
	Product_subcategory Subcategories
	// Product_color       Colors
}

type AddProduct struct {
	Product_id             int
	Product_name           string
	Product_category_id    int
	Product_brand_id       int
	Product_subcategory_id int
	Product_price          decimal.Decimal
	Product_desc           string
	Product_quantity       int
	Product_color          string
}

type SampleProduct struct {
	Product_id    int
	Product_name  string
	Product_desc  string
	Product_price decimal.Decimal
}

// type Colors struct {
// 	Color_id   int    `json:"color_id"`
// 	Color      string `json:"color"`
// 	Product_id int    `json:"product_id"`
// }

type Images struct {
	Image_id      int    `json:"image_id"`
	Product_id    int    `json:"product_id"`
	Product_image string `json:"product_image"`
}

type Brands struct {
	Brand_id         int        `json:"brand_id"`
	Brand_name       string     `json:"brand_name"`
	Brand_desc       string     `json:"brand_desc"`
	Brand_created_at time.Time  `json:"brand_created_at"`
	Brand_updated_at *time.Time `json:"brand_updated_at"`
	Brand_deleted_at *time.Time `json:"brand_deleted_at"`
}

type Categories struct {
	Category_id         int        `json:"category_id"`
	Category_name       string     `json:"category_name"`
	Category_desc       string     `json:"category_desc"`
	Category_created_at time.Time  `json:"category_created_at"`
	Category_updated_at *time.Time `json:"category_updated_at"`
	Category_deleted_at *time.Time `json:"category_deleted_at"`
}

type Subcategories struct {
	Subcategory_id         int        `json:"subcategory_id"`
	Subcategory_name       string     `json:"subcategory_name"`
	Subcategory_desc       string     `json:"subcategory_desc"`
	Subcategory_created_at time.Time  `json:"subcategory_created_at"`
	Subcategory_updated_at *time.Time `json:"subcategory_updated_at"`
	Subcategory_deleted_at *time.Time `json:"subcategory_deleted_at"`
}

