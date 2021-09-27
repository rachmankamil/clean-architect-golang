package books

import (
	"kampus-merdeka-ca/bussiness/books"

	"gorm.io/gorm"
)

type repoBooks struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) books.Repository {
	return &repoBooks{
		DBConn: db,
	}
}

func (repo *repoBooks) Insert(book *books.Domain) (*books.Domain, error) {
	recordBook := fromDomain(*book)
	if err := repo.DBConn.Create(&recordBook).Error; err != nil {
		return &books.Domain{}, err
	}

	record, err := repo.FindByID(int(recordBook.ID))
	if err != nil {
		return &books.Domain{}, err
	}
	return record, nil
}

func (repo *repoBooks) Update(book *books.Domain, id int) (*books.Domain, error) {
	return &books.Domain{}, nil
}
func (repo *repoBooks) FindByID(id int) (*books.Domain, error) {
	var recordBook Books

	if err := repo.DBConn.Where("books.id = ?", id).Joins("Author").Find(&recordBook).Error; err != nil {
		return &books.Domain{}, err
	}
	result := toDomain(recordBook)
	return &result, nil
}
func (repo *repoBooks) FindAll(generalSearch string, availability bool) []books.Domain {
	return []books.Domain{}
}
