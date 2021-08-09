package news

import (
	newsUsecase "ca-amartha/bussiness/news"
	"time"
)

type News struct {
	Id        int
	Title     string
	Content   string `gorm:"column:content_data"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserStat  string
}

func fromDomain(domain *newsUsecase.Domain) *News {
	return &News{
		Id:       domain.Id,
		Title:    domain.Title,
		Content:  domain.Content,
		UserStat: domain.IPStat,
	}
}

func (rec *News) toDomain() newsUsecase.Domain {
	return newsUsecase.Domain{
		Id:        rec.Id,
		Title:     rec.Title,
		Content:   rec.Content,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		IPStat:    rec.UserStat,
	}
}
