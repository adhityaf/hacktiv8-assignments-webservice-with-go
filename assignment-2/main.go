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

	router := gin.Default()

	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)

	router.POST("/items", itemController.CreateNewItem)

	// router.POST("/orders")
	// router.GET("/orders")
	// router.PUT("/orders/:orderId")
	// router.DELETE("/orders/:orderId")
	router.Run(database.APP_PORT)
}
