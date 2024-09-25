package models

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"groupie/services"
)

// SearchHandler handles the search functionality
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error fetching artist data")
		return
	}

	var filteredArtists []services.Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	tmpl, err := template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error loading template")
		return
	}

	err = tmpl.Execute(w, filteredArtists)
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error rendering template")
	}
}
