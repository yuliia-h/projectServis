package main

import (
	"fmt"
	"net/http"
	"projectServis/infrastructure"
)

func main() {

	http.HandleFunc("/struct/", infrastructure.HandleResizeImage)

	http.HandleFunc("/getimageId/{id}/", infrastructure.GetImageId)

	http.HandleFunc("/getimages/", infrastructure.GetImages)

	http.HandleFunc("/updateimage/{id}/", infrastructure.UpdateImage)

	//for check THEN: delete
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello mу friend :)"))
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":45998", nil)
}
