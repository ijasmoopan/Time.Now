package models

import "time"

// CategoryOffer struct for category offer model.
type CategoryOffer struct {
	ID       int        `json:"offerID"`
	Name     string     `json:"offerName"`
	Offer    int        `json:"offer"`
	From     string     `json:"offerFrom"`
	To       string     `json:"offerTo"`
	Status   bool       `json:"offerStatus"`
	Category Categories `json:"category"`
}

// Offer model for offers
type Offer struct {
	ID       int        `json:"offerID"`
	Name     string     `json:"offerName"`
	Offer    int        `json:"offer"`
	From     string     `json:"offerFrom"`
	To       string     `json:"offerTo"`
	Status   bool       `json:"offerStatus"`
}

// OfferProducts model
type OfferProducts struct {
	ID                int       `json:"offerID"`
	Name              string    `json:"offerName"`
	Offer             int       `json:"offer"`
	From              time.Time `json:"offerFrom"`
	To                time.Time `json:"offerTo"`
	Status            bool      `json:"offerStatus"`
	ProductID         int       `json:"ProductID"`
	ProductName       string    `json:"ProductName"`
	ProductCategoryID int       `json:"categoryID"`
	ProductCategory   string    `json:"category"`
	ProductPrice      float64   `json:"productPrice"`
	OfferPrice        float64   `json:"offerPrice"`
}

// CategoryOfferRequest model for accessing offers
type CategoryOfferRequest struct {
	ID         *int    `json:"offerID"`
	Name       *string `json:"offerName"`
	Offer      *int    `json:"offer"`
	CategoryID *int    `json:"categoryID"`
	OfferFrom  *string `json:"offerFrom"`
	OfferTo    *string `json:"offerTo"`
}

// ProductOffer struct for product offer model.
type ProductOffer struct {
	ID      int
	Name    string
	Offer   int
	From    time.Time
	To      time.Time
	Product Product
}

// Coupon struct for coupon offer model.
type Coupon struct {
	ID    int
	Name  string
	Code  string
	Offer int
	From  time.Time
	To    time.Time
}
