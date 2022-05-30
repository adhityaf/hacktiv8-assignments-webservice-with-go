package repositories

import (
	"assignment-2/models"

	"github.com/jinzhu/gorm"
)

type OrderRepo interface {
	FindAll() (*[]models.Order, error)
	FindById(id uint) (*models.Order, error)
	Create(order *models.Order) (*models.Order, error)
	Update(order *models.Order) (*models.Order, error)
	Delete(order *models.Order) (*models.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) FindAll() (*[]models.Order, error) {
	var orders []models.Order

	err := o.db.Preload("Items").Find(&orders).Error

	return &orders, err
}

func (o *orderRepo) FindById(id uint) (*models.Order, error) {
	var order models.Order

	err := o.db.Preload("Items").First(&order, "id=?", id).Error

	return &order, err
}

func (o *orderRepo) Create(order *models.Order) (*models.Order, error) {
	err := o.db.Create(&order).Error
	return order, err
}

func (o *orderRepo) Update(order *models.Order) (*models.Order, error) {
	err := o.db.Save(&order).Error
	return order, err
}

func (o *orderRepo) Delete(order *models.Order) (*models.Order, error) {
	err := o.db.Delete(&order).Error
	return order, err
}
