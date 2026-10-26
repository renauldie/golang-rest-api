package products

import "gorm.io/gorm"

type Repository interface {
	FindAll(limit int, offset int) ([]Product, int64, error)
	FindById(id int) (Product, error)
	CreateProduct(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

// Constructor for repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(limit int, offset int) ([]Product, int64, error) {
	//var products []Product
	//err := r.db.Find(&products).Error
	//return products, err
	var products []Product
	var total int64

	err := r.db.Model(&products).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}

func (r *repository) CreateProduct(product Product) (Product, error) {
	// create only
	//err := r.db.Create(&product).Error
	//if err != nil {
	//	return Product{}, err
	//}
	//return product, nil

	// save or update with id
	err := r.db.Save(&product).Error
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *repository) FindById(id int) (Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
