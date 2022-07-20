package models

import (
	"time"

	"github.com/google/uuid"
)

// Address is a model struct used to describing address of users.
type Address struct {
	ID          int        `json:"addressID"`
	UserID      int        `json:"userID"`
	Name        string     `json:"Name"`
	Phone       string     `json:"Phone"`
	Pincode     string     `json:"Pincode"`
	HouseName   string     `json:"Housename"`
	StreetName  string     `json:"Street"`
	City        string     `json:"City"`
	State       string     `json:"State"`
	Description string     `json:"Description"`
}

// AddressInDetail for detailed details about address.
type AddressInDetail struct {
	ID          int        `json:"addressID"`
	UserID      int        `json:"userID"`
	Name        string     `json:"Name"`
	Phone       string     `json:"Phone"`
	Pincode     string     `json:"Pincode"`
	HouseName   string     `json:"Housename"`
	StreetName  string     `json:"Streetname"`
	City        string     `json:"City"`
	State       string     `json:"State"`
	Description string     `json:"Desc"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	UpdatedAt   *time.Time `json:"UpdatedAt"`
	DeletedAt   *time.Time `json:"DeletedAt"`
}

// User model is for describing struct for user.
type User struct {
	ID         int        `json:"userID"`
	FirstName  string     `json:"firstname"`
	SecondName string     `json:"secondname"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Gender     string     `json:"gender"`
	Status     bool       `json:"status"`
	Referral   string     `json:"referral"`
	Image      string     `json:"image"`
}

// UserInDetail for detailed details of user.
type UserInDetail struct {
	ID         int        `json:"userID"`
	FirstName  string     `json:"firstname"`
	SecondName string     `json:"secondname"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Gender     string     `json:"gender"`
	Status     bool       `json:"status"`
	Referral   string     `json:"referral"`
	Address    Address    `json:"address"`
	Image      string     `json:"image"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

// UserLogin is a struct model for describing user when logging time.
type UserLogin struct {
	ID       int    `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRegister is a struct model used for describing user while registration.
type UserRegister struct {
	ID              int        `json:"userID"`
	FirstName       string     `json:"firstname"`
	SecondName      string     `json:"secondname"`
	Email           string     `json:"email"`
	Gender          string     `json:"gender"`
	Phone           string     `json:"phone"`
	Referral        *uuid.UUID `json:"referral"`
	Password        string     `json:"password"`
	ConfirmPassword string     `json:"confirm"`
}
