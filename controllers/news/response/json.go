package response

import (
	"ca-amartha/businesses/news"
	"time"
)

type News struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CategoryID   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(domain news.Domain) News {
	return News{
		Id:           domain.ID,
		Title:        domain.Title,
		Content:      domain.Content,
		CategoryID:   domain.CategoryID,
		CategoryName: domain.CategoryName,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
