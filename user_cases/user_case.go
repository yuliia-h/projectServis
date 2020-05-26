package user_cases

import (
	"log"
	"math/rand"
	"time"
)

type Image struct {
	Id     int
	Width  int
	Height int
	Buffer string
	Link   string
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
	Resize(image Image) (Image, error)
	History() ([]Image, error)
	GetDataById(id int) (Image, error)
	UpdateDataById(id int) (Image, error)
}

func (s Service) Resize(image Image) (Image, error) {

	resizedImg, err := s.library.ResizeImageLibrary(image)
	if err != nil {
		log.Println(err)
	}
	resizedImg.Link = String(30)

	imgInfo, err := s.repository.SaveImage(resizedImg)
	if err != nil {
	}

	imgInfo.Buffer = resizedImg.Buffer

	return imgInfo, err
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
	SaveImage(image Image) (Image, error)
}
