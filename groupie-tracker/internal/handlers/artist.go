package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/internal/models"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		RenderError(w, http.StatusBadRequest, "Missing artist ID parameter")
		return
	}
//strconv here ensures its num and not id=abc or its under 0 like id=-6
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		RenderError(w, http.StatusBadRequest, "Invalid artist ID parameter")
		return
	}

	detail, err := models.FetchArtistDetail(id)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to load artist details from API")
		return
	}
// it give us 404 error, ex: id=999 is a valid nub but it doesnt exist in our data
	if detail == nil {
		RenderError(w, http.StatusNotFound, "Artist not found")
		return
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to load artist template")
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, detail); err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to render artist page")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = buf.WriteTo(w)
}
