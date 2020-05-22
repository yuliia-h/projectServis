package user_cases

import "log"

type Image struct {
	Id     int
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
	GetDataById(id int) (Image, error)
	UpdateDataById(id int) (Image, error)
}

func (s Service) Resize(image Image) (Image, error) {
	image, err := s.library.ResizeImageLibrary(image)
	if err != nil {
		log.Println(err)
	}

	err = s.repository.SaveImage(image)
	if err != nil {
		log.Println(err)
	}

	return image, err
}

func (s Service) History() ([]Image, error) {
	return s.repository.HistoryImages()
}

func (s Service) GetDataById(id int) (Image, error) {
	return s.repository.FindImageId(id)
}

func (s Service) UpdateDataById(id int) (Image, error) {
	return s.repository.ChangeImageId(id)
}

type LibraryImager interface {
	ResizeImageLibrary(image Image) (Image, error)
}

type RepositoryImager interface {
	HistoryImages() ([]Image, error)
	FindImageId(id int) (Image, error)
	ChangeImageId(id int) (Image, error)
	SaveImage(image Image) error
}
