package role

import "gorm.io/gorm"

type Repository interface {
	GetByID(id int) (Role, error)
	Store(role Role) (Role, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetByID(id int) (Role, error) {
	var role Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return Role{}, err
	}

	return role, nil
}

func (r *repository) Store(role Role) (Role, error) {
	err := r.db.Save(&role).Error
	if err != nil {
		return Role{}, err
	}

	return role, nil
}
