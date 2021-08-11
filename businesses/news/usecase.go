package news

import (
	"ca-amartha/businesses"
	"ca-amartha/businesses/category"
	"ca-amartha/businesses/iplocator"
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
	ipLocator       iplocator.Repository
}

func NewNewsUsecase(nr Repository, cu category.Usecase, timeout time.Duration, il iplocator.Repository) Usecase {
	return &newsUsecase{
		newsRepository:  nr,
		categoryUsecase: cu,
		contextTimeout:  timeout,
		ipLocator:       il,
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
func (nu *newsUsecase) Store(ctx context.Context, ip string, newsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	_, err := nu.categoryUsecase.GetByID(ctx, newsDomain.CategoryID)
	if err != nil {
		return Domain{}, businesses.ErrCategoryNotFound
	}

	existedNews, err := nu.newsRepository.GetByTitle(ctx, newsDomain.Title)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if existedNews != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	if strings.TrimSpace(ip) != "" {
		ipLoc, err := nu.ipLocator.GetLocationByIP(ctx, ip)
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		jsonMarshal, err := json.Marshal(ipLoc)
		if err != nil {
			log.Default().Printf("%+v", err)
		}

		newsDomain.IPStat = string(jsonMarshal)
	}

	result, err := nu.newsRepository.Store(ctx, newsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (nu *newsUsecase) Update(ctx context.Context, newsDomain *Domain) (*Domain, error) {
	existedNews, err := nu.newsRepository.GetByID(ctx, newsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	newsDomain.ID = existedNews.ID

	result, err := nu.newsRepository.Update(ctx, newsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
