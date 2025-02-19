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
	// Load environment variables
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

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Define API routes under /api
	routes.HolidayRoutes(router)

	// Serve static files only in production
	if os.Getenv("ENV") == "production" {
		// Adjust paths to include the "backend" folder, assuming Railway's working directory is the repository root.
		router.Static("/static", "./backend/build/static")
		router.StaticFile("/favicon.ico", "./backend/build/favicon.ico")
		router.StaticFile("/manifest.json", "./backend/build/manifest.json")
		router.NoRoute(func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
				return
			}
			c.File("./backend/build/index.html")
		})
	}

	// Use Railway provided port or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server running on http://0.0.0.0:" + port)
	err := router.Run("0.0.0.0:" + port)
	if err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
