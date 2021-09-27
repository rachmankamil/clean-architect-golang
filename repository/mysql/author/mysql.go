package author

import (
	"kampus-merdeka-ca/bussiness/author"

	"gorm.io/gorm"
)

type repoAuthor struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) author.Repository {
	return &repoAuthor{
		DBConn: db,
	}
}

func (repo *repoAuthor) FindByID(id int) (*author.Domain, error) {
	var recordAuthor Author
	if err := repo.DBConn.Where("id = ?", id).Find(&recordAuthor).Error; err != nil {
		return &author.Domain{}, err
	}
	result := toDomain(recordAuthor)
	return &result, nil
}
