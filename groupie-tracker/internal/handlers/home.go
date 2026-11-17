package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/internal/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound, "Page not found")
		return
	}

	artists, err := models.FetchArtists()
	if err != nil {
		// Print the real error to your Mac terminal so you can see it
		log.Printf("ERROR FETCHING ARTISTS: %v", err)
		RenderError(w, http.StatusInternalServerError, "Failed to load artists from API: "+err.Error())
		return
	}
	data := models.IndexPageData{
		Query:   "",
		Artists: artists,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to load index template")
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to render home page")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = buf.WriteTo(w)
}
