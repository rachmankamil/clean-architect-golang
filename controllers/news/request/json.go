package request

import "ca-amartha/businesses/news"

type News struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (req *News) ToDomain() *news.Domain {
	return &news.Domain{
		Title:   req.Title,
		Content: req.Content,
	}
}
