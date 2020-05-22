package interfaces

import (
	"crypto/rand"
	"log"
	"os"
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

func (r RepositoryImages) SaveImage(image user_cases.Image) error {

	img := Image{}
	namefile, _ := randomfilename16char()
	f, err := os.Create("image" + namefile + ".png")
	if err != nil {
		log.Println(err)
	}
	fileurl := "C:\\Users\\user\\go\\src\\projectServis\\repository\\" + f.Name()
	img.Link = fileurl
	err = r.db.SaveImage(img)
	if err != nil {
		log.Println(err)
	}
	return err
}

func randomfilename16char() (s string, err error) {
	b := make([]byte, 8)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	return s, err
}

type Dbmager interface {
	HistoryAll() ([]Image, error)
	FindImageId(id int) Image
	ChangeImageId(id int)
	SaveImage(image Image) error
}
