package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	dbURL := "postgres://postgres:ijasmoopan@localhost:5432/timenow"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db

}

func CloseDB(db *gorm.DB) {
	
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()

}