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
	// Initialize database
	db, err := database.ConnectDB("mysql")
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	route := gin.Default()

	itemRepo := repositories.NewItemRepo(db)
	orderRepo := repositories.NewOrderRepo(db)

	orderService := services.NewOrderService(orderRepo, itemRepo)
	orderController := controllers.NewOrderController(orderService)

	route.POST("/orders", orderController.CreateOrder)
	route.GET("/orders", orderController.GetOrders)
	route.PUT("/orders/:orderId", orderController.UpdateOrder)
	route.DELETE("/orders/:orderId", orderController.DeleteOrder)

	route.Run(database.APP_PORT)
}
