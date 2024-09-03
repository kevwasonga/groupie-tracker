package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Artist represents the JSON structure for each artist
type Artist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"-"`
	Relations    map[string][]string `json:"-"`
}

// Locations represents the JSON structure for artist locations
type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// LocationsData holds the index of all locations
type LocationsData struct {
	Index []Locations `json:"index"`
}

// Relations represents the JSON structure for artist relations
type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// RelationsData holds the index of all relations
type RelationsData struct {
	Index []Relations `json:"index"`
}

// FetchAndUnmarshalArtists fetches JSON data from the URL and unmarshals it
func FetchAndUnmarshalArtists() ([]Artist, error) {
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	relationsURL := "https://groupietrackers.herokuapp.com/api/relation"

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Set the timeout duration
	}

	// Fetch JSON data for artists
	artistsResp, err := client.Get(artistsURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching artists JSON data: %w", err)
	}
	defer artistsResp.Body.Close()

	// Check if the request was successful
	if artistsResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d for artists", artistsResp.StatusCode)
	}

	// Decode the JSON data into a slice of Artist structs
	var artists []Artist
	err = json.NewDecoder(artistsResp.Body).Decode(&artists)
	if err != nil {
		return nil, fmt.Errorf("error decoding artists JSON data: %w", err)
	}

	// Fetch JSON data for locations
	locationsResp, err := client.Get(locationsURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching locations JSON data: %w", err)
	}
	defer locationsResp.Body.Close()

	// Check if the request was successful
	if locationsResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d for locations", locationsResp.StatusCode)
	}

	// Decode the JSON data into a LocationsData struct
	var locationsData LocationsData
	err = json.NewDecoder(locationsResp.Body).Decode(&locationsData)
	if err != nil {
		return nil, fmt.Errorf("error decoding locations JSON data: %w", err)
	}

	// Fetch JSON data for relations
	relationsResp, err := client.Get(relationsURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching relations JSON data: %w", err)
	}
	defer relationsResp.Body.Close()

	// Check if the request was successful
	if relationsResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d for relations", relationsResp.StatusCode)
	}

	// Decode the JSON data into a RelationsData struct
	var relationsData RelationsData
	err = json.NewDecoder(relationsResp.Body).Decode(&relationsData)
	if err != nil {
		return nil, fmt.Errorf("error decoding relations JSON data: %w", err)
	}

	// Map locations and relations to their corresponding artists
	for i, artist := range artists {
		for _, loc := range locationsData.Index {
			if artist.ID == loc.ID {
				artists[i].Locations = loc.Locations
				break
			}
		}

		for _, rel := range relationsData.Index {
			if artist.ID == rel.ID {
				artists[i].Relations = rel.DatesLocations
				break
			}
		}
	}

	return artists, nil
}
