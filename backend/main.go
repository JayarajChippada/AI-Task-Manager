package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"backend/config"
	"backend/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	// Initialize Fiber
	app := fiber.New()

	// Connect to MongoDB
	config.ConnectDB()

	// Register routes
	routes.SetupAuthRoutes(app)
	routes.SetupTaskRoutes(app)
	routes.SetupAIRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(app.Listen(":" + port))
}
