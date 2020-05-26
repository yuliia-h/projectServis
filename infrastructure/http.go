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
	Id     int    `json:"id"`
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
	//var  b bytes.Buffer
	//err = json.NewEncoder(&b).Encode(ans)
	//if err != nil{
	//	log.Println(err)
	//}

	b, err := json.Marshal(ans)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

// данные картинки по id
func (h Handlers) HandleGetImageById(w http.ResponseWriter, r *http.Request) {
	//считываем весь реквест в body
	//body, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//	panic(err)
	//}
	//
	////формируем ответ передаем в метод структуру и возвращаем ошибку
	//err = h.imgService.GetDataById( )
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	_, err := w.Write([]byte(err.Error()))
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	return
	//}
	//// отправляем статус 200
	////сделать: завернуть картинку и отправить...
	//w.WriteHeader(http.StatusOK)
}

// изменить данные картинки по id
func (h Handlers) HandleUpdateImageById(w http.ResponseWriter, r *http.Request) {
	////считываем весь реквест в body
	//body, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//	panic(err)
	//}
	//
	////создаем структуру
	//image := &user_cases.Image{}
	//
	////парсим json в эту структуру
	//err = json.Unmarshal(body, image)
	//
	////формируем ответ передаем в метод структуру и возвращаем ошибку
	//err = h.imgService.UpdateDataById(image)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	_, err := w.Write([]byte(err.Error()))
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	return
	//}
	//// отправляем статус 200
	////сделать: завернуть картинку и отправить...
	//w.WriteHeader(http.StatusOK)
}
