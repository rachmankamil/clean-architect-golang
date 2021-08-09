package response

import (
	"ca-amartha/bussiness/news"
	"time"
)

type News struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain news.Domain) News {
	return News{
		Id:        domain.Id,
		Title:     domain.Title,
		Content:   domain.Content,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
