package models

import (
	"time"

	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	User_id int `json:"user_id" gorm:"type:bigserial;primaryKey;autoIncrement"`
	User_firstname string `json:"user_firstname" gorm:"type:varchar(255);not null"`
	User_secondname string `json:"user_secondname" gorm:"type:varchar(255);not null"`
	User_email string `json:"user_email" gorm:"type:varchar(255);not null"`
	User_password string `json:"user_password" gorm:"type:varchar(255);not null"`
	User_phone string `json:"user_phone" gorm:"type:varchar(10);not null"`
	User_gender string `json:"user_gender" gorm:"type:varchar(10);not null"`
	User_status bool `json:"status" gorm:"type:bool;default:true"`
	User_referral string `json:"user_referral" gorm:"type:uuid;default:uuid_generate_v4()"`
	Created_at time.Time `json:"created_at" gorm:"default:now()"`
	Updated_at *time.Time `json:"updated_at" gorm:"default:NULL"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
}

type UserLogin struct {
	User_id int
	User_email string
	User_password string
}

type UserRegister struct {
	User_id int
	User_firstname string
	User_secondname string
	User_email string
	User_gender string
	User_phone string
	User_referral string
	User_password string
	User_confirm string
}