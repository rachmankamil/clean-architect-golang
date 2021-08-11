package news

import (
	"ca-amartha/businesses"
	"ca-amartha/businesses/category"
	"ca-amartha/drivers/thirdparties/iplocator"
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"
)

type newsUsecase struct {
	newsRepository  Repository
	categoryUsecase category.Usecase
	contextTimeout  time.Duration
	IpLocator       *iplocator.IPLocator
}

func NewNewsUsecase(nr Repository, cu category.Usecase, timeout time.Duration) Usecase {
	return &newsUsecase{
		newsRepository:  nr,
		categoryUsecase: cu,
		contextTimeout:  timeout,
		IpLocator:       iplocator.NewIPLocator(),
	}
}

func (nu *newsUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := nu.newsRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
func (nu *newsUsecase) GetByID(ctx context.Context, newsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if newsId <= 0 {
		return Domain{}, businesses.ErrNewsIDResource
	}
	res, err := nu.newsRepository.GetByID(ctx, newsId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (nu *newsUsecase) GetByTitle(ctx context.Context, newsTitle string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(newsTitle) == "" {
		return Domain{}, businesses.ErrNewsTitleResource
	}
	res, err := nu.newsRepository.GetByTitle(ctx, newsTitle)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (nu *newsUsecase) Store(ctx context.Context, ip string, newsDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	_, err := nu.categoryUsecase.GetByID(ctx, newsDomain.CategoryID)
	if err != nil {
		return businesses.ErrCategoryNotFound
	}

	existedNews, err := nu.newsRepository.GetByTitle(ctx, newsDomain.Title)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedNews != (Domain{}) {
		return businesses.ErrDuplicateData
	}

	if strings.TrimSpace(ip) != "" {
		ipLoc, err := nu.IpLocator.NewsGetLocationByIP(ctx, ip)
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		jsonMarshal, err := json.Marshal(ipLoc)
		if err != nil {
			log.Default().Printf("%+v", err)
		}

		newsDomain.IPStat = string(jsonMarshal)
	}

	err = nu.newsRepository.Store(ctx, newsDomain)
	if err != nil {
		return err
	}

	return nil
}
