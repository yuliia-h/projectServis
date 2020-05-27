package interfaces

import (
	"projectServis/user_cases"
)

type ImageDb struct {
	Id     int    `db:"id"`
	Width  int    `db:"width"`
	Height int    `db:"height"`
	Link   string `db:"link"`
}

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

	img, err := r.db.HistoryAll()
	var count int = len(img)
	var imageNew = make([]user_cases.Image, count)

	for i := 0; i < len(imageNew); i++ {
		imageNew[i].Id = img[i].Id
		imageNew[i].Width = img[i].Width
		imageNew[i].Height = img[i].Height
		imageNew[i].Link = img[i].Link
	}

	return imageNew, err
}

func (r RepositoryImages) FindImageId(s int) (user_cases.Image, error) {

	img := r.db.FindImageId(s)
	i := user_cases.Image{
		Id:     img.Id,
		Width:  img.Width,
		Height: img.Height,
		Link:   img.Link,
	}
	return i, nil
}

func (r RepositoryImages) ChangeImageId(image user_cases.Image) (user_cases.Image, error) {

	imgDb := ImageDb{
		Id:     image.Id,
		Width:  image.Width,
		Height: image.Height,
		Link:   image.Link,
	}
	imgUpdate := ImageDb{}

	var err error
	_, err = r.db.ChangeImageId(imgDb)
	imgUpdate = r.db.FindImageId(image.Id)
	imguser := user_cases.Image{
		Id:   imgUpdate.Id,
		Link: imgUpdate.Link,
	}
	return imguser, err
}

func (r RepositoryImages) SaveImage(image user_cases.Image) (user_cases.Image, error) {

	imgDb := ImageDb{
		Width:  image.Width,
		Height: image.Height,
		Link:   image.Link,
	}
	imgId := ImageDb{}
	imguser := user_cases.Image{}
	var err error
	if image.Link != "" && image.Height != 0 && image.Width != 0 {
		imgId, err = r.db.SaveImage(imgDb)
		imguser = user_cases.Image{
			Id:     imgId.Id,
			Width:  imgId.Width,
			Height: imgId.Height,
			Link:   imgId.Link,
		}
	}

	return imguser, err
}

type DbImager interface {
	HistoryAll() ([]ImageDb, error)
	FindImageId(id int) ImageDb
	ChangeImageId(ImageDb) (ImageDb, error)
	SaveImage(ImageDb) (ImageDb, error)
}
