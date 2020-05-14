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
	hadnlfield user_cases.Servicer
}

//конструктор
func NewHandlers(service user_cases.Servicer) *Handlers {
	return &Handlers{
		hadnlfield: service,
	}
}

//--------------------------------------------------------
func (handler Handlers) HandleResizeImage(w http.ResponseWriter, r *http.Request) {

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
	err = handler.hadnlfield.Resize(*image)
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
func (handler Handlers) HistoryImages(w http.ResponseWriter, r *http.Request) {
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
	err = handler.hadnlfield.History(*image)
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
func (handler Handlers) GetImageId(w http.ResponseWriter, r *http.Request) {
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
	err = handler.hadnlfield.GetId(*image)
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
func (handler Handlers) UpdateImage(w http.ResponseWriter, r *http.Request) {
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
	err = handler.hadnlfield.UpdateId(*image)
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
