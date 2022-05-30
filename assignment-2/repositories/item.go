package repositories

import (
	"assignment-2/models"

	"github.com/jinzhu/gorm"
)

type ItemRepo interface {
	FindAll() (*[]models.Item, error)
	FindByCode(code, orderId uint) (*models.Item, error)
	Create(item *models.Item) (*models.Item, error)
	Update(item *models.Item) (*models.Item, error)
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return &itemRepo{
		db: db,
	}
}

func (i *itemRepo) FindAll() (*[]models.Item, error) {
	var items []models.Item

	err := i.db.Find(&items).Error

	return &items, err
}

func (i *itemRepo) FindByCode(code, orderId uint) (*models.Item, error) {
	var item models.Item
	err := i.db.Where("item_code=?",code).Where("order_id=?", orderId).First(&item).Error
	return &item, err
}

func (i *itemRepo) Create(item *models.Item) (*models.Item, error) {
	err := i.db.Create(item).Error
	return item, err
}

func (i *itemRepo) Update(item *models.Item) (*models.Item, error) {
	err := i.db.Save(item).Error
	return item, err
}
