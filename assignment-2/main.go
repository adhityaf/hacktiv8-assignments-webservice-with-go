package main

import (
	"assignment-2/controllers"
	"assignment-2/database"
	"assignment-2/repositories"
	"assignment-2/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	route := gin.Default()

	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)

	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	route.POST("/items", itemController.CreateItem)
	route.GET("/items", itemController.GetItems)

	route.POST("/orders", orderController.CreateOrder)
	route.GET("/orders", orderController.GetOrders)
	route.PUT("/orders/:orderId", orderController.UpdateOrder)
	route.DELETE("/orders/:orderId", orderController.DeleteOrder)
	route.Run(database.APP_PORT)
}
