package category

import (
	"ca-amartha/businesses/category"
	"context"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) category.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) Find(ctx context.Context, active string) ([]category.Domain, error) {
	rec := []Category{}

	query := cr.conn.Where("archive = ?", false)

	if active != "" {
		query = query.Where("active = ?", active)
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

func (cr *categoryRepository) FindByID(id int) (category.Domain, error) {
	rec := Category{}

	if err := cr.conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return category.Domain{}, err
	}
	return rec.ToDomain(), nil
}
