package products

import "time"

type Product struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Price     int       `json:"price" gorm:"not null"`
	Stock     int       `json:"stock" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
