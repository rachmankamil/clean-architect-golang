package news

import (
	"context"
	"time"
)

type Domain struct {
	Id         int
	Title      string
	Content    string
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IPStat     string
}

type IPStatDomain struct {
	IP          string
	Version     string
	City        string
	Region      string
	CountryName string
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, newsId int) (Domain, error)
	GetByTitle(ctx context.Context, newsTitle string) (Domain, error)
	Store(ctx context.Context, ip string, newsDomain *Domain) error
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, newsId int) (Domain, error)
	GetByTitle(ctx context.Context, newsTitle string) (Domain, error)
	Store(ctx context.Context, newsDomain *Domain) error
}

type IPLocatorRepository interface {
	NewsGetLocationByIP(ctx context.Context, ip string) (IPStatDomain, error)
}
