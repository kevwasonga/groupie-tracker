package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"strings"

	"groupie/models"
	"groupie/services"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/artist/", models.ArtistDetailsHandler)
	mux.HandleFunc("/search", models.SearchHandler)

	mux.HandleFunc("/dates", protectedHandler(models.DatesHandler))
	mux.HandleFunc("/locations", protectedHandler(models.LocationsHandler))

	fmt.Println("Server is starting...")
	fmt.Println("Go on http://localhost:8080/")
	fmt.Println("To shut down the server press CTRL + C")
	http.ListenAndServe(":8080", mux)
}

func protectedHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		referer := r.Header.Get("Referer")
		if referer == "" || !strings.HasPrefix(referer, "http://localhost:8080") {
			models.HandleError(w, errors.New("access denied"), 403, "Access Denied Cant parse directly on url")
			return
		}

		handler(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	allowedURLs := []string{"/", "/artist/", "/search", "/dates", "/locations"}
	if !(slices.Contains(allowedURLs, string(r.URL.Path))) {
		models.HandleError(w, errors.New("not found"), 404, "Not Found in our servers")
		return
	}
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, artists)
}
