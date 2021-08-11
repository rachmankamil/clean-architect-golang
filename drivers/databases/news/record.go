package news

import (
	newsUsecase "ca-amartha/businesses/news"
	"ca-amartha/drivers/databases/category"
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
		Id:         domain.Id,
		Title:      domain.Title,
		Content:    domain.Content,
		CategoryID: domain.CategoryID,
		UserStat:   domain.IPStat,
	}
}

func (rec *News) toDomain() newsUsecase.Domain {
	return newsUsecase.Domain{
		Id:         rec.Id,
		Title:      rec.Title,
		Content:    rec.Content,
		CategoryID: rec.CategoryID,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		IPStat:     rec.UserStat,
	}
}
