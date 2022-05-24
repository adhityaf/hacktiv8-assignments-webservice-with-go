package repository

import (
	"sesi-7-gorm/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(*models.Product) error
	GetAllProducts() (*[]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProductByID(id uint, request *models.Product) error
	DeleteProductByID(id uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) CreateProduct(request *models.Product) error {
	err := r.db.Create(request).Error
	return err
}
func (r *productRepo) GetAllProducts() (*[]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return &products, err
}
func (r *productRepo) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, "id=?", id).Error
	return &product, err
}
func (r *productRepo) UpdateProductByID(id uint, request *models.Product) error {
	err := r.db.Where("id=?", id).Updates(request).Error
	return err
}
func (r *productRepo) DeleteProductByID(id uint) error {
	var product *models.Product
	err := r.db.Where("id=?", id).Delete(&product).Error
	return err
}
