package user_cases

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Image struct {
	Id     int    `json:"id,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Buffer string `json:"buffer,omitempty"`
	Link   string `json:"link,omitempty"`
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
	UpdateDataById(Image) (Image, error)
}

const path_to_file = "C:\\Users\\user\\go\\src\\projectServis\\images"

func (s Service) Resize(image Image) (Image, error) {

	if image.Buffer == "" || image.Height == 0 || image.Width == 0 {
		return Image{}, errors.New("not correct data")
	}
	resizedImg, err := s.library.ResizeImageLibrary(image)
	if err != nil {
	}

	//f, err := os.Create("./images/outimage.png")
	//if err != nil {}
	//defer f.Close()

	//Декодируем base64 в байты
	outPngData, err := base64.StdEncoding.DecodeString(resizedImg.Buffer)
	if err != nil {
	}
	fileName := String(10)
	err = ioutil.WriteFile("./images/"+fileName+".png", outPngData, 0644)

	resizedImg.Link = fileName

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

func (s Service) UpdateDataById(image Image) (Image, error) {

	imageFindName, err := s.repository.FindImageId(image.Id)

	files, err := ioutil.ReadDir(path_to_file)
	if err != nil {
	}

	findName := imageFindName.Link + ".png"
	for _, file := range files {
		if file.Name() == findName {
			err := os.Remove(path_to_file + "\\" + findName)
			if err != nil {
			}
		}
		//fmt.Println(file.Name(), file.IsDir())
	}

	//f, err := os.Create("./images/outimage.png")
	//if err != nil {}
	//defer f.Close()

	//Декодируем base64 в байты
	outPngData, err := base64.StdEncoding.DecodeString(image.Buffer)
	if err != nil {
	}
	fileName := String(10)
	err = ioutil.WriteFile("./images/"+fileName+".png", outPngData, 0644)
	image.Link = fileName

	return s.repository.ChangeImageId(image)
}

type LibraryImager interface {
	ResizeImageLibrary(image Image) (Image, error)
}

type RepositoryImager interface {
	HistoryImages() ([]Image, error)
	FindImageId(id int) (Image, error)
	ChangeImageId(Image) (Image, error)
	SaveImage(image Image) (Image, error)
}
