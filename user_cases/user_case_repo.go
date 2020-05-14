package user_cases

//------------Repository----------------------------------

type RepositoryImager interface {
	HistoryImages(image Image) error
	FindImageId(image Image) error
	ChangeImageId(image Image) error
	GetAllImage() ([]Image, error)
	SaveImage(image Image) error
}

type RepositoryImages struct {
}

// история по измененным картинкам
func (repo RepositoryImages) HistoryImages(image Image) error {
	return nil
}

// данные картинки по id
func (repo RepositoryImages) FindImageId(image Image) error {
	return nil
}

// изменить данные картинки по id
func (repo RepositoryImages) ChangeImageId(image Image) error {
	return nil
}

// выдача всех картинок
func (repo RepositoryImages) GetAllImage() ([]Image, error) {
	return nil, nil
}

// сохранить картинку ?????????????????????
func (repo RepositoryImages) SaveImage(image Image) error {
	return nil
}
