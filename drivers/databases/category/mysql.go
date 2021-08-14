package category

import (
	"ca-amartha/businesses/category"
	"context"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	conn *gorm.DB
}

//NewMySQLRepository we need this to work around the repository test
func NewMySQLRepository(conn *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		conn: conn,
	}
}

func (cr *MySQLRepository) Find(ctx context.Context, active string) ([]category.Domain, error) {
	rec := []Category{}

	query := cr.conn.Debug().Where("archive = ?", false)

	if active != "" {
		if active == "false" {
			query = query.Where("active = ?", false)
		} else {
			query = query.Where("active = ?", true)
		}
	}

	err := query.Find(&rec).Error
	if err != nil {
		return []category.Domain{}, err
	}

	categoryDomain := []category.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.ToDomain())
	}

	return categoryDomain, nil
}

func (cr *MySQLRepository) FindByID(id int) (category.Domain, error) {
	rec := Category{}

	if err := cr.conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return category.Domain{}, err
	}
	return rec.ToDomain(), nil
}
