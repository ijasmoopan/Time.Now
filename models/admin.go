package models

type Admin struct {

	Admin_id int `json:"admin_id" gorm:"type:bigserial;primaryKey;autoIncrement"`
	Admin_name string `json:"admin_name" gorm:"type:varchar(255);not null"`
	Admin_password string `json:"admin_password" gorm:"type:varchar(255);not null"`   

}