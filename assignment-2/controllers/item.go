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

func (i *ItemController) CreateNewItem(c *gin.Context){
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

	response := i.itemService.CreateItem(req)
	c.JSON(response.Status, response)
}