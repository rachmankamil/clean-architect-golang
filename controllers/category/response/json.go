package response

import (
	"ca-amartha/businesses/category"
	"time"
)

type Category struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Active      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain category.Domain) Category {
	return Category{
		Id:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		Active:      domain.Active,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
