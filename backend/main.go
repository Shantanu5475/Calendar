package main

import (
	"fmt"
	"holiday-api/config"
	"holiday-api/controllers"
	"holiday-api/routes"
	"net/http"
	"os"
	"strings"

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

	// In production, serve the React app's static assets
	if os.Getenv("ENV") == "production" {
		// Serve static assets from /static (the build folder typically has its assets in "build/static")
		router.Static("/static", "./build/static")
		// Serve other static files individually if needed (like favicon, manifest, etc.)
		router.StaticFile("/favicon.ico", "./build/favicon.ico")
		router.StaticFile("/manifest.json", "./build/manifest.json")

		// Define a NoRoute handler for client-side routing.
		// This will serve index.html for any route that does NOT start with /api.
		router.NoRoute(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
				return
			}
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
