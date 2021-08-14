package users

import (
	"ca-amartha/businesses/users"
	"context"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		Conn: conn,
	}
}

func (nr *MySQLRepository) Fetch(ctx context.Context, page, perpage int) ([]users.Domain, int, error) {
	rec := []Users{}

	offset := (page - 1) * perpage
	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var domainNews []users.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (nr *MySQLRepository) GetByID(ctx context.Context, userId int) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Where("id = ?", userId).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *MySQLRepository) GetByUsername(ctx context.Context, username string) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Where("username = ?", username).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *MySQLRepository) Store(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
