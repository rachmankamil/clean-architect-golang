package category_test

import (
	"ca-amartha/businesses"
	category "ca-amartha/businesses/category"
	categoryMock "ca-amartha/businesses/category/mocks"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository categoryMock.Repository
	categoryUsecase    category.Usecase
)

func setup() {
	categoryUsecase = category.NewCategoryUsecase(2, &categoryRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := category.Domain{
			ID:          1,
			Title:       "Sport",
			Description: "TestCase1",
			Active:      true,
		}
		categoryRepository.On("FindByID", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := categoryUsecase.GetByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.Title, result.Title)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := categoryUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, result, category.Domain{})
		assert.Equal(t, err, businesses.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		categoryRepository.On("FindByID", mock.AnythingOfType("int")).Return(category.Domain{}, errNotFound).Once()
		result, err := categoryUsecase.GetByID(context.Background(), 10)

		assert.Equal(t, result, category.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestGetByActive(t *testing.T) {
	t.Run("test case 1, active", func(t *testing.T) {
		domain := []category.Domain{
			{
				ID:          1,
				Title:       "Sport",
				Description: "TestCase1",
				Active:      true,
			},
			{
				ID:          2,
				Title:       "Olympic",
				Description: "TestCase1",
				Active:      true,
			},
		}
		categoryRepository.On("Find", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := categoryUsecase.GetByActive(context.Background(), true)

		assert.Equal(t, 2, len(result))
		assert.Nil(t, err)
	})

	t.Run("test case 2, false", func(t *testing.T) {
		domain := []category.Domain{
			{
				ID:          1,
				Title:       "Sport",
				Description: "TestCase1",
				Active:      false,
			},
		}
		categoryRepository.On("Find", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := categoryUsecase.GetByActive(context.Background(), false)

		assert.Equal(t, 1, len(result))
		assert.Nil(t, err)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errRepository := errors.New("mysql not running")
		categoryRepository.On("Find", mock.Anything, mock.AnythingOfType("string")).Return([]category.Domain{}, errRepository).Once()

		result, err := categoryUsecase.GetByActive(context.Background(), true)

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}
