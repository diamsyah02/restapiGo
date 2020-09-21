package main

import (
	"log"
	"net/http"

	"github.com/restapiGo/controllers"
)

func main() {

	// route
	http.HandleFunc("/mahasiswa", controllers.GetMahasiswa)
	http.HandleFunc("/mahasiswa/create", controllers.PostMahasiswa)
	http.HandleFunc("/mahasiswa/update", controllers.UpdateMahasiswa)
	http.HandleFunc("/mahasiswa/delete", controllers.DeleteMahasiswa)

	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
