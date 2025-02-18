package main

import (
	"holiday-api/config"
	"holiday-api/controllers"
	"holiday-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()

	holidayCollection := config.GetCollection("holidays")
	controllers.InitHolidayCollection(holidayCollection)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	routes.HolidayRoutes(router)

	// Start the server on port 8080
	router.Run(":8080")
}
