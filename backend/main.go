package main

import (
	"fmt"
	"holiday-api/config"
	"holiday-api/controllers"
	"holiday-api/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables (useful for local development)
	godotenv.Load()

	// Connect to MongoDB
	fmt.Println("Connecting to MongoDB...")
	if err := config.ConnectDB(); err != nil {
		fmt.Println("❌ Failed to connect to MongoDB:", err)
		os.Exit(1) // Exit the program if MongoDB connection fails
	}
	fmt.Println("MongoDB connected successfully!")

	// Initialize the holiday collection
	holidayCollection := config.GetCollection("holidays")
	controllers.InitHolidayCollection(holidayCollection)

	// if os.Getenv("ENV")=="production"{
	// 	app.static("/","./cl")
	// }

	// Create a new Gin router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Define routes
	routes.HolidayRoutes(router)

	// Get port from environment variable (Render provides this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 for local development
	}

	// Print the port before starting the server
	fmt.Println("🚀 Server running on http://0.0.0.0:" + port)

	// Start the server
	fmt.Println("Starting server...")
	err := router.Run("0.0.0.0:" + port) // Bind to the port provided by Render
	if err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
