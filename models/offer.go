package models

import "time"

// CategoryOffer struct for category offer model.
type CategoryOffer struct {
	ID int 
	Name string
	Offer int
	From time.Time
	To time.Time
	Category Categories
}

// ProductOffer struct for product offer model.
type ProductOffer struct {
	ID int
	Name string
	Offer int
	From time.Time
	To time.Time
	Product Product
}

// Coupon struct for coupon offer model.
type Coupon struct {
	ID int
	Name string
	Code string
	Offer int
	From time.Time
	To time.Time
}