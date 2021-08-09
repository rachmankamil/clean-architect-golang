package bussiness

import "errors"

var (
	ErrIDNotFound = errors.New("id not found")

	ErrNewsIDResource = errors.New("(NewsID) not found or empty")

	ErrNewsTitleResource = errors.New("(NewsTitle) not found or empty")

	ErrCategoryNotFound = errors.New("category not found")

	ErrDuplicateData = errors.New("duplicate data")
)
