package books

import (
	"kampus-merdeka-ca/bussiness/author"
	_authorDomain "kampus-merdeka-ca/bussiness/author"
	_cacheDomain "kampus-merdeka-ca/bussiness/cache"
)

type serviceBooks struct {
	repository      Repository
	authorDomain    _authorDomain.Service
	cacheRepository _cacheDomain.Repository
}

func NewService(repoBook Repository, repoCache _cacheDomain.Repository, authorServ author.Service) Service {
	return &serviceBooks{
		repository:      repoBook,
		cacheRepository: repoCache,
		authorDomain:    authorServ,
	}
}

func (servBook *serviceBooks) Append(book *Domain) (*Domain, error) {

	respo, err := servBook.authorDomain.FindByID(book.AuthorID)
	if err != nil {
		return &Domain{}, err
	}
	book.Author = respo.Name

	result, err := servBook.repository.Insert(book)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}
func (servBook *serviceBooks) Update(book *Domain, id int) (*Domain, error) {
	return &Domain{}, nil
}
func (servBook *serviceBooks) FindByID(id int) (*Domain, error) {
	return &Domain{}, nil
}
func (servBook *serviceBooks) Available(generalSearch string) []Domain {
	return []Domain{}
}
