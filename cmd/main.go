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
	"projectServis/interfaces"
	"projectServis/user_cases"
)

func main() {

	// инициализация переменной

	libImage := interfaces.NewLibraryImages()

	repoImage := interfaces.NewRepositoryImages()

	resizeImager := user_cases.NewService(libImage, repoImage)

	handlers := infrastructure.NewHandlers(resizeImager)

	// изменить размер картинки
	http.HandleFunc("/struct/", handlers.HandleResizeImage)

	// история по измененным картинкам
	http.HandleFunc("/historyimages/", handlers.HandleHistoryImages)

	// данные картинки по id
	http.HandleFunc("/getimage/{id}/", handlers.HandleGetImageById)

	// изменить данные картинки по id
	http.HandleFunc("/updateimage/{id}/", handlers.HandleUpdateImageById)

	//for check THEN: delete
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "form.html")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":45998", nil)
}
