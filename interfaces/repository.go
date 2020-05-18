package interfaces

import (
	"projectServis/user_cases"
)

type RepositoryImages struct {
}

func NewRepositoryImages() *RepositoryImages {
	return &RepositoryImages{}
}

func (r RepositoryImages) HistoryImages() ([]user_cases.Image, error) {

	i := []user_cases.Image{}
	return i, nil
}

func (r RepositoryImages) FindImageId(s string) (user_cases.Image, error) {
	i := user_cases.Image{}
	return i, nil
}

func (r RepositoryImages) ChangeImageId(s string) (user_cases.Image, error) {
	i := user_cases.Image{}
	return i, nil
}

func (r RepositoryImages) GetAllImage() ([]user_cases.Image, error) {
	return nil, nil
}

func (r RepositoryImages) SaveImage(image user_cases.Image) error {
	return nil
}
