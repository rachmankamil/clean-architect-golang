package response

import (
	"kampus-merdeka-ca/bussiness/books"
	"time"
)

type Books struct {
	ISBN      string    `json:"isbn"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Status    bool      `json:"status"`
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain books.Domain) Books {
	return Books{
		ISBN:      domain.ISBN,
		Title:     domain.Title,
		Author:    domain.Author,
		Status:    domain.Status,
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
