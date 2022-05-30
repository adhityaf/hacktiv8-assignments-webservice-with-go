package controllers

import (
	"assignment-2/params"
	"assignment-2/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{
		orderService: *service,
	}
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var req params.CreateOrder

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status: http.StatusBadRequest,
			Error:  "BAD REQUEST",
		})
		return
	}

	response := o.orderService.Create(req)
	fmt.Println(response)
	c.JSON(response.Status, response)
}

func(o *OrderController) GetOrders(c *gin.Context){
	response := o.orderService.FindAll()
	c.JSON(response.Status, response)
}

func (o *OrderController) UpdateOrder(c *gin.Context) {
	var req params.UpdateOrder

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status: http.StatusBadRequest,
			Error:  "BAD REQUEST",
		})
		return
	}

	orderIdString := c.Param("orderId")
	id, err := strconv.Atoi(orderIdString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status: http.StatusBadRequest,
			Error:  "BAD REQUEST",
		})
		return
	}

	response := o.orderService.Update(uint(id), req)
	c.JSON(response.Status, response)
}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	orderIdString := c.Param("orderId")
	id, err := strconv.Atoi(orderIdString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status: http.StatusBadRequest,
			Error:  "BAD REQUEST",
		})
		return
	}

	response := o.orderService.Delete(uint(id))
	c.JSON(response.Status, response)
}