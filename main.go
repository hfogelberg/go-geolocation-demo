package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

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

	// fmt.Println("Location struct: ", loc)

	name, err := reverseGeocode(loc.Lat, loc.Lon)
	if err != nil {
		fmt.Println(err)
	}

	loc.Name = name
	fmt.Println(loc.Name)

	var templates = template.Must(template.New("location").ParseFiles("templates/location.html"))
	if err := templates.ExecuteTemplate(w, "location.html", loc); err != nil {
		panic(err)
	}
}

func reverseGeocode(lat string, lon string) (string, error) {
	var name string
	key := os.Getenv("GOOGLE_MAPS_KEY")
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s", lat, lon, key)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var g = new(Response)
	err = json.NewDecoder(resp.Body).Decode(g)

	if err != nil {
		return "", err
	}

	acs := g.Results[0].AddressComponents

	for _, ac := range acs {
		t := ac.Types[0]
		if t == "postal_town" {
			name = ac.ShortName
		}
		if t == "locality" && name == "" {
			name = ac.ShortName
		}
		if t == "administrative_area_level_1" && name == "" {
			name = ac.ShortName
		}
		if t == "administrative_area_level_2" && name == "" {
			name = ac.ShortName
		}
	}

	return name, nil
}
