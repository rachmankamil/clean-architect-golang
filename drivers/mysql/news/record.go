package news

import (
	newsUsecase "ca-amartha/businesses/news"
	"ca-amartha/drivers/mysql/category"
	"time"
)

type News struct {
	Id         int
	Title      string
	Content    string `gorm:"column:content_data"`
	CategoryID int
	Category   category.Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserStat   string
}

func fromDomain(domain *newsUsecase.Domain) *News {
	return &News{
		Id:         domain.ID,
		Title:      domain.Title,
		Content:    domain.Content,
		CategoryID: domain.CategoryID,
		UserStat:   domain.IPStat,
	}
}

func (rec *News) toDomain() newsUsecase.Domain {
	return newsUsecase.Domain{
		ID:           rec.Id,
		Title:        rec.Title,
		Content:      rec.Content,
		CategoryID:   rec.CategoryID,
		CategoryName: rec.Category.Title,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		IPStat:       rec.UserStat,
	}
}
