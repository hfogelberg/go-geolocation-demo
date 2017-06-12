package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Geolocation struct {
	Lat string
	Lon string
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/location/{lat}/{lon}", location)
	http.ListenAndServe(":8080", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.New("index").ParseFiles("templates/index.html"))
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		panic(err)
	}

	prompt := "Detecting your location. Please click the 'Allow' button."
	w.Write([]byte(prompt))
}

func location(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

	fmt.Println("Lat: " + vars["lat"])
	fmt.Println("Lon: " + vars["lon"])

	loc := Geolocation{
		Lat: vars["lat"],
		Lon: vars["lon"],
	}

	fmt.Println("Location struct: ", loc)

	var templates = template.Must(template.New("location").ParseFiles("templates/location.html"))
	if err := templates.ExecuteTemplate(w, "location.html", loc); err != nil {
		panic(err)
	}

}
