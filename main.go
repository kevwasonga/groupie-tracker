package main

import (
	"fmt"
	"html/template" // Package to handle HTML templates
	"net/http"      // Package to create an HTTP server

	"groupie/models"   // Import custom package 'models' to handle data models and request handling
	"groupie/services" // Import custom package 'services' to handle API calls and other services
)

func main() {
	// Route handlers for various endpoints
	http.HandleFunc("/dates", models.DatesHandler)           // Handles requests to "/dates", shows artist dates
	http.HandleFunc("/locations", models.LocationsHandler)   // Handles requests to "/locations", shows artist locations
	http.HandleFunc("/", indexHandler)                       // Handles requests to root, serves the homepage with artists
	http.HandleFunc("/artist/", models.ArtistDetailsHandler) // Handles requests to "/artist/", shows details for a specific artist

	http.HandleFunc("/search", models.SearchHandler) // Handles search requests to "/search"

	// Start the server and display some helpful information in the console
	fmt.Println("Server is starting...")
	fmt.Println("Go on http://localhost:8080/") // Tell user where to access the server
	fmt.Println("To shut down the server press CTRL + C")

	// Start the server on port 8080 and listen for incoming requests
	http.ListenAndServe(":8080", nil) // Use nil for the default multiplexer
}

// Handler for the root URL, which displays the main page (list of artists)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch artists from an external service (probably from an API) using the FetchAndUnmarshalArtists function
	artists, err := services.FetchAndUnmarshalArtists() // Fetch artist data
	if err != nil {
		// If there's an error fetching the data, send a 500 Internal Server Error response with the error message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse and execute the HTML template for the index page
	// 'template.Must' ensures that the parsing error is handled automatically (panics if there is an error)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Execute the template and pass the artist data to it
	// 'w' is the ResponseWriter to send the rendered HTML back to the browser
	tmpl.Execute(w, artists)
}
