package services

import (
	"assignment-2/models"
	"assignment-2/params"
	"assignment-2/repositories"
	"crypto/rand"
	"io"
)

type ItemService struct {
	itemRepo repositories.ItemRepo
}

func NewItemService(repo repositories.ItemRepo) *ItemService {
	return &ItemService{
		itemRepo: repo,
	}
}

func (i *ItemService) CreateItem(request params.CreateItem) *params.Response {
	code := generateItemCode()
	model := models.Item{
		ItemCode:    code,
		Description: request.Description,
		Quantity:    request.Quantity,
	}

	err := i.itemRepo.CreateItem(&model)
	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  201,
		Message: "Created Success",
		Payload: request,
	}
}

// Function to generate 6 Code for item_code field
var table = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func generateItemCode() string {
	b := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, b, 6)
	if n != 6 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
