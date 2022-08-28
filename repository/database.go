package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	port, _ := strconv.Atoi(os.Getenv("DB_LOCAL_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")
	// fmt.Println("Connecting..")
	
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	db, err := sql.Open(driver, postgresInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect after 10 secs.")
			break
		}
	}
	fmt.Println("Connected:", db.Ping() == nil)

	return db
}
