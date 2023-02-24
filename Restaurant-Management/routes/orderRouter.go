package routes

import (
	"github.com/gin-gonic/gin"

	controller "Golang-RestaurantManagement/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/Oder", controller.GetOders())
	incomingRoutes.GET("/Oder/Oder_id", controller.GetOder())
	incomingRoutes.POST("/Oder", controller.CreateOder())
	incomingRoutes.PATCH("/Oder/Oder_id", controller.UpdateOder())
}
