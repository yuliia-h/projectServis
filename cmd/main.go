//Задание:
//Реализовать сервис, который позволяет
//- изменять размер картинок
//- запросить историю по измененным картинкам
//- запросить данные картинки по id
//- изменить данные картинки по id

package main

import (
	"fmt"
	"net/http"
	"projectServis/infrastructure"
	"projectServis/user_cases"
)

func main() {

	// инициализация переменной
	resizeImager := user_cases.NewResizeImager()

	handlers := infrastructure.NewHandlers(resizeImager)

	// изменить размер картинки
	http.HandleFunc("/struct/", handlers.HandleResizeImage)

	// история по измененным картинкам
	http.HandleFunc("/historyimages/", handlers.HistoryImages)

	// данные картинки по id
	http.HandleFunc("/getimage/{id}/", handlers.GetImageId)

	// изменить данные картинки по id
	http.HandleFunc("/updateimage/{id}/", handlers.UpdateImage)

	//for check THEN: delete
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello mу friend :)"))
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":45998", nil)
}
