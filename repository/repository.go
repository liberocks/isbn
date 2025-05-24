package repository

type RepositoryInterface interface {
	BookRepositoryInterface
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}
