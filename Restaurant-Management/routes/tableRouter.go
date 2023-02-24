package routes

import (
	"github.com/gin-gonic/gin"

	controller "Golang-RestaurantManagement/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/table_id", GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/table_id", controller.UpdateTable())
}
