package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
       
var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(" ⚠️ Error loading .env file")

	} else {
		log.Println(" .env file loaded successfully")
	}

	// Build DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Connect to DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Fatal("❌ Failed to connect to database: ", err)

	}

	DB = db
	fmt.Println("✅ Connected to MySQL database")

}
