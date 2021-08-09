package category

import (
	"ca-amartha/bussiness/category"
	"time"
)

type Category struct {
	ID          int
	Title       string
	Description string
	Active      bool
	Archive     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (rec *Category) ToDomain() category.Domain {
	return category.Domain{
		ID:        rec.ID,
		Title:     rec.Title,
		Active:    rec.Active,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
