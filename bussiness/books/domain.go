package books

import "time"

// domain layer / entity layer -> acuan utama dalam domain.
type Domain struct {
	ID        int
	Title     string
	ISBN      string
	AuthorID  int
	Author    string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// inteface of bussiness layer -> fungsi fungsi yang di butuhkan oleh domain (bussiness logic)
type Service interface {
	Append(book *Domain) (*Domain, error)
	Update(book *Domain, id int) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Available(generalSearch string) []Domain
}

// interface of data layer -> fungsi fungsi untuk membaca / memerintah database, 3rd Party, ataupun package lain.
type Repository interface {
	Insert(book *Domain) (*Domain, error)
	Update(book *Domain, id int) (*Domain, error)
	FindByID(id int) (*Domain, error)
	FindAll(generalSearch string, availability bool) []Domain
}
