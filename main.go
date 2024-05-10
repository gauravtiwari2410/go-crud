package main

import (
	"fiber-crud/config" // Import the config package
	"log"

	"fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to the database
	db, err := config.DBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform auto-migration
	config.AutoMigrate(db)

	app := fiber.New()
	app.Use(app)
	routes.Setup(app)
	app.Listen(":3001")

}
