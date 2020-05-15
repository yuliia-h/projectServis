package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"projectServis/user_cases"
)

//интерфейсный тип
type Handlers struct {
	imgService user_cases.Servicer
}

//конструктор
func NewHandlers(service user_cases.Servicer) *Handlers {
	return &Handlers{
		imgService: service,
	}
}

type Image struct {
	Id     string `json:"id"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Buffer []byte `json:"buffer"`
}

//--------------------------------------------------------
func (h Handlers) HandleResizeImage(w http.ResponseWriter, r *http.Request) {

	//считываем весь реквест в body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//создаем структуру
	image := Image{}
	i := user_cases.Image(image)

	//парсим json в эту структуру
	err = json.Unmarshal(body, i)

	//формируем ответ передаем в метод структуру и возвращаем ошибку
	err = h.imgService.Resize(*image)
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

// сделать: история по измененным картинкам....
func (h Handlers) HistoryImages(w http.ResponseWriter, r *http.Request) {
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
	err = h.imgService.History(*image)
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

// данные картинки по id
func (h Handlers) GetImageId(w http.ResponseWriter, r *http.Request) {
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
	err = h.imgService.GetDataById(image)
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

// изменить данные картинки по id
func (h Handlers) UpdateImage(w http.ResponseWriter, r *http.Request) {
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
	err = h.imgService.UpdateDataById(image)
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
