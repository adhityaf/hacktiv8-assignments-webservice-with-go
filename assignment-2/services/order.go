package services

import (
	"assignment-2/models"
	"assignment-2/params"
	"assignment-2/repositories"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
}

func NewOrderService(repo repositories.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: repo,
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

	if len(*orders) == 0 {
		return &params.Response{
			Status:         http.StatusNoContent,
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

	// orderId := order.ID
	// for i := 0; i < 25; i++ {
	// 	code := GenerateItemCode()
	// 	ItemController.itemService.Create(models.Item{
	// 		ItemCode:    code,
	// 		Description: "Decription Item with code " + code,
	// 		Quantity:    10,
	// 		OrderID:     orderId,
	// 	})
	// }

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

// Function to generate 6 Code for item_code field
var table = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateItemCode() string {
	b := make([]byte, 2)
	n, err := io.ReadAtLeast(rand.Reader, b, 2)
	if n != 2 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
