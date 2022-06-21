package models

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	Product_id int `json:"product_id" gorm:"type:bigserial;primaryKey;autoIncrement"`
	Product_name string `json:"user_firstname" gorm:"type:varchar(255);not null"`
	Product_price decimal.Decimal `json:"product_price" sql:"type:decimal(10,2);"`
	Product_color string `json:"product_color" gorm:"type:varchar(20);not null"`
	Product_desc string `json:"product_desc" gorm:"type:varchar(255);not null"`
	Product_quantity uint `json:"product_quantity" gorm:"not null"`
	Product_image string `json:"product_image" gorm:"type:varchar(255);"`
}


