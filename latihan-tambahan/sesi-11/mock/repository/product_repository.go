package repository

import "sesi-11/mock/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	GetAll() *[]entity.Product
}
