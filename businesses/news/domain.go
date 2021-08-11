package news

import (
	"context"
	"time"
)

type Domain struct {
	ID           int
	Title        string
	Content      string
	CategoryID   int
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IPStat       string
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, newsId int) (Domain, error)
	GetByTitle(ctx context.Context, newsTitle string) (Domain, error)
	Store(ctx context.Context, ip string, newsDomain *Domain) (Domain, error)
	Update(ctx context.Context, newsDomain *Domain) (*Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, newsId int) (Domain, error)
	GetByTitle(ctx context.Context, newsTitle string) (Domain, error)
	Store(ctx context.Context, newsDomain *Domain) (Domain, error)
	Update(ctx context.Context, newsDomain *Domain) (Domain, error)
}
