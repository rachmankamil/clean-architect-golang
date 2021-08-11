package request

import "ca-amartha/businesses/news"

type News struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID int    `json:"category_id"`
}

func (req *News) ToDomain() *news.Domain {
	return &news.Domain{
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
	}
}
