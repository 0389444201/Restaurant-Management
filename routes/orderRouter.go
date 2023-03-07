package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurant_management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/Oder", controller.GetOrders())
	incomingRoutes.GET("/Oder/Oder_id", controller.GetOrder())
	incomingRoutes.POST("/Oder", controller.CreateOrder())
	incomingRoutes.PATCH("/Oder/Oder_id", controller.UpdateOrder())
}
