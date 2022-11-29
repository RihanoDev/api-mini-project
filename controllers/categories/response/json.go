package response

import (
	"api-mini-project/businesses/categories"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" form:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
	Name      string         `json:"name" form:"name"`
}

func FromDomain(domain categories.Domain) Category {
	return Category{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
