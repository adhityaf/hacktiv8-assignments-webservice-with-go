package controllers

import (
	"assignment-2/params"
	"assignment-2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController struct{
	itemService services.ItemService
}

func NewItemController(service *services.ItemService) *ItemController{
	return &ItemController{
		itemService: *service,
	}
}

func (i *ItemController) CreateItem(c *gin.Context){
	var req params.CreateItem

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status: http.StatusBadRequest,
			Error: "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}

	response := i.itemService.Create(req)
	c.JSON(response.Status, response)
}

func (i *ItemController) GetItems(c *gin.Context){
	response := i.itemService.GetItems()
	c.JSON(response.Status, response)
}