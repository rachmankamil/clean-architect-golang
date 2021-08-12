// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	news "ca-amartha/businesses/news"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, page, perpage
func (_m *Usecase) Fetch(ctx context.Context, page int, perpage int) ([]news.Domain, int, error) {
	ret := _m.Called(ctx, page, perpage)

	var r0 []news.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []news.Domain); ok {
		r0 = rf(ctx, page, perpage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]news.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, int, int) int); ok {
		r1 = rf(ctx, page, perpage)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int, int) error); ok {
		r2 = rf(ctx, page, perpage)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, newsId
func (_m *Usecase) GetByID(ctx context.Context, newsId int) (news.Domain, error) {
	ret := _m.Called(ctx, newsId)

	var r0 news.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) news.Domain); ok {
		r0 = rf(ctx, newsId)
	} else {
		r0 = ret.Get(0).(news.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, newsId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, newsTitle
func (_m *Usecase) GetByTitle(ctx context.Context, newsTitle string) (news.Domain, error) {
	ret := _m.Called(ctx, newsTitle)

	var r0 news.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) news.Domain); ok {
		r0 = rf(ctx, newsTitle)
	} else {
		r0 = ret.Get(0).(news.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, newsTitle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, ip, newsDomain
func (_m *Usecase) Store(ctx context.Context, ip string, newsDomain *news.Domain) (news.Domain, error) {
	ret := _m.Called(ctx, ip, newsDomain)

	var r0 news.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, *news.Domain) news.Domain); ok {
		r0 = rf(ctx, ip, newsDomain)
	} else {
		r0 = ret.Get(0).(news.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *news.Domain) error); ok {
		r1 = rf(ctx, ip, newsDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, newsDomain
func (_m *Usecase) Update(ctx context.Context, newsDomain *news.Domain) (*news.Domain, error) {
	ret := _m.Called(ctx, newsDomain)

	var r0 *news.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *news.Domain) *news.Domain); ok {
		r0 = rf(ctx, newsDomain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*news.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *news.Domain) error); ok {
		r1 = rf(ctx, newsDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
