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
	fmt.Println("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("❌ Failed to load environment variables:", err)
		return
	}
	fmt.Println("Environment variables loaded successfully!")

	// Connect to MongoDB
	fmt.Println("Connecting to MongoDB...")
	err = config.ConnectDB()
	if err != nil {
		fmt.Println("❌ Failed to connect to MongoDB:", err)
		return
	}
	fmt.Println("MongoDB connected successfully!")

	// Initialize the holiday collection
	fmt.Println("Initializing holiday collection...")
	holidayCollection := config.GetCollection("holidays")
	controllers.InitHolidayCollection(holidayCollection)
	fmt.Println("Holiday collection initialized successfully!")

	// Create a new Gin router
	fmt.Println("Creating Gin router...")
	router := gin.Default()
	fmt.Println("Gin router created successfully!")

	// Configure CORS
	fmt.Println("Configuring CORS...")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))
	fmt.Println("CORS configured successfully!")

	// Define routes
	fmt.Println("Defining routes...")
	routes.HolidayRoutes(router)
	fmt.Println("Routes defined successfully!")

	// Get port from environment variable (Render provides this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Using PORT:", port)

	// Start the server
	fmt.Println("Starting server...")
	err = router.Run("0.0.0.0:" + port)
	if err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
