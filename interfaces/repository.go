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

//
//type ImageLink struct {
//	Id     int
//	Width  int
//	Height int
//	Link   string
//	Buffer string
//}

type RepositoryImages struct {
	//use interface
	db DbImager
}

func NewRepositoryImages(db DbImager) *RepositoryImages {
	return &RepositoryImages{
		db: db,
	}
}

func (r RepositoryImages) HistoryImages() ([]user_cases.Image, error) {

	i := []user_cases.Image{}

	return i, nil
}

func (r RepositoryImages) FindImageId(s int) (user_cases.Image, error) {

	i := user_cases.Image{}
	return i, nil
}

func (r RepositoryImages) ChangeImageId(s int) (user_cases.Image, error) {
	i := user_cases.Image{}
	return i, nil
}

func (r RepositoryImages) SaveImage(image user_cases.Image) (user_cases.Image, error) {

	//imgId, err := r.db.SaveImage(image)
	//imgReturn := ImageDb{
	//	Id:     imgId.Id,
	//	Width:  imgId.Width,
	//	Height: imgId.Height,
	//	Link:   imgId.Link,
	//}
	//imguser := user_cases.Image{
	//	Width:  imgReturn.Id,
	//	Height: imgReturn.Width,
	//	Buffer: imgReturn.Link,
	//}

	return image, nil
}

type DbImager interface {
	HistoryAll() ([]ImageDb, error)
	FindImageId(id int) ImageDb
	ChangeImageId(id int)
	SaveImage(ImageDb) (ImageDb, error)
}
