package models

import (
	"time"
)

type User struct {
	User_id         int        `json:"user_id"`
	User_firstname  string     `json:"user_firstname"`
	User_secondname string     `json:"user_secondname"`
	User_email      string     `json:"user_email"`
	User_password   string     `json:"user_password"`
	User_phone      string     `json:"user_phone"`
	User_gender     string     `json:"user_gender"`
	User_status     bool       `json:"status"`
	User_referral   string     `json:"user_referral"`
	Created_at      time.Time  `json:"created_at"`
	Updated_at      *time.Time `json:"updated_at"`
	Deleted_at      *time.Time `json:"deleted_at"`
	Address_id      Address    `json:"address_id"`
}

type Address struct {
	address_id         int        `json:"address_id"`
	address_name       string     `json:"address_name"`
	address_phone      string     `json:"addess_phone"`
	address_pincode    string     `json:"addess_pincode"`
	address_housename  string     `json:"addess_housename"`
	address_streetname string     `json:"addess_streetname"`
	address_city       string     `json:"addess_city"`
	address_state      string     `json:"addess_state"`
	address_desc       string     `json:"addess_desc"`
	address_created_at time.Time  `json:"addess_created_at"`
	address_updated_at *time.Time `json:"addess_updated_at"`
	address_deleted_at *time.Time `json:"addess_deleted_at"`
}

type UserLogin struct {
	User_id       int
	User_email    string
	User_password string
}

type UserRegister struct {
	User_id         int
	User_firstname  string
	User_secondname string
	User_email      string
	User_gender     string
	User_phone      string
	User_referral   string
	User_password   string
	User_confirm    string
}
