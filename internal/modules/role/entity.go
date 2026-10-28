package role

import "time"

type Role struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name" gorm:"type:varchar(60);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime"`
	UpdatedAt time.Time `json:"updated_at"`
}
