package author

type serviceAuthor struct {
	repository Repository
}

func NewService(repoBook Repository) Service {
	return &serviceAuthor{
		repository: repoBook,
	}
}

func (servBook *serviceAuthor) FindByID(book int) (*Domain, error) {
	result, err := servBook.repository.FindByID(book)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
