package user_cases

type Image struct {
	Id     string
	Height int
	Width  int
	Buffer []byte
}

//------------Repository----------------------------------
type RepositoryImager interface {
	HistoryImages() (Image, error)
	FindImageId(str string) (Image, error)
	ChangeImageId(image Image) error
	SaveImage(image Image) error
}

//------------Library----------------------------------
type LibraryImager interface {
	ResizeImageLibrary(image Image) error
}

//------------Servicer------------------
type Servicer interface {
	Resize(image Image) (Image, error)
	History() (Image, error)
	GetDataById(string) (Image, error)
	UpdateDataById(string) (Image, error)
}

func NewResizeImager() *Service {
	return nil
}

type Service struct {
	repository RepositoryImager
	library    LibraryImager
}

// изменить размер
func (service Service) Resize(image Image) (Image, error) {
	service.library.ResizeImageLibrary(image)
	return image, nil
}

// история по измененным картинкам
func (service Service) History() (Image, error) {
	return service.repository.HistoryImages()
}

// данные картинки по id
func (service Service) GetDataById(id string) (Image, error) {
	return service.repository.FindImageId(id)
}

// изменить данные картинки по id
func (service Service) UpdateDataById(id string) (Image, error) {
	return service.repository.ChangeImageId(id)
}
