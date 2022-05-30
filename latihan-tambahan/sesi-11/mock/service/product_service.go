package service

import (
	"errors"
	"sesi-11/mock/entity"
	"sesi-11/mock/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProduct() (*[]entity.Product, error) {
	products := service.Repository.GetAll()
	if products == nil {
		return nil, errors.New("product not found")
	}

	return products, nil
}
