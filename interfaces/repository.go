package interfaces

import (
	"projectServis/user_cases"
)

type ImagMap = map[string]user_cases.Image

type RepositoryImages struct {
}

// история по измененным картинкам
func (repo RepositoryImages) HistoryImages() (user_cases.Image, error) {
	i := user_cases.Image{}
	return i, nil
}

// данные картинки по id
func (repo RepositoryImages) FindImageId(s string) (user_cases.Image, error) {
	i := user_cases.Image{}
	return i, nil
}

// изменить данные картинки по id
func (repo RepositoryImages) ChangeImageId(s string) (user_cases.Image, error) {
	i := user_cases.Image{}
	return i, nil
}

// выдача всех картинок
func (repo RepositoryImages) GetAllImage() ([]Image, error) {
	return nil, nil
}

// сохранить картинку ?????????????????????
func (repo RepositoryImages) SaveImage(image Image) error {
	return nil
}
