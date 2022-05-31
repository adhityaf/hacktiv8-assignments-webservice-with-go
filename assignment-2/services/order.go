package services

import (
	"assignment-2/models"
	"assignment-2/params"
	"assignment-2/repositories"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
	itemRepo  repositories.ItemRepo
}

func NewOrderService(orderRepo repositories.OrderRepo, itemRepo repositories.ItemRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		itemRepo:  itemRepo,
	}
}

func (o *OrderService) FindAll() *params.Response {
	orders, err := o.orderRepo.FindAll()
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	if orders == nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "Data Not Exist",
			AdditionalInfo: err.Error(),
		}
	}
	
	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve all data",
		Payload: orders,
	}
}

func (o *OrderService) Create(request params.CreateOrder) *params.Response {
	model := models.Order{
		CustomerName: request.CustomerName,
		OrderedAt:    time.Now(),
	}

	order, err := o.orderRepo.Create(&model)
	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	orderId := order.ID
	for i := 0; i < 25; i++ {
		rand.Seed(time.Now().UnixNano())
		code := rand.Intn(10-1) + 1
		
		if item, err := o.itemRepo.FindByCode(uint(code), orderId); err != nil {
			// create new item if code and orderId not found
			modelItem := models.Item{ItemCode: uint(code),
				Description: fmt.Sprintf("Decription Item with code %d ", code),
				Quantity:    1,
				OrderID:     orderId,
			}

			o.itemRepo.Create(&modelItem)
		} else {
			// if item already exist
			// increase quantity and update the data
			item.Quantity += 1

			o.itemRepo.Update(item)
		}
	}

	return &params.Response{
		Status:  201,
		Message: "Order Create Success",
		Payload: order,
	}
}

func (o *OrderService) Update(id uint, request params.UpdateOrder) *params.Response {
	order, err := o.orderRepo.FindById(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "Data Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	order.CustomerName = request.CustomerName

	order, err = o.orderRepo.Update(order)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Success Update Data with id: %d", id),
		Payload: order,
	}
}

func (o *OrderService) Delete(id uint) *params.Response {
	order, err := o.orderRepo.FindById(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "Data Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	order, err = o.orderRepo.Delete(order)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Success Delete Data with id: %d", id),
		Payload: order,
	}
}
