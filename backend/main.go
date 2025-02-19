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
	config.ConnectDB()
	fmt.Println("🚀 Connected to MongoDB successfully!")
	fmt.Println("MongoDB connected successfully!")

	// Initialize the holiday collection
	holidayCollection := config.GetCollection("holidays")
	controllers.InitHolidayCollection(holidayCollection)

	// Create a new Gin router
	router := gin.Default()

	// Configure CORS (if testing locally with separate origins)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Define API routes under the /api prefix
	routes.HolidayRoutes(router)

	// Serve static files from the "build" folder when in production.
	if os.Getenv("ENV") == "production" {
		// Serve all files from the build folder at the root.
		router.Static("/", "./build")
		// For client-side routing (React Router): serve index.html for unmatched routes.
		router.NoRoute(func(c *gin.Context) {
			c.File("./build/index.html")
		})
	}

	// Get port from environment variable (Railway will provide this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local development
	}

	fmt.Println("🚀 Server running on http://0.0.0.0:" + port)
	fmt.Println("Starting server...")
	err := router.Run("0.0.0.0:" + port)
	if err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
