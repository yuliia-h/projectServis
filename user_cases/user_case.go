package user_cases

type Image struct {
	Id     string `json:"id"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Buffer []byte `json:"buffer"`
}

//------------Servicer------------------
type Servicer interface {
	Resize(image Image) error
	History(image Image) error
	GetId(image Image) error
	UpdateId(image Image) error
}

func NewResizeImager() *Service {
	return nil
}

type Service struct {
	repository RepositoryImages
	library    LibraryImages
}

// изменить размер
func (service Service) Resize(image Image) error {
	service.library.ResizeImageLibrary(image)
	return nil
}

// история по измененным картинкам
func (service Service) History(image Image) error {
	service.repository.HistoryImages(image)
	return nil
}

// данные картинки по id
func (service Service) GetId(image Image) error {
	service.repository.FindImageId(image)
	return nil
}

// изменить данные картинки по id
func (service Service) UpdateId(image Image) error {
	service.repository.ChangeImageId(image)
	return nil
}
