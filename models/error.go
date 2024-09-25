package models

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// HandleError is a centralized error handler function
func HandleError(w http.ResponseWriter, err error, statusCode int, userMessage string) {
	if err != nil {
		log.Printf("Error: %v", err) // Log the error for debugging
	} else {
		log.Printf("User Message: %s", userMessage) // Log user-friendly message
	}

	// Set the response status code
	w.WriteHeader(statusCode)

	// Parse the error template
	tmpl, templateErr := template.ParseFiles(filepath.Join("templates", "error.html"))
	if templateErr != nil {
		log.Printf("Template Error: %v", templateErr)
		http.Error(w, "Error loading error template", http.StatusInternalServerError)
		return
	}

	// Execute the template with error message
	errExec := tmpl.Execute(w, struct {
		Message string
		Status  int
	}{
		Message: userMessage,
		Status:  statusCode,
	})
	if errExec != nil {
		log.Printf("Execution Error: %v", errExec)
		http.Error(w, "Error rendering error template", http.StatusInternalServerError)
	}
}
