package user_cases

import (
	"log"
	"math/rand"
	"projectServis/interfaces"
	"time"
)

type Image struct {
	Width  int
	Height int
	Buffer string
}

//random numbers
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//random string
func String(length int) string {
	return StringWithCharset(length, charset)
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
	Resize(image Image) (interfaces.ImageLink, error)
	History() ([]interfaces.ImageLink, error)
	GetDataById(id int) (interfaces.ImageLink, error)
	UpdateDataById(id int) (Image, error)
}

func (s Service) Resize(image Image) (interfaces.ImageLink, error) {

	resizedImg, err := s.library.ResizeImageLibrary(image)
	if err != nil {
		log.Println(err)
	}
	imDb := interfaces.ImageDb{
		Width:  resizedImg.Width,
		Height: resizedImg.Height,
		Link:   String(30),
	}

	imgInfo, err := s.repository.SaveImage(imDb)
	if err != nil {
	}
	imgInfo.Buffer = resizedImg.Buffer

	return imgInfo, err
}

func (s Service) History() ([]interfaces.ImageLink, error) {
	return s.repository.HistoryImages()
}

func (s Service) GetDataById(id int) (interfaces.ImageLink, error) {
	return s.repository.FindImageId(id)
}

func (s Service) UpdateDataById(id int) (Image, error) {
	return s.repository.ChangeImageId(id)
}

type LibraryImager interface {
	ResizeImageLibrary(image Image) (Image, error)
}

type RepositoryImager interface {
	HistoryImages() ([]interfaces.ImageLink, error)
	FindImageId(id int) (interfaces.ImageLink, error)
	ChangeImageId(id int) (Image, error)
	SaveImage(interfaces.ImageDb) (interfaces.ImageLink, error)
}
