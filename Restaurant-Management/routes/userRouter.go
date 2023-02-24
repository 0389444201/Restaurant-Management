package routes

import (
	controller "Golang-RestaurantManagement/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/user", controller.GetUsers())
	incomingRoutes.GET("/user/user_id", controller.GetUser())
	incomingRoutes.POST("/user/signup", controller.SignUp())
	incomingRoutes.POST("/user/login", controller.Login())
}
