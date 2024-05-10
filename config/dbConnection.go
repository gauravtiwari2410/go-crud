package config

import (
	"fiber-crud/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql" // Import GORM's MySQL driver
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	// Automatically migrate your schema
	err := db.Debug().AutoMigrate(
		&model.Cashier{},
		&model.Category{},
		&model.Discount{},
		&model.Order{},
		&model.Payment{},
		&model.PaymentType{},
		&model.Product{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func DBConnection() (*gorm.DB, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to the database using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object from GORM: %v", err)
	}

	// Check if the connection is successful by pinging the MySQL server
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to the database using GORM!")
	return db, nil
}
