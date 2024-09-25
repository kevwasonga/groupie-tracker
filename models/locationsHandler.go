package models

import (
	"html/template"
	"net/http"
	"path/filepath"

	"groupie/services"
)

// ArtistLocation represents the combined data of artist names and locations
type ArtistLocation struct {
	Name      string   `json:"name"`
	Locations []string `json:"locations"`
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := services.FetchAndUnmarshalArtists()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error fetching artist data")
		return
	}

	artistNameMap := make(map[int]string)
	for _, artist := range artists {
		artistNameMap[artist.ID] = artist.Name
	}

	locationsData, err := services.FetchAndUnmarshalLocations()
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error fetching locations data")
		return
	}

	var artistLocations []ArtistLocation
	for _, loc := range locationsData.Index {
		name, found := artistNameMap[loc.ID]
		if !found {
			name = "Unknown"
		}
		artistLocations = append(artistLocations, ArtistLocation{
			Name:      name,
			Locations: loc.Locations,
		})
	}

	tmpl, err := template.ParseFiles(filepath.Join("templates", "locations.html"))
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error loading template")
		return
	}

	err = tmpl.Execute(w, struct{ Index []ArtistLocation }{Index: artistLocations})
	if err != nil {
		HandleError(w, err, http.StatusInternalServerError, "Error rendering template")
	}
}
