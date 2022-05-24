package repositories

import (
	"assignment-2/models"

	"github.com/jinzhu/gorm"
)

type ItemRepo interface{
	CreateItem (item *models.Item) error
}

type itemRepo struct{
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo{
	return &itemRepo{
		db: db,
	}
}

func (i *itemRepo) CreateItem(item *models.Item) error{
	return i.db.Create(item).Error
}