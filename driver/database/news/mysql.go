package news

import (
	"ca-amartha/bussiness/news"
	"context"

	"gorm.io/gorm"
)

type mysqlNewsRepository struct {
	Conn *gorm.DB
}

func NewMySQLNewsRepository(conn *gorm.DB) news.Repository {
	return &mysqlNewsRepository{
		Conn: conn,
	}
}

func (nr *mysqlNewsRepository) Fetch(ctx context.Context, page, perpage int) ([]news.Domain, int, error) {
	rec := []News{}

	offset := (page - 1) * perpage
	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []news.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []news.Domain{}, 0, err
	}

	var domainNews []news.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (nr *mysqlNewsRepository) GetByID(ctx context.Context, newsId int) (news.Domain, error) {
	rec := News{}
	err := nr.Conn.Where("id = ?", newsId).First(&rec).Error
	if err != nil {
		return news.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlNewsRepository) GetByTitle(ctx context.Context, newsTitle string) (news.Domain, error) {
	rec := News{}
	err := nr.Conn.Where("title = ?", newsTitle).First(&rec).Error
	if err != nil {
		return news.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlNewsRepository) Store(ctx context.Context, newsDomain *news.Domain) error {
	rec := fromDomain(newsDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
