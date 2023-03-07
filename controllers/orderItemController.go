package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"restaurant_management/database"
	"restaurant_management/models"
	"time"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})
			return
		}
		var allOrderItems []bson.M
		if err := result.All(ctx, &allOrderItems); err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}
func GetOrderItemByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")
		allOrderItems, err := ItemsByOrder(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occrured while listing order items by order"})
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}
func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var orderItemPack OrderItemPack
		var order models.Order
		if err := c.BindJSON(&orderItemPack); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		order.Order_Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		orderItemToBeInserted := []interface{}{}
		order.Table_id = orderItemPack.Table_id
		order_id := OrderItemOrderCreator(order)
		for _, orderItem := range orderItemPack.Order_items {
			orderItem.Order_id = order_id
			validationErr := validate.Struct(orderItem)
			if validationErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
				return
			}
			order.ID = primitive.NewObjectID()
			orderItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Order_item_id = orderItem.ID.Hex()
			var num = toFixed(*orderItem.Unit_price, 2)
			orderItem.Unit_price = &num
			orderItemToBeInserted = append(orderItemToBeInserted, orderItem)
		}
		insertedOrderItems, err := orderItemCollection.InsertMany(ctx, orderItemToBeInserted)
		if err != nil {
			log.Fatal(err)
		}
		defer cancel()
		c.JSON(http.StatusOK, insertedOrderItems)
	}
}
func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	matchStage := bson.D{
		{"$match", bson.D{
			{"order_id", id}}}}
	lookupStage := bson.D{
		{"$lookup", bson.D{
			{"form", "food"},
			{"localField", "food_id"},
			{"foreignField", "food_id"},
			{"as", "food"}}}}
	unwindStage := bson.D{
		{"$unwind", bson.D{
			{"path", "$food"},
			{"preserveNullAndEmptyArrays", true}}}}
	lookupOrderStage := bson.D{
		{"$lookup", bson.D{
			{"from", "order"},
			{"localField", "order_id"},
			{"foreignField", "order_id"},
			{"as", "order"}}}}
	unwindOrderStage := bson.D{
		{"$unwind", bson.D{
			{"path", "$order"},
			{"preserveNullEmptyArrays", true}}}}
	lookupTableStage := bson.D{
		{"$lookup", bson.D{
			{"from", "table"},
			{"localField", "order.table_id"},
			{"foreignField", "table_id"},
			{"as", "table"}}}}
	unwindTableStage := bson.D{
		{"$unwind", bson.D{
			{"path", "$table"},
			{"preserveNullAndEmptyArrays", true}}}}
	projectStage := bson.D{
		{"$project", bson.D{
			{"id", 0},
			{"amont", "$food_price"},
			{"total_cont", 1},
			{"food_name", "$food.name"},
			{"foo_image", "$food.food_image"},
			{"table_number", "$table.table_number"},
			{"table_id", "table.table_id"},
			{"order_id", "$order.order_id"},
			{"price", "$food.price"},
			{"quantity", 1},
		}},
	}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", bson.D{
				{"order_id", "$order_id"},
				{"table_id", "$table_id"},
				{"table_number", "$table_number"},
				{"payment_due", bson.D{
					{"$sum", "$amount"},
				}}}}, {"total_cont", bson.D{
				{"$sum", 1},
				{"order_items", "$order_items"},
			}},
		}}}
	projectStage2 := bson.D{
		{"$project", bson.D{
			{"id", 0},
			{"", ""},
			{"payment_due", 1},
			{"total_count", 1},
			{"table_number", "$_id.table_number"},
			{"order_time", 1},
		}},
	}
	result, err := orderItemCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		lookupStage,
		unwindStage,
		lookupOrderStage,
		unwindOrderStage,
		lookupTableStage,
		unwindTableStage,
		projectStage,
		groupStage,
		projectStage2,
	})
	if err != nil {
		panic(err)
	}
	if err = result.All(ctx, &OrderItems); err != nil {
		panic(err)
	}
	defer cancel()
	return OrderItems, err
}
func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		orderItemId := c.Param("order_item_id")
		var orderItem models.OrderItem
		err := orderItemCollection.FindOne(ctx, bson.M{"order_item_id": orderItemId})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing ordered item"})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var orderItem models.OrderItem
		orderItemId := c.Param("order_item_id")
		filter := bson.M{"order_item_id": orderItemId}
		var updateObj primitive.D
		if orderItem.Unit_price != nil {
			updateObj = append(updateObj, bson.E{"unit_price", *&orderItem.Unit_price})
		}
		if orderItem.Quantity != nil {
			updateObj = append(updateObj, bson.E{"quantity", *orderItem.Quantity})
		}
		if orderItem.Food_id != nil {
			updateObj = append(updateObj, bson.E{"food_id", *orderItem.Food_id})
		}
		orderItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", orderItem.Updated_at})
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := orderItemCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)
		if err != nil {
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
