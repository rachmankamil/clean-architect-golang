package category

import (
	"ca-amartha/businesses"
	"context"
	"time"
)

type categoryUsecase struct {
	categoryRespository Repository
	contextTimeout      time.Duration
}

func NewCategoryUsecase(timeout time.Duration, cr Repository) Usecase {
	return &categoryUsecase{
		contextTimeout:      timeout,
		categoryRespository: cr,
	}
}

func (cu *categoryUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.categoryRespository.Find(ctx, "")
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (cu *categoryUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	if id <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}

	resp, err := cu.categoryRespository.FindByID(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (cu *categoryUsecase) GetByActive(ctx context.Context, active bool) ([]Domain, error) {
	findActive := "false"
	if active {
		findActive = "true"
	}
	resp, err := cu.categoryRespository.Find(ctx, findActive)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}
