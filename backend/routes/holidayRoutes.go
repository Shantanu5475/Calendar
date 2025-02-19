package routes

import (
	"holiday-api/controllers"

	"github.com/gin-gonic/gin"
)

func HolidayRoutes(router *gin.Engine) {
	// Group the API endpoints under /api to avoid conflicts with the static file serving.
	api := router.Group("/api")
	{
		api.GET("/holidays", controllers.ListHolidays)
		api.POST("/holidays", controllers.AddHoliday)
		api.DELETE("/holidays/:id", controllers.DeleteHoliday)
	}
}
