package repository

type RepositoryInterface interface{}

type Repository struct {
	DataBase string
}

func NewRepository(db string) RepositoryInterface {
	return &Repository{DataBase: db}
}
