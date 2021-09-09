package category_test

import (
	"ca-amartha/businesses"
	"ca-amartha/businesses/category"
	_categoryMock "ca-amartha/businesses/category/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository _categoryMock.Repository
	categoryService    category.Usecase

	categoryDomain category.Domain
)

func TestMain(m *testing.M) {
	categoryService = category.NewCategoryUsecase(time.Hour*1, &categoryRepository)
	categoryDomain = category.Domain{
		ID:          1,
		Title:       "Test Category",
		Description: "Description",
		Active:      true,
	}
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		categoryRepository.On("Find", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return([]category.Domain{categoryDomain}, nil).Once()

		result, err := categoryService.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		categoryRepository.On("Find", mock.AnythingOfType("context.Context"), mock.AnythingOfType("string")).Return([]category.Domain{}, businesses.ErrCategoryNotFound).Once()

		_, err := categoryService.GetAll(context.Background())

		assert.NotNil(t, err)
	})
}
