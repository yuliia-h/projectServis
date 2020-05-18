package user_cases

type Image struct {
	Id     string
	Width  int
	Height int
	Buffer string
}

type Service struct {
	library    LibraryImager
	repository RepositoryImager
}

func NewService(lib LibraryImager, repo RepositoryImager) *Service {
	return &Service{
		library:    lib,
		repository: repo,
	}
}

type Servicer interface {
	Resize(image Image) (Image, error)
	History() ([]Image, error)
	GetDataById(id string) (Image, error)
	UpdateDataById(id string) (Image, error)
}

func (s Service) Resize(image Image) (Image, error) {
	return s.library.ResizeImageLibrary(image)
}

func (s Service) History() ([]Image, error) {
	return s.repository.HistoryImages()
}

func (s Service) GetDataById(id string) (Image, error) {
	return s.repository.FindImageId(id)
}

func (s Service) UpdateDataById(id string) (Image, error) {
	return s.repository.ChangeImageId(id)
}

type LibraryImager interface {
	ResizeImageLibrary(image Image) (Image, error)
}

type RepositoryImager interface {
	HistoryImages() ([]Image, error)
	FindImageId(id string) (Image, error)
	ChangeImageId(id string) (Image, error)
	SaveImage(image Image) error
}
