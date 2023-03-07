package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"restaurant_management/database"
	"restaurant_management/middleware"
	"restaurant_management/routes"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	os.Setenv("PORT", "localhost:8000")
	port := os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	router.Use(middleware.Authentication())
	router.Run(port)
}
