package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"time"

	_ "github.com/pressly/goose/v3"

	// "github.com/ijasmoopan/Time.Now/usecases"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


// ConnectDB method
func ConnectDB() *sql.DB {

	err := godotenv.Load("app.env")
	if err != nil {
		log.Println("Can't open env file.")
	}
	driver := os.Getenv("DB_DRIVER")
	log.Println("Driver:", driver)
	
	connString := "postgresql://root:secret@postgres:5432/timenow?sslmode=disable"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect after 10 secs.")
			break
		}
	}
	fmt.Println("Connected:", db.Ping() == nil)

	// Database migration using goose package.
	// goose.SetDialect("postgres")
	// goose.SetBaseFS(embedMigrations)

	// if err = goose.Up(db, "github.com/ijasmoopan/Time.Now/migrations"); err != nil {
	// 	panic(err)
	// }
	// if err = goose.Version(db, "github.com/ijasmoopan/Time.Now/migrations"); err != nil {
	// 	log.Fatal(err)
	// }

	return db
}
