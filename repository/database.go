package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	// "time"

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
	// host := os.Getenv("DB_CONTAINER")
	port := os.Getenv("DB_LOCAL_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	connString := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable", "postgresql", user, password, host, port, dbname)

	db, err := sql.Open(driver, connString)
	if err != nil {
		fmt.Println(err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("DB String: ", connString)
		fmt.Println("Failed to connect database:", err)
		return nil
	}

	// start := time.Now()
	// for db.Ping() != nil {
	// 	if start.After(start.Add(10 * time.Second)) {
	// 		log.Fatalln("Failed to connect db after 10 secs.")
	// 		break
	// 	}
	// }
	fmt.Println("Database connection: ", db.Ping() == nil)

	// defer db.Close()

	return db
}
