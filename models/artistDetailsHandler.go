package models

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"groupie/services"
)

// ArtistDetailsHandler serves the details of a single artist
func ArtistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the artist ID from the URL path
	artistIDStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		HandleError(w, err, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error fetching artist data")
		return
	}

	// Find the artist with the given ID
	var artist *services.Artist
	for _, a := range artists {
		if a.ID == artistID {
			artist = &a
			break
		}
	}

	if artist == nil {
		HandleError(w, err, http.StatusNotFound, "Artist not found")
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join("templates", "artistdetails.html"))
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error loading artist details template")
		return
	}

	err = tmpl.Execute(w, artist)
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error rendering artist details")
	}
}
