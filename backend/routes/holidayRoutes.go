package routes

import (
	"holiday-api/controllers"

	"github.com/gin-gonic/gin"
)

func HolidayRoutes(router *gin.Engine) {
	holidayRoutes := router.Group("/holidays")
	{
		holidayRoutes.GET("", controllers.ListHolidays)
		holidayRoutes.POST("", controllers.AddHoliday)
		holidayRoutes.DELETE("/:id", controllers.DeleteHoliday)
	}
}
