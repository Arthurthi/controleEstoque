package products

import "gorm.io/gorm"

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindByID(id uint) (*Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) Update(product *Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&Product{}, id).Error
}
