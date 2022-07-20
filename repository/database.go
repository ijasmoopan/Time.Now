package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	
	_ "github.com/pressly/goose/v3" 
	// "github.com/ijasmoopan/Time.Now/usecases"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// ConnectDB method
func ConnectDB() *sql.DB {

	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Println("Can't open env file.")
	}
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	connString := fmt.Sprint(driver,"://", user, ":", password, "@", host, ":", port, "/", dbname, "?sslmode=disable")

	db, err := sql.Open(driver, connString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database connected...")
	// defer db.Close()

	return db
}
