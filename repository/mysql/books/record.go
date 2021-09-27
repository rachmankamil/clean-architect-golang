package books

import (
	"fmt"
	"kampus-merdeka-ca/bussiness/books"
	"kampus-merdeka-ca/repository/mysql/author"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title    string
	ISBN     string
	Status   bool
	AuthorID int
	Author   author.Author `gorm:"foreignKey:author_id"`
}

func toDomain(rec Books) books.Domain {
	fmt.Printf("%+v", rec)
	return books.Domain{
		ID:        int(rec.ID),
		Title:     rec.Title,
		ISBN:      rec.ISBN,
		Status:    rec.Status,
		Author:    rec.Author.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain books.Domain) Books {
	return Books{
		Title:    domain.Title,
		ISBN:     domain.ISBN,
		Status:   domain.Status,
		AuthorID: domain.AuthorID,
	}
}
