package request

import "kampus-merdeka-ca/bussiness/books"

type BookInsert struct {
	ISBN     string `json:"isbn"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	Status   bool   `json:"status"`
}

type BookUpdate struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func ToDomain(request BookInsert) *books.Domain {
	return &books.Domain{
		ISBN:     request.ISBN,
		Title:    request.Title,
		AuthorID: request.AuthorID,
		Status:   request.Status,
	}
}
