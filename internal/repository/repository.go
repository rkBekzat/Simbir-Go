package repository

type Repo interface {
}

type repo struct {
}

func NewRepo() Repo {
	return &repo{}
}
