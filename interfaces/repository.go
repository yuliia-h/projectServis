package interfaces

import (
	"log"
	"projectServis/user_cases"
)

type Image struct {
	Id   int    `db:" id "`
	Link string `db:" link "`
}

type RepositoryImages struct {
	//use interface
	db Dbmager
}

func NewRepositoryImages(db Dbmager) *RepositoryImages {
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

func (r RepositoryImages) SaveImage(image user_cases.ImageDb) (user_cases.ImageDb, error) {

	image, err := r.db.SaveImage(image)
	if err != nil {
		log.Println(err)
	}
	return image, err
}

type Dbmager interface {
	HistoryAll() ([]Image, error)
	FindImageId(id int) Image
	ChangeImageId(id int)
	SaveImage(image user_cases.ImageDb) (user_cases.ImageDb, error)
}
