package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projectServis/user_cases"
	"strconv"
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
	Id     int    `json:"id,omitempty"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Buffer string `json:"buffer"`
	Link   string `json: "link"`
}

//--------------------------------------------------------
func (h Handlers) HandleResizeImage(w http.ResponseWriter, r *http.Request) {

	//считываем весь реквест в body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	//создаем структуру
	image := Image{}

	//парсим json в эту структуру
	err = json.Unmarshal(body, &image)
	if err != nil {
		log.Println(err)
	}

	i := user_cases.Image(image)

	//формируем ответ передаем в метод структуру и возвращаем ошибку
	ans, err := h.imgService.Resize(i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}
	b, err := json.Marshal(ans)
	if err != nil {
		log.Println(err)
	}

	w.Write(b)
	// отправляем статус 200
	//сделать: завернуть картинку и отправить...
	//w.WriteHeader(http.StatusOK)
}

// сделать: история по измененным картинкам....
func (h Handlers) HandleHistoryImages(w http.ResponseWriter, r *http.Request) {

	//формируем ответ передаем в метод структуру и возвращаем ошибку
	ans, err := h.imgService.History()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	b, err := json.Marshal(ans)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

// данные картинки по id
func (h Handlers) HandleGetImageById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	//формируем ответ передаем в метод структуру и возвращаем ошибку
	ans, err := h.imgService.GetDataById(s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}
	b, err := json.Marshal(ans)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

// изменить данные картинки по id
func (h Handlers) HandleUpdateImageById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	s, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	//считываем весь реквест в body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	image := Image{}

	//парсим json в эту структуру
	err = json.Unmarshal(body, &image)

	//создаем структуру
	i := user_cases.Image{
		Id:     s,
		Width:  image.Width,
		Height: image.Height,
		Buffer: image.Buffer,
		Link:   image.Link,
	}
	//формируем ответ передаем в метод структуру и возвращаем ошибку
	ans, err := h.imgService.UpdateDataById(i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}
	b, err := json.Marshal(ans)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}
