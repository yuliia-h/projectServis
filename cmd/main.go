//Задание:
//Реализовать сервис, который позволяет
//- изменять размер картинок
//- запросить историю по измененным картинкам
//- запросить данные картинки по id
//- изменить данные картинки по id

package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"projectServis/infrastructure"
	"projectServis/interfaces"
	"projectServis/user_cases"
)

func main() {

	//for check THEN: delete
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "form.html")
	})

	connStr := "user=postgres password=ihavetoget5588 dbname=postgres sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbimage := infrastructure.NewDbimageConnect(db)

	libImage := interfaces.NewLibraryImages()

	repoImage := interfaces.NewRepositoryImages(dbimage)

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

	fmt.Println("Server is listening...")
	http.ListenAndServe(":45998", nil)
}
