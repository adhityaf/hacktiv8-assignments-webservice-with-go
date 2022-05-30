package service

import (
	"sesi-11/mock/entity"
	"sesi-11/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T){
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T){
	product := entity.Product{Id: "2", Name: "kacamata"}
	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.Id, result.Id, "result id has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result name has to be 'kacamata'")
	assert.Equal(t, &product, result, "result has to be product with id '2'")
}

func TestProductServiceGetAllProducts(t *testing.T){
	products := []entity.Product{
		{Id: "2", Name: "kacamata"},
		{Id: "3", Name: "kacamata"},
	}
	productRepository.Mock.On("GetAll").Return(products)

	result, err := productService.GetAllProduct()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestProductServiceGetAllProductsNotFound(t *testing.T){
	productRepository.Mock.On("GetAll").Return(nil)

	result, err := productService.GetAllProduct()
	assert.Nil(t, err)
	assert.NotNil(t, result)
}