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

type ItemService struct {
	itemRepo repositories.ItemRepo
}

func NewItemService(repo repositories.ItemRepo) *ItemService {
	return &ItemService{
		itemRepo: repo,
	}
}

func (i *ItemService) Create(request params.CreateItem) *params.Response {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(10 - 1) + 1
	item, err := i.itemRepo.FindByCode(uint(code), request.OrderID)

	if err != nil {
		// create new item if id not found
		model := models.Item{
			ItemCode:    uint(code),
			Description: request.Description,
			Quantity:    request.Quantity,
			OrderID:     request.OrderID,
		}

		item, err := i.itemRepo.Create(&model)
		if err != nil {
			return &params.Response{
				Status:         http.StatusBadRequest,
				Error:          "BAD REQUEST",
				AdditionalInfo: err.Error(),
			}
		}

		return &params.Response{
			Status:  http.StatusCreated,
			Message: "Item Created Success",
			Payload: item,
		}
	}
	// if item already exist
	// increase quantity
	item.Quantity += 1
	item, err = i.itemRepo.Update(item)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Data quantity with id %d increase by 1", code),
		Payload: item,
	}
}

func (i *ItemService) GetItems() *params.Response {
	items, err := i.itemRepo.FindAll()
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	if len(*items) == 0 {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "Data Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success retrieve all items data",
		Payload: items,
	}
}

// Function to generate 6 Code for item_code field
// var data = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// func generateItemCode() string {
// 	b := make([]byte, 2)
// 	n, err := io.ReadAtLeast(rand.Reader, b, 2)
// 	if n != 2 {
// 		panic(err)
// 	}
// 	for i := 0; i < len(b); i++ {
// 		b[i] = data[int(b[i])%len(data)]
// 	}
// 	return string(b)
// }
