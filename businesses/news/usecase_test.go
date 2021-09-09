package news_test

import (
	"ca-amartha/businesses/category"
	_categoryMock "ca-amartha/businesses/category/mocks"
	"ca-amartha/businesses/iplocator"
	_ipLocatorMock "ca-amartha/businesses/iplocator/mocks"
	"ca-amartha/businesses/news"
	_newsMock "ca-amartha/businesses/news/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryUsecase _categoryMock.Usecase
	newsRepository  _newsMock.Repository
	ipLocRepository _ipLocatorMock.Repository

	newsUsecase news.Usecase

	newsDomain              news.Domain
	categoryDomainActive    category.Domain
	categoryDomainNotActive category.Domain
	ipLocDomain             iplocator.Domain
)

func TestMain(m *testing.M) {
	newsUsecase = news.NewNewsUsecase(&newsRepository, &categoryUsecase, time.Hour*1, &ipLocRepository)

	newsDomain = news.Domain{
		ID:         1,
		Title:      "Test Category",
		Content:    "Olympic",
		CategoryID: 1,
	}

	categoryDomainActive = category.Domain{
		ID:     1,
		Title:  "Sport",
		Active: true,
	}
	categoryDomainNotActive = category.Domain{
		ID:     1,
		Title:  "Sport",
		Active: false,
	}
	ipLocDomain = iplocator.Domain{
		City: "Malang",
	}
}

func TestStore(t *testing.T) {
	t.Run("Store | Valid", func(t *testing.T) {
		categoryUsecase.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(categoryDomainActive, nil).Once()
		newsRepository.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(news.Domain{}, nil).Once()
		ipLocRepository.On("GetLocationByIP", mock.Anything, mock.AnythingOfType("string")).Return(ipLocDomain, nil).Once()
		newsRepository.On("Store", mock.Anything, mock.AnythingOfType("*Domain")).Return(newsDomain, nil).Once()

		result, err := newsUsecase.Store(context.Background(), "0.0.0.0", &newsDomain)

		assert.Nil(t, err)
		assert.Contains(t, "Malang", result.IPStat)
	})
}
