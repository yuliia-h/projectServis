package interfaces

import (
	"projectServis/user_cases"
)

type ImageDb struct {
	Id     int    `db:" Id "`
	Width  int    `db:" Width "`
	Height int    `db:" Height "`
	Link   string `db:" Link "`
}

type ImageLink struct {
	Id     int
	Width  int
	Height int
	Link   string
	Buffer string
}

type RepositoryImages struct {
	//use interface
	db DbImager
}

func NewRepositoryImages(db DbImager) *RepositoryImages {
	return &RepositoryImages{
		db: db,
	}
}

func (r RepositoryImages) HistoryImages() ([]ImageLink, error) {

	i := []ImageLink{}

	return i, nil
}

func (r RepositoryImages) FindImageId(s int) (ImageLink, error) {

	i := ImageLink{}
	return i, nil
}

func (r RepositoryImages) ChangeImageId(s int) (user_cases.Image, error) {
	i := user_cases.Image{}
	return i, nil
}

func (r RepositoryImages) SaveImage(image ImageDb) (ImageLink, error) {

	imgId, err := r.db.SaveImage(image)
	imgReturn := ImageLink{
		Id:     imgId.Id,
		Width:  imgId.Width,
		Height: imgId.Height,
		Link:   imgId.Link,
		Buffer: "",
	}

	return imgReturn, err
}

type DbImager interface {
	HistoryAll() ([]ImageLink, error)
	FindImageId(id int) ImageLink
	ChangeImageId(id int)
	SaveImage(image ImageDb) (ImageDb, error)
}
