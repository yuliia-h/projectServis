package infrastructure

import (
	"encoding/json"
	_ "github.com/facebookgo/inject"
	"io/ioutil"
	"log"
	"net/http"
	"projectServis/user_cases"
	_ "projectServis/user_cases"
)

type ServConn struct {
	image   *user_cases.Image   `inject:""`
	service *user_cases.Service `inject:""`
}

func New(image *user_cases.Image) *ServConn {
	return &ServConn{
		image: image,
	}
}

//var conn ServConn

func HandleResizeImage(w http.ResponseWriter, r *http.Request) {

	//считываем весь реквест в body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//создаем структуру
	image := &user_cases.Image{}

	//парсим json в эту структуру
	err = json.Unmarshal(body, image)

	//формируем ответ передаем в метод структуру и возвращаем ошибку
	err = conn.service.Resize(*image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}
	// отправляем статус 200
	//сделать: завернуть картинку и отправить...
	w.WriteHeader(http.StatusOK)
}

// выводит весь массив картинок на экран
//возвращать массив картинок и ошибку....
func GetImages(w http.ResponseWriter, r *http.Request) {
	conn.service.GetImages()
}

//поиск картинки по id
func GetImageId(w http.ResponseWriter, r *http.Request) {
	conn.service.GetImages()
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {

}
