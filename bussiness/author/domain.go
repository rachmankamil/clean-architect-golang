package author

type Domain struct {
	Name      string
	Publisher string
}

// inteface of bussiness layer -> fungsi fungsi yang di butuhkan oleh domain (bussiness logic)
type Service interface {
	FindByID(id int) (*Domain, error)
}

// interface of data layer -> fungsi fungsi untuk membaca / memerintah database, 3rd Party, ataupun package lain.
type Repository interface {
	FindByID(id int) (*Domain, error)
}

type RepositoryPublisher interface {
	FindByID(id int) (*Domain, error)
}
